package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	hHealth "bitbucket.org/Amartha/go-dlq-retrier/internal/http/handler/health"
	"github.com/labstack/echo/v4"
)

func TestGet_Handle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		get     *hHealth.Get
		wantErr bool
	}{
		{
			name:    "success",
			get:     hHealth.NewGet(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				req := httptest.NewRequest(http.MethodGet, "/health", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				w := httptest.NewRecorder()

				e := echo.New()
				ctx := e.NewContext(req, w)

				get := tt.get
				if err := get.Handle(ctx); (err != nil) != tt.wantErr {
					t.Errorf("Get.Handle() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}
