package answers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/matheus-hrm/curiously/internal/auth"
	"gitub.com/matheus-hrm/curiously/types"
	"gitub.com/matheus-hrm/curiously/utils"
)

type Handler struct {
	answerStore types.AnswerStorage
	userStore   types.UserStorage
}

func NewHandler(answerStore types.AnswerStorage, userStore types.UserStorage) *Handler {
	return &Handler{
		answerStore: answerStore,
		userStore:   userStore,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/answer", auth.WithJWTAuth(h.userStore, h.handleCreateAnswer))
	router.GET("/answer/:id", h.handleGetAnswer)
	router.GET("/answers/question/:id", h.handleGetAnswersByQuestionID)
}

func (h *Handler) handleCreateAnswer(c *gin.Context) {
	var payload types.CreateAnswerPayload

	if err := utils.ParseJson(c.Request, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}
	answer, err := h.answerStore.CreateAnswer(payload, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, errors.New("error creating answer"))
		return
	}

	utils.WriteJson(c, http.StatusOK, answer)
}

func (h *Handler) handleGetAnswer(c *gin.Context) {
	id, err := utils.GetIDFromParam(c, "id")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	answer, err := h.answerStore.GetAnswerByID(id, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, answer)
}

func (h *Handler) handleGetAnswersByQuestionID(c *gin.Context) {
	id, err := utils.GetIDFromParam(c, "id")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	answers, err := h.answerStore.GetAnswersByQuestionID(id, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, answers)
}
