package server_handler

import (
	"fmt"
	"lawSite/internal/mail"
	"lawSite/internal/tools"

	"lawSite/internal/config"
	"lawSite/internal/database"
)

var ServerClient *Server

type Server struct {
	db            *database.Client
	mailClient    *mail.Client
	linkGenerator *tools.Generator
}

func NewServer() *Server {
	return &Server{}
}

func init() {
	if err := config.ReadConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read config: %s", err))
	}

	db, err := database.NewClient(config.Config)
	if err != nil {
		panic(fmt.Sprintf("Failed connect to database: %s", err))
	}

	generator := tools.NewGenerator()
	mailClient := mail.NewClient(config.Config)

	ServerClient = &Server{
		db:            db,
		mailClient:    mailClient,
		linkGenerator: generator,
	}
}
