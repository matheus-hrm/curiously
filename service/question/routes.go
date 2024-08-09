package question

import "gitub.com/matheus-hrm/curiously/types"

type Handler struct {
	store types.QuestionStorage
}

func NewHandler(store types.QuestionStorage) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {

}