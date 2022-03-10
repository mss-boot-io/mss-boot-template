package main

import (
	"context"
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/pkg/config"

	"{{.service}}/cfg"
	"{{.service}}/router"
)

// @title {{.service}} API
// @version {{.version}}
// @description {{.service}}接口文档

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @host localhost:{{.port}}
// @BasePath /{{.service}}
func main() {
	c := &cfg.Config{}
	err := config.Init(flag.Lookup("c").Value.String(), c)
	if err != nil {
		log.Printf("cfg init failed, %s\n", err.Error())
		return
	}
	ctx := context.Background()

	mgr := server.New()

	r := gin.Default()
	router.Init(r.Group("/{{.service}}"))

	c.Init(mgr, r)

	log.Println("starting {{.service}} manage")

	err = mgr.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
