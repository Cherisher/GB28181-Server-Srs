package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Cherisher/GB28181-Server-Srs/config"
	"github.com/Cherisher/GB28181-Server-Srs/route"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfigJson("conf/app.json"); err != nil {
		log.Fatalf("InitConfigJson: %s\n", err)
	}

	conf := config.Config()

	gin.SetMode(gin.DebugMode)

	r := route.SetupRouter()
	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(conf.Self.Port), 10),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
