package resolver

import (
	"context"

	"github.com/fajarsep12zx/go-micro-starter/api/login/model"
	coreclient "github.com/fajarsep12zx/go-micro-starter/core/client"
	servicemodel "github.com/fajarsep12zx/go-micro-starter/core/proto"
	"github.com/fajarsep12zx/go-micro-starter/core/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// DoLogin ...
func DoLogin(username, password string) (result *model.LoginData, err error) {
	service := coreclient.GetUserService()
	response, err := service.Login(
		context.Background(),
		&servicemodel.UserRequest{
			User: &servicemodel.UserData{
				Username: username,
				Password: password,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.LoginData)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}

// DoRefreshToken ...
func DoRefreshToken(refreshToken string) (result *model.LoginData, err error) {
	service := coreclient.GetUserService()
	response, err := service.RefreshToken(
		context.Background(),
		&servicemodel.UserRequest{
			RefreshToken: refreshToken,
		},
	)

	if err != nil {
		return nil, err
	}

	result = new(model.LoginData)
	err = utils.CopyObject(response.Info, &result)
	return result, err
}
