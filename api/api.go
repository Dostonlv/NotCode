package api

import (
	"github.com/Dostonlv/NotCode.git/api/handlers"
	"github.com/Dostonlv/NotCode.git/config"
	"github.com/Dostonlv/NotCode.git/storage"
	"github.com/Dostonlv/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewAPI(r *gin.Engine, cfg *config.Config, notcode storage.StorageI, logger logger.LoggerI) {
	handler := handlers.NewHandler(cfg, notcode, logger)
	r.POST("/api/v1/compiler", handler.Compiler)
}
