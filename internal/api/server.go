package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zanovru/gin-rest-api/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(router *gin.Engine, configs *config.Configs) error {
	s.httpServer = &http.Server{
		Addr:           ":" + configs.Server.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(configs.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(configs.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx *gin.Context) error {
	return s.httpServer.Shutdown(ctx)
}
