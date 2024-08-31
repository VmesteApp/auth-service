package usecase

import (
	"context"

	"github.com/VmesteApp/auth-service/internal/entity"
)

type (
	User interface {
		CreateAccount(ctx context.Context, email, password string) error
		Login(ctx context.Context, email, password string) (string, error)
		VkLogin(ctx context.Context, userAccessToken string) (string, error)
	}
	UserRepo interface {
		SaveUser(ctx context.Context, email string, hassPash []byte) error
		User(ctx context.Context, email string) (*entity.User, error)
		SaveSocialUser(ctx context.Context, provider, providerID string) (*entity.User, error)
		SocialUser(ctx context.Context, provider, providerID string) (*entity.User, error)
	}
	VkWebApi interface {
		ValidateUserAccessToken(userAccessToken string) (*entity.VkTokenInfo, error)
	}
)
