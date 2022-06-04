package service

import (
	"context"

	coreclient "github.com/fajarsep12zx/go-micro-starter/core/client"
	servicemodel "github.com/fajarsep12zx/go-micro-starter/core/proto"
)

// GetValidateTokenSession ...
func GetValidateTokenSession(token string) (*servicemodel.AuthData, error) {
	service := coreclient.GetAuthService()
	response, err := service.ValidateAccessTokenSession(
		context.Background(),
		&servicemodel.AuthRequest{
			AccessToken: token,
		},
	)

	if err != nil {
		return nil, err
	}

	return response.AuthToken, nil
}
