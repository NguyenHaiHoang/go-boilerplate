package transporter

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type restfulApi struct {
	Router *gin.Engine
	Server *http.Server
}

func RestTransporter(host string, port int, setupRoute func(engine *gin.Engine) error) (Transporter, error) {
	rest := &restfulApi{}
	rest.Router = gin.Default()
	rest.Server = &http.Server{
		Handler: rest.Router,
		Addr:    fmt.Sprintf("%s:%d", host, port),
	}
	return rest, setupRoute(rest.Router)
}

func (r restfulApi) Name() string {
	return "rest"
}

func (r restfulApi) Start() error {
	return r.Server.ListenAndServe()
}

func (r restfulApi) Stop(ctx context.Context) {
	_ = r.Server.Shutdown(ctx)
}
