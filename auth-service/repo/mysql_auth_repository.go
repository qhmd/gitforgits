package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/qhmd/gitforgits/auth-service/config"
	"github.com/qhmd/gitforgits/auth-service/model"
	"github.com/qhmd/gitforgits/shared/models"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

// UpdateMe implements auth.AuthRepository.

func NewMySQLAuthRepository(db *gorm.DB) model.AuthRepository {
	return &mysqlAuthRepository{db: db}
}

func (m *mysqlAuthRepository) UpdateMe(ctx context.Context, user *models.Auth) (*models.Auth, error) {
	fmt.Print("isi semuanya ", user)
	if err := m.db.WithContext(ctx).Model(user).Where("id = ?", &user.ID).Updates(user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, config.ErrUserExists
			}
		}
		return nil, err
	}
	return user, nil
}

// DeleteUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) DeleteUser(ctx context.Context, id uint) error {
	if err := m.db.WithContext(ctx).Delete(&models.Auth{}, id); err != nil {
		return err.Error
	}
	return nil
}

// ListUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) ListUser(ctx context.Context) ([]*models.Auth, error) {
	panic("unimplemented")
}

// LogoutUser implements auth.AuthRepository.
func (m *mysqlAuthRepository) LogoutUser(ctx context.Context, token string) error {
	panic("unimplemented")
}

func (m *mysqlAuthRepository) RegisterUser(ctx context.Context, auth *models.Auth) error {
	return m.db.WithContext(ctx).Create(auth).Error
}

func (m *mysqlAuthRepository) FindByEmail(ctx context.Context, email string) (*models.Auth, error) {
	var user models.Auth
	err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlAuthRepository) GetUserByID(ctx context.Context, id uint) (*models.Auth, error) {
	var u models.Auth
	err := m.db.WithContext(ctx).First(&u, id).Error
	return &u, err
}
