package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gitub.com/matheus-hrm/curiously/types"
	"gitub.com/matheus-hrm/curiously/utils"
)

func CreateJWT(secret []byte, id int) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("error creating token")
	}
	return tokenString, nil
}

func WithJWTAuth(store types.UserStorage, handler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getTokenFromRequest(c)

		token, err := ValidateToken(tokenString)
		if err != nil {
			log.Printf("error validating token: %s", err)
			PermissionDenied(c)
			return
		}

		if !token.Valid {
			log.Printf("token is invalid")
			PermissionDenied(c)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["id"].(string)
		id, err := strconv.Atoi(str)

		log.Printf("id: %d", id)
		user, err := store.GetUserByID(id, c)
		if err != nil {
			log.Fatal(err)
			utils.WriteError(c, http.StatusNotFound, errors.New("user not found"))
			return
		}

		c.Set("user", user)
		handler(c)
	}
}

func getTokenFromRequest(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.WriteError(c, http.StatusUnauthorized, errors.New("missing Authorization header"))
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid Authorization header"))
		return ""
	}

	return splitToken[1]
}

func ValidateToken(t string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func PermissionDenied(c *gin.Context) {
	utils.WriteError(c, http.StatusForbidden, errors.New("permission denied"))
}

func GetIDFromToken(c *gin.Context) (int, error) {
	tokenString := getTokenFromRequest(c)
	token, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	str := claims["id"].(string)
	id, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return id, nil
}
