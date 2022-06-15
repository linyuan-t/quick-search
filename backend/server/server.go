package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/linyuan-t/quick-search/backend/config"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Run is used to run server.
func Run(filename string) {

	var config config.Config
	config.FileName = filename

	s, err := initServer(&config)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	srv := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: initRouter(s),
	}

	go func() {
		// service connections
		log.Println("Run server successfully")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -3 is syscall.SIGQUIT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
