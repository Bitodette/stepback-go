package handler

import (
	"net/http"

	"stepback-golang/internal/dto"
	"stepback-golang/internal/service"
	"stepback-golang/internal/utils"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if !utils.ParseBody(r, &req) {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if errs := utils.ValidateStruct(req); errs != nil {
		utils.ValidationError(w, errs)
		return
	}

	user, err := h.service.Register(r.Context(), &req)
	if err != nil {
		utils.Error(w, http.StatusConflict, err.Error())
		return
	}

	utils.Created(w, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if !utils.ParseBody(r, &req) {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if errs := utils.ValidateStruct(req); errs != nil {
		utils.ValidationError(w, errs)
		return
	}

	tokens, err := h.service.Login(r.Context(), &req)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(w, tokens)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshRequest
	if !utils.ParseBody(r, &req) {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if errs := utils.ValidateStruct(req); errs != nil {
		utils.ValidationError(w, errs)
		return
	}

	tokens, err := h.service.RefreshToken(r.Context(), &req)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(w, tokens)
}
