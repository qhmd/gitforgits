package repo

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/qhmd/gitforgits/shared/dto"
	"github.com/qhmd/gitforgits/shared/models"
	"github.com/qhmd/gitforgits/users-service/config"
	"github.com/qhmd/gitforgits/users-service/model"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

// DeleteUser implements user.UserRepository.
func (m *mysqlUserRepository) DeleteUser(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).Delete(&models.Auth{}, id).Error
}

// FindByEmail implements user.UserRepository.
func (m *mysqlUserRepository) FindByEmail(ctx context.Context, email string) (*models.Auth, error) {
	var user *models.Auth
	if err := m.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser implements user.UserRepository.
func (m *mysqlUserRepository) GetUser(ctx context.Context, id int) (*models.Auth, error) {
	var user *models.Auth
	if err := m.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// ListUser implements user.UserRepository.
func (m *mysqlUserRepository) ListUser(ctx context.Context) ([]*models.Auth, error) {
	var users []*models.Auth
	if err := m.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser implements user.UserRepository.
func (m *mysqlUserRepository) UpdateUser(ctx context.Context, users *dto.UserResponse, id int) (*dto.UserResponse, error) {
	if err := m.db.WithContext(ctx).Model(&models.Auth{}).Where("id = ?", id).Updates(&users).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil, config.ErrUserExists
			}
		}
		return nil, err
	}
	fmt.Print("isi dari user di mysql", users)
	return users, nil
}

func NewUserMySqlRepo(db *gorm.DB) model.UserRepository {
	return &mysqlUserRepository{db: db}
}
