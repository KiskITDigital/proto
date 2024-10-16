package auth

import "context"

type SignInParams struct {
}

type SignInResult struct {
}

func SignIn(ctx context.Context, params SignInParams) (SignInResult, error) {
	return SignInResult{}, nil
}
