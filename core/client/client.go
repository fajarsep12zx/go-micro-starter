package client

import (
	"sync"
	"time"

	configFile "github.com/fajarsep12zx/go-micro-starter/core/config"
	service "github.com/fajarsep12zx/go-micro-starter/core/proto"
	"github.com/fajarsep12zx/go-micro-starter/core/server"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func newGRPCClient() client.Client {
	options := getClientOptions()
	return grpc.NewClient(options...)
}

func getClientOptions() []client.Option {
	config := configFile.GetApplicationConfig()
	options := []client.Option{
		client.RequestTimeout(config.GRPCTimeout * time.Second),
	}

	if server.IsUseKubernetes() {
		options = append(options, client.Registry(kubernetes.NewRegistry()))
	}

	return options
}

var (
	authOnce    sync.Once
	authService service.AuthService
)

// GetAuthService ...
func GetAuthService() service.AuthService {
	authOnce.Do(func() {
		authService = service.NewAuthService(server.AuthService, newGRPCClient())
	})

	return authService
}

var (
	userOnce    sync.Once
	userService service.UserService
)

// GetUserService ...
func GetUserService() service.UserService {
	userOnce.Do(func() {
		userService = service.NewUserService(server.UserService, newGRPCClient())
	})

	return userService
}
