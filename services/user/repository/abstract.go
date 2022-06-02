package repository

import (
	cache "github.com/patrickmn/go-cache"
	core "zebrax.id/product/dmaa/core/proto"
	"zebrax.id/product/dmaa/core/utils"
	"zebrax.id/product/dmaa/services/user/repository/postgres"
)

// Repository interface
type Repository interface {
	GetUserByIDQuery(userID int) ([][]string, error)
	GetUserByUsernameQuery(username string) ([][]string, error)
	GetUserDataByIDQuery(userID int) ([][]string, error)
	GetUserDataByUsernameQuery(username string) ([][]string, error)
	GetUserPermissionQuery(menu string) ([][]string, error)
	GetUserPermissionBulkQuery(menu []string) ([][]string, error)
	GetUserRoleQuery(name string) ([][]string, error)
	UpdatePasswordQuery(userID int, password string) error
	CreateUserQuery(user *core.UserData) (int, error)
	CreateUserPermissionQuery(permission *core.UserPermission) error
	CreateUserRoleQuery(role *core.UserRole) error
}

// AbstractRepository ..
type AbstractRepository struct {
	Repository
	Cache *cache.Cache
}

// NewRepository ...
func NewRepository() *AbstractRepository {
	newRepo := &AbstractRepository{
		Cache: utils.GetCacheHandler(),
	}

	if utils.GetDatasourceInfo() == utils.Postgres {
		newRepo.Repository = &postgres.Repository{
			Connection: utils.GetPostgresHandler(),
		}
	} else {
		newRepo.Repository = nil
	}

	return newRepo
}
