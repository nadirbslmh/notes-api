package users

import "context"

type userUsecase struct {
	userRepository Repository
}

func NewUserUseCase(ur Repository) Usecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) Register(ctx context.Context, userDomain *Domain) (Domain, error) {
	//TODO: implement this
	return Domain{}, nil
}

func (uu *userUsecase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	//TODO: implement this
	return "belum selesai", nil
}
