package user

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitub.com/matheus-hrm/curiously/internal/auth"
	"gitub.com/matheus-hrm/curiously/types"
	"gitub.com/matheus-hrm/curiously/utils"
)

type Handler struct {
	userStore   types.UserStorage
	answerStore types.AnswerStorage
}

func NewHandler(userStore types.UserStorage, answerStore types.AnswerStorage) *Handler {
	return &Handler{
		userStore:   userStore,
		answerStore: answerStore,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
	router.GET("/user/:username", h.handleGetUserProfile)
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
	u, err := h.userStore.GetUserByEmail(payload.Email, c)
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

	utils.WriteJson(c, http.StatusOK, gin.H{"token": token, "username": u.Username})
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
	_, err := h.userStore.GetUserByEmail(payload.Email, c)
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
	if err := h.userStore.CreateUser(user, c); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	utils.WriteJson(c, http.StatusCreated, gin.H{"message": "user created"})
}

func (h *Handler) handleGetUserProfile(c *gin.Context) {
	username := c.Param("username")
	user, err := h.userStore.GetUserByUsername(username, c)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid username"))
		return
	}
	questions, err := h.userStore.GetQuestionsByUserID(user.ID, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("failed to fetch questions"))
		return
	}
	profile := types.UserProfile{
		Email:     user.Email,
		Username:  user.Username,
		Questions: make([]types.ProfileQuestion, len(questions)),
		CreatedAt: user.CreatedAt,
	}

	for _, q := range questions {
		pq := types.ProfileQuestion{
			ID:        q.ID,
			Content:   q.Content,
			CreatedAt: q.CreatedAt,
			Answer:    []string{},
			Answered:  false,
		}

		ans, err := h.answerStore.GetAnswersByQuestionID(q.ID, c)
		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, errors.New("failed to fetch answers"))
			return
		}
		if len(ans) > 0 {
			pq.Answered = true
			for _, a := range ans {
				pq.Answer = append(pq.Answer, a.Content)
			}
		}
		profile.Questions = append(profile.Questions, pq)
	}

	utils.WriteJson(c, http.StatusOK, profile)
}
