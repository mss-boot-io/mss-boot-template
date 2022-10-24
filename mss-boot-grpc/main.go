package main

import (
	"context"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/grpc"

	"github.com/WhiteMatrixTech/matrix-cloud-monorepo/{{.serviceName}}/cfg"
	"github.com/WhiteMatrixTech/matrix-cloud-monorepo/{{.serviceName}}/handlers"
	pb "github.com/WhiteMatrixTech/matrix-cloud-monorepo/{{.serviceName}}/proto/v1"
)

func main() {
	ctx := context.Background()

	cfg.Cfg.Init(func(srv *grpc.Server) {
		pb.RegisterHelloworldServer(srv.Server(), handlers.New("{{.serviceName}}"))
	})

	log.Info("starting {{.serviceName}} manage")

	err := server.Manage.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
