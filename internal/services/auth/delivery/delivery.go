package delivery

import (
	"context"

	"snakealive/m/internal/services/auth/models"
	"snakealive/m/internal/services/auth/usecase"
	"snakealive/m/pkg/error_adapter"
	auth_service "snakealive/m/pkg/services/auth"

	"github.com/golang/protobuf/ptypes/empty"
)

type authDelivery struct {
	authUsecase  usecase.AuthUseCase
	errorAdapter error_adapter.ErrorAdapter
	auth_service.UnimplementedAuthServiceServer
}

func (a *authDelivery) ValidateSession(ctx context.Context, session *auth_service.Session) (*auth_service.ValidateSessionResponse, error) {
	response, err := a.authUsecase.ValidateSession(ctx, session.Cookie)
	if err != nil {
		return &auth_service.ValidateSessionResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.ValidateSessionResponse{UserId: response}, nil
}

func (a *authDelivery) LogoutUser(ctx context.Context, session *auth_service.Session) (*empty.Empty, error) {
	return &empty.Empty{}, a.authUsecase.LogoutUser(ctx, session.Cookie)
}

func (a *authDelivery) LoginUser(ctx context.Context, request *auth_service.LoginRequest) (*auth_service.LoginResponse, error) {
	response, err := a.authUsecase.LoginUser(ctx, &models.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &auth_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.LoginResponse{
		Cookie: response.Cookie,
		Token:  response.Token,
	}, nil
}

func (a *authDelivery) RegisterUser(ctx context.Context, request *auth_service.RegisterRequest) (*auth_service.LoginResponse, error) {
	response, err := a.authUsecase.RegisterUser(ctx, &models.User{
		Name:     request.Name,
		Surname:  request.Surname,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &auth_service.LoginResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.LoginResponse{
		Cookie: response.Cookie,
		Token:  response.Token,
	}, nil
}

func (a *authDelivery) GetUser(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.GetUserResponse, error) {
	response, err := a.authUsecase.GetUser(ctx, request.Id)
	if err != nil {
		return &auth_service.GetUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetUserResponse{
		Name:        response.Name,
		Surname:     response.Surname,
		Email:       response.Email,
		Image:       response.Image,
		Description: response.Description,
	}, nil
}

func (a *authDelivery) UpdateUser(ctx context.Context, request *auth_service.UpdateUserRequest) (*auth_service.GetUserResponse, error) {
	response, err := a.authUsecase.UpdateUser(ctx, &models.User{
		ID:          request.Id,
		Name:        request.Name,
		Surname:     request.Surname,
		Email:       request.Email,
		Password:    request.Password,
		Image:       request.Image,
		Description: request.Description,
	})
	if err != nil {
		return &auth_service.GetUserResponse{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.GetUserResponse{
		Name:        response.Name,
		Surname:     response.Surname,
		Email:       response.Email,
		Image:       response.Image,
		Description: response.Description,
	}, nil
}

func (a *authDelivery) GetUserInfo(ctx context.Context, request *auth_service.GetUserRequest) (*auth_service.UserInfo, error) {
	responce, err := a.authUsecase.GetUserInfo(ctx, int(request.Id))
	if err != nil {
		return &auth_service.UserInfo{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.UserInfo{
		UserId:  responce.ID,
		Name:    responce.Name,
		Surname: responce.Surname,
		Image:   responce.Image,
	}, nil
}

func (a *authDelivery) GetUserByEmail(ctx context.Context, request *auth_service.UserEmailRequest) (*auth_service.UserId, error) {
	responce, err := a.authUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return &auth_service.UserId{}, a.errorAdapter.AdaptError(err)
	}

	return &auth_service.UserId{Id: responce.ID}, nil
}

func NewAuthDelivery(
	authUsecase usecase.AuthUseCase,
	errorAdapter error_adapter.ErrorAdapter,
) auth_service.AuthServiceServer {
	return &authDelivery{
		authUsecase:  authUsecase,
		errorAdapter: errorAdapter,
	}
}
