package repository

import (
	"context"
	"errors"

	"github.com/qhmd/gitforgits/internal/domain/auth"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

// DeleteUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) DeleteUser(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// ListUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) ListUser(ctx context.Context) ([]*auth.Auth, error) {
	panic("unimplemented")
}

// LogoutUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) LogoutUser(ctx context.Context, token string) error {
	panic("unimplemented")
}

func (m *mysqlAuthRepository) RegisterUser(ctx context.Context, auth *auth.Auth) error {
	return m.db.WithContext(ctx).Create(auth).Error
}

func (m *mysqlAuthRepository) FindByEmail(ctx context.Context, email string) (*auth.Auth, error) {
	var user auth.Auth
	err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlAuthRepository) GetUserByID(ctx context.Context, id uint) (*auth.Auth, error) {
	var u auth.Auth
	err := m.db.WithContext(ctx).First(&u, id).Error
	return &u, err
}

func NewMySQLAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &mysqlAuthRepository{db: db}
}
