package service

import (
	"context"
	"errors"

	"stepback-golang/internal/dto"
	"stepback-golang/internal/model"
	"stepback-golang/internal/repository"
	"stepback-golang/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, req *dto.RegisterRequest) (*model.User, error) {
	existing, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &model.User{
		Email:        req.Email,
		PasswordHash: string(hash),
		Name:         req.Name,
		PhoneNumber:  req.PhoneNumber,
		Role:         "customer",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.TokenResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.TokenResponse{
		Access:  accessToken,
		Refresh: refreshToken,
		User:    user,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *dto.RefreshRequest) (*dto.TokenResponse, error) {
	claims, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	user, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.TokenResponse{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}
