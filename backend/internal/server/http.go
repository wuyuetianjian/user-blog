package server

import (
	"net/http"

	"qizhan/backend/internal/service"

	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(addr string, about *service.AboutService) *kratoshttp.Server {
	srv := kratoshttp.NewServer(
		kratoshttp.Address(addr),
	)

	srv.HandleFunc("/api/about", func(w http.ResponseWriter, r *http.Request) {
		data, err := about.GetAbout()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"error":"internal error"}`))
			return
		}
		_ = srv.JSON(w, http.StatusOK, data)
	})

	return srv
}
