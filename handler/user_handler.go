package handler

import (
	"library-api/common/dto"
	"library-api/model"
	"library-api/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	req := dto.GetRequestData[model.User](r)

	err := h.service.CreateNewUser(req)
	if err != nil {
		res := dto.NewResponse(err.Error(), false)
		res.Send(w)
		return
	}

	res := dto.NewResponse("created", true)
	res.Send(w)
}

func (h *UserHandler) NewSession(w http.ResponseWriter, r *http.Request) {
	req := dto.GetRequestData[model.User](r)

	err := h.service.CreateNewSession(req)
	if err != nil {
		res := dto.NewResponse(err.Error(), false)
		res.Send(w)
		return
	}

	res := dto.NewResponse("logged in", true)
	res.Send(w)
}

func (h *UserHandler) ClearSession(w http.ResponseWriter, r *http.Request) {
	err := h.service.ClearCurrSession()
	if err != nil {
		res := dto.NewResponse(err.Error(), false)
		res.Send(w)
		return
	}

	res := dto.NewResponse("logged out", true)
	res.Send(w)
}
