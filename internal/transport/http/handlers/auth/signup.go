package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

func (h *Handler) V1AuthSignupPost(ctx context.Context, req *api.V1AuthSignupPostReq) (api.V1AuthSignupPostRes, error) {
	resp, err := h.svc.SignUp(ctx, auth.SignUpParams{
		Email:        req.GetEmail(),
		Phone:        req.GetPhone(),
		Password:     req.GetPassword(),
		FirstName:    req.GetFirstName(),
		LastName:     req.GetLastName(),
		MiddleName:   req.GetMiddleName(),
		AvatarURL:    req.GetAvatarURL().Value,
		INN:          req.GetInn(),
		IsContractor: req.GetIsContractor(),
	})
	if err != nil {
		return nil, fmt.Errorf("signup: %w", err)
	}

	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    resp.Session.ID,
		Path:     "/",
		Expires:  resp.Session.ExpiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return &api.V1AuthSignupPostCreatedHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AuthSignupPostCreated{
			AccessToken: resp.AccessToken,
		},
	}, nil
}
