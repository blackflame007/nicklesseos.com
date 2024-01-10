package layouts

import (
	"context"
)

func GetAuthenticatedUser(ctx context.Context) string {
	user, ok := ctx.Value("user").(string)
	if !ok {
		return "Not OK"
	}
	return user

}

func IsAuthenticated(ctx context.Context) string {
	auth, ok := ctx.Value("authenticated").(string)
	if !ok {
		return "Not OK"
	}
	return auth

}
