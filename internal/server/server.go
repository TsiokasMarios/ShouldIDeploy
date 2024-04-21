package server

import (
	"ShouldIDeploy/cmd/utils"
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {

	utils.LoadBundles()

	//port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: 8080,
	}
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	println("Server is running on port: ", NewServer.port)

	return server
}
