package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
}

type Usecase interface {
	Register(ctx context.Context, userDomain *Domain) Domain
	Login(ctx context.Context, userDomain *Domain) string
}

type Repository interface {
	Register(ctx context.Context, userDomain *Domain) Domain
	GetByEmail(ctx context.Context, userDomain *Domain) Domain
}
