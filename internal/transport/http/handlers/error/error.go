package error

import (
	"context"
	"net/http"
)

func (h *Handler) HandleError(ctx context.Context, w http.ResponseWriter, _ *http.Request, err error) {
	_, _ = w.Write([]byte(err.Error()))
}
