package usecase

import (
	"github.com/fajarsep12zx/go-micro-starter/services/user/client"
	"github.com/fajarsep12zx/go-micro-starter/services/user/repository"
)

// UseCase ...
type UseCase struct {
	repository *repository.AbstractRepository
	client     *client.Client
}

// NewUseCase ...
func NewUseCase(repo *repository.AbstractRepository, cl *client.Client) *UseCase {
	return &UseCase{
		repository: repo,
		client:     cl,
	}
}
