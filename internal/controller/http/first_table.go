package http

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api/internal/domain/entity"
)

const (
	URL = "/getall"
)

type FirstUsecase interface {
	GetAllService() []entity.FirstTable
}

type FirstHandler struct {
	firstUsecase FirstUsecase
}

func NewFirstHandler(firstUsecase FirstUsecase) *FirstHandler {
	return &FirstHandler{firstUsecase: firstUsecase}
}

func (h *FirstHandler) Register(router *httprouter.Router) {
	router.GET(URL, h.GetAllService)
}

func (h *FirstHandler) GetAllService(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users := h.firstUsecase.GetAllService()
	fmt.Println(users)
}
