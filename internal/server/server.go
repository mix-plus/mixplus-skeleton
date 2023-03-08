package server

import (
	"api1/internal/service"
	"api1/internal/svc"
	"net/http"

	"github.com/google/wire"
	"github.com/mix-plus/go-mixplus/mrpc"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

type AppServer struct {
	SvcCtx *svc.ServiceContext
	Hs     *http.Server
	gs     *mrpc.RpcServer

	HelloService *service.HelloService
}

func NewApp(svcCtx *svc.ServiceContext, helloService *service.HelloService, hs *http.Server, gs *mrpc.RpcServer) (*AppServer, error) {
	return &AppServer{
		SvcCtx:       svcCtx,
		HelloService: helloService,
		Hs:           hs,
		gs:           gs,
	}, nil
}

func (a *AppServer) Run() {

	go func() {
		a.Hs.ListenAndServe()
	}()

	a.gs.Start()

	defer a.gs.Stop()
}
