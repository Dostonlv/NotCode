package main

import (
	"fmt"
	"github.com/Dostonlv/NotCode.git/api"
	"github.com/Dostonlv/NotCode.git/config"
	"github.com/Dostonlv/NotCode.git/storage/postgres"
	"github.com/Dostonlv/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// call logger
	r.Use(gin.Recovery(), gin.Logger())
	cfg := config.Load()
	//
	var loggerLevel = new(string)
	*loggerLevel = logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)

	}

	log := logger.NewLogger("app", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	notcode, err := postgres.NewConnectPostgresql(&cfg)

	api.NewAPI(r, &cfg, notcode, log)

	fmt.Println("Server running on port", cfg.ServerHost+cfg.ServerPort)
	err = r.Run(cfg.ServerHost + cfg.ServerPort)
	if err != nil {
		log.Panic("Error listening server: ", logger.Error(err))
		return
	}

	//var s = models.Req{
	//	Version:  "3.10",
	//	Language: "python",
	//	Code:     "input()\nprint(inp[::-1])",
	//	Cases:    "hello world",
	//}
	//
	//d, err := postgres.NewConnectPostgresql(&cfg)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//b, err := d.Compiler().Compile(s)
	//if err != nil {
	//	logger.Error(err)
	//}
	//
	//fmt.Println(b.GetOutput())

}
