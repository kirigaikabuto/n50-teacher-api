package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"github.com/kirigaikabuto/n50-teacher-api/users"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configName              = "main"
	configPath              = "/config/"
	version                 = "0.0.1"
	s3endpoint              = ""
	s3bucket                = ""
	s3accessKey             = ""
	s3secretKey             = ""
	s3uploadedFilesBasePath = ""
	s3region                = ""
	postgresUser            = ""
	postgresPassword        = ""
	postgresDatabaseName    = ""
	postgresHost            = ""
	postgresPort            = 5432
	postgresParams          = ""
	port                    = "8080"
	flags                   = []cli.Flag{
		&cli.StringFlag{
			Name:        "config, c",
			Usage:       "path to .env config file",
			Value:       "main",
			Destination: &configName,
		},
	}
)

func parseEnvFile() {
	filepath, err := os.Getwd()
	if err != nil {
		panic("main, get rootDir error" + err.Error())
		return
	}
	viper.AddConfigPath(filepath + configPath)
	viper.SetConfigName(configName)
	err = viper.ReadInConfig()
	if err != nil {
		panic("main, fatal error while reading config file: " + err.Error())
		return
	}
	postgresUser = viper.GetString("db.primary.user")
	postgresPassword = viper.GetString("db.primary.pass")
	postgresDatabaseName = viper.GetString("db.primary.name")
	postgresParams = viper.GetString("db.primary.param")
	postgresPort = viper.GetInt("db.primary.port")
	postgresHost = viper.GetString("db.primary.host")
	s3endpoint = viper.GetString("s3.primary.s3endpoint")
	s3bucket = viper.GetString("s3.primary.s3bucket")
	s3accessKey = viper.GetString("s3.primary.s3accessKey")
	s3secretKey = viper.GetString("s3.primary.s3secretKey")
	s3uploadedFilesBasePath = viper.GetString("s3.primary.s3uploadedFilesBasePath")
	s3region = viper.GetString("s3.primary.s3region")
}

func run(c *cli.Context) error {
	parseEnvFile()
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	gin.SetMode(gin.ReleaseMode)
	cfg := common.PostgresConfig{
		Host:     postgresHost,
		Port:     postgresPort,
		User:     postgresUser,
		Password: postgresPassword,
		Database: postgresDatabaseName,
		Params:   postgresParams,
	}
	//applications
	_, err := common.NewS3Uploader(
		s3endpoint,
		s3accessKey,
		s3secretKey,
		s3bucket,
		s3uploadedFilesBasePath,
		s3region)
	if err != nil {
		return err
	}
	usersPostgreStore, err := users.NewPostgresUsersStore(cfg)
	if err != nil {
		return err
	}
	usersService := users.NewUserService(usersPostgreStore)
	fmt.Println(usersService)
	r := gin.Default()
	//r.Use(apiKewMdw.MakeCorsMiddleware())
	authGroup := r.Group("/users")
	{
		authGroup.POST("/")
	}
	log.Info().Msg("app is running on port:" + port)
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting.")
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "n50-teacher-api"
	app.Description = ""
	app.Usage = "n50-teacher-api"
	app.UsageText = "n50-teacher-api"
	app.Version = version
	app.Flags = flags
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Err(err)
	}
}
