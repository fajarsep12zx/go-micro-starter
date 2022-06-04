package usecase

import (
	"github.com/fajarsep12zx/go-micro-starter/core/cache"
	"github.com/fajarsep12zx/go-micro-starter/services/auth/repository"
)

// UseCase ...
type UseCase struct {
	repository *repository.AbstractRepository
	session    cache.ICache
}

// NewUseCase ...
func NewUseCase(
	repo *repository.AbstractRepository,
	sess cache.ICache,
) *UseCase {
	return &UseCase{
		repository: repo,
		session:    sess,
	}
}
