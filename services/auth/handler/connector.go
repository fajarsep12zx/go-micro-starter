package handler

import (
	"sync"

	"github.com/fajarsep12zx/go-micro-starter/core/cache"
	"github.com/fajarsep12zx/go-micro-starter/services/auth/repository"
	"github.com/fajarsep12zx/go-micro-starter/services/auth/usecase"
)

var uc *usecase.UseCase
var oneUc sync.Once

// GetUsecase ...
func GetUsecase() *usecase.UseCase {
	oneUc.Do(func() {
		uc = usecase.NewUseCase(
			getRepository(),
			getSession(),
		)
	})

	return uc
}

var (
	repo    *repository.AbstractRepository
	oneRepo sync.Once
)

func getRepository() *repository.AbstractRepository {
	oneRepo.Do(func() {
		repo = repository.NewRepository()
	})

	return repo
}

var (
	sess    cache.ICache
	oneSess sync.Once
)

func getSession() cache.ICache {
	oneSess.Do(func() {
		sess = cache.NewSession()
	})

	return sess
}
