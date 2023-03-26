package users

import (
	"context"
	"fmt"
	"notes-api/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(ctx context.Context, userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) GetByEmail(ctx context.Context, userDomain *users.Domain) users.Domain {
	var user User

	ur.conn.WithContext(ctx).First(&user, "email = ?", userDomain.Email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return users.Domain{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		fmt.Println("password failed!")
		return users.Domain{}
	}

	return user.ToDomain()
}
