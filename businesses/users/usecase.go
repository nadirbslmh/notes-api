package users

import (
	"context"
	"notes-api/app/middlewares"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.JWTConfig
}

func NewUserUseCase(ur Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *userUsecase) Register(ctx context.Context, userDomain *Domain) (Domain, error) {
	return uu.userRepository.Register(ctx, userDomain)
}

func (uu *userUsecase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	user, err := uu.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", err
	}

	token, err := uu.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
