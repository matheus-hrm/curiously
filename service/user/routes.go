package user

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitub.com/matheus-hrm/curiously/internal/auth"
	"gitub.com/matheus-hrm/curiously/types"
	"gitub.com/matheus-hrm/curiously/utils"
)

type Handler struct {
	store types.UserStorage
}

func NewHandler(store types.UserStorage) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
	router.GET("/user/:id", h.handleGetUserProfile)
}

func (h *Handler) handleLogin(c *gin.Context) {
	var payload types.LoginUserPayload

	if err := utils.ParseJson(c.Request, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}
	u, err := h.store.GetUserByEmail(payload.Email, c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid email"))
		return
	}

	if !auth.ComparePasswords(u.Password_Hash, []byte(payload.Password)) {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid email or password"))
		return
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	utils.WriteJson(c, http.StatusOK, gin.H{"token": token})
}

func (h *Handler) handleRegister(c *gin.Context) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(c.Request, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}
	_, err := h.store.GetUserByEmail(payload.Email, c)
	if err == nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("email already in use"))
		return
	}
	hash, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	user := types.User{
		Email:         payload.Email,
		Username:      payload.Username,
		Password_Hash: hash,
	}
	if err := h.store.CreateUser(user, c); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	utils.WriteJson(c, http.StatusCreated, gin.H{"message": "user created"})
}

func (h *Handler) handleGetUserProfile(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	user, err := h.store.GetUserByID(userId, c)
	if err != nil {
		utils.WriteError(c, http.StatusNotFound, errors.New("user not found"))
		return
	}

	questions, err := h.store.GetQuestionsByUserID(userId, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	profile := types.UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Questions: make([]types.ProfileQuestion, len(questions)),
		CreatedAt: user.CreatedAt,
	}
	// TODO: check if questions are answered
	for i, q := range questions {
		profile.Questions[i] = types.ProfileQuestion{
			ID:        q.ID,
			Content:   q.Content,
			CreatedAt: q.CreatedAt,
			// Answered: false,
		}
	}

	utils.WriteJson(c, http.StatusOK, profile)
}
