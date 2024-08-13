package question

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitub.com/matheus-hrm/curiously/types"
	"gitub.com/matheus-hrm/curiously/utils"
)

type Handler struct {
	store types.QuestionStorage
}

func NewHandler(store types.QuestionStorage) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.POST("/question", h.handleCreateQuestion)
	router.GET("/question/:id", h.handleGetQuestion)
	router.GET("/question", h.handleGetQuestions)
}

func (h *Handler) handleCreateQuestion(c *gin.Context) {
	var payload types.CreateQuestionPayload

	if err := utils.ParseJson(c.Request, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	question, err := h.store.CreateQuestion(payload, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, question)
}

// TODO: FIX isAnonymous texto to bool conversion from db
func (h *Handler) handleGetQuestion(c *gin.Context) {
	id, err := utils.GetIDFromParam(c, "id")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	question, err := h.store.GetQuestionByID(id, c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, question)
}

// TODO: FIX isAnonymous texto to bool conversion from db
func (h *Handler) handleGetQuestions(c *gin.Context) {
	questions, err := h.store.GetQuestions(c)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(c, http.StatusOK, questions)
}

