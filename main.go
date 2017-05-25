package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/yursan9/tulis/pkg/server"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	app := server.New()
	go func() {
		log.Println("Listening to port http://127.0.0.1" + app.Addr)
		log.Fatal(app.ListenAndServe())
	}()
	<-stop

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	log.Println("Shutting down the server...")
	app.Shutdown(ctx)
	log.Println("Server stopped.")
}
