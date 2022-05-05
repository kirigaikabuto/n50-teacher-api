package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kirigaikabuto/n50-teacher-api/auth"
	"github.com/kirigaikabuto/n50-teacher-api/common"
	"github.com/kirigaikabuto/n50-teacher-api/groups"
	"github.com/kirigaikabuto/n50-teacher-api/lessons"
	"github.com/kirigaikabuto/n50-teacher-api/subjects"
	"github.com/kirigaikabuto/n50-teacher-api/users"
	setdata_common "github.com/kirigaikabuto/setdata-common"
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
	port                    = "5000"
	redisHost               = ""
	redisPort               = ""
	adminUsername           = ""
	adminPassword           = ""
	adminEmail              = ""
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
	redisHost = viper.GetString("redis.primary.host")
	redisPort = viper.GetString("redis.primary.port")
	adminUsername = viper.GetString("user.admin.username")
	adminPassword = viper.GetString("user.admin.password")
	adminEmail = viper.GetString("user.admin.email")
}

func run(c *cli.Context) error {
	parseEnvFile()
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
	s3Uploader, err := common.NewS3Uploader(
		s3endpoint,
		s3accessKey,
		s3secretKey,
		s3bucket,
		s3uploadedFilesBasePath,
		s3region)
	if err != nil {
		return err
	}
	authTokenStore, err := auth.NewTokenStore(auth.RedisConfig{
		Host: redisHost,
		Port: redisPort,
	})
	if err != nil {
		return err
	}

	//users store
	usersPostgreStore, err := users.NewPostgresUsersStore(cfg)
	if err != nil {
		return err
	}
	usersService := users.NewUserService(usersPostgreStore, authTokenStore)
	usersHttpEndpoints := users.NewUsersHttpEndpoints(setdata_common.NewCommandHandler(usersService))

	//groups store
	groupPostgreStore, err := groups.NewUserGroupPostgreStore(cfg)
	if err != nil {
		return err
	}
	groupService := groups.NewUserGroupService(groupPostgreStore, usersPostgreStore)
	groupHttpEndpoints := groups.NewUserGroupHttpEndpoints(setdata_common.NewCommandHandler(groupService))
	authMdw := auth.NewMiddleware(authTokenStore)

	//subjects store
	subjectPostgreStore, err := subjects.NewSubjectPostgreStore(cfg)
	if err != nil {
		return err
	}
	subjectService := subjects.NewSubjectService(subjectPostgreStore, usersPostgreStore, groupPostgreStore)
	subjectHttpEndpoints := subjects.NewSubjectsHttpEndpoints(setdata_common.NewCommandHandler(subjectService))

	//lessons store
	lessonsPostgreStore, err := lessons.NewLessonsPostgreStore(cfg)
	if err != nil {
		return err
	}
	lessonService := lessons.NewLessonService(lessonsPostgreStore, subjectPostgreStore, s3Uploader)
	lessonHttpEndpoints := lessons.NewLessonHttpEndpoints(setdata_common.NewCommandHandler(lessonService))

	//create admin
	usersService.CreateUser(&users.CreateUserCommand{
		Username:  adminUsername,
		Password:  adminPassword,
		Email:     adminEmail,
		FirstName: "",
		LastName:  "",
		Type:      common.Admin.ToString(),
	})

	r := gin.Default()
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", usersHttpEndpoints.MakeLoginEndpoint())
		authGroup.POST("/register", authMdw.MakeMiddleware(), usersHttpEndpoints.MakeCreateUserByAdminEndpoint())
	}
	groupGroups := r.Group("/group", authMdw.MakeMiddleware())
	{
		groupGroups.POST("/", groupHttpEndpoints.MakeCreateGroupEndpoint())
		groupGroups.GET("/", groupHttpEndpoints.MakeListGroupEndpoint())
		groupGroups.GET("/id", groupHttpEndpoints.MakeGetGroupByIdEndpoint())
	}
	userGroupGroups := r.Group("/userGroup", authMdw.MakeMiddleware())
	{
		userGroupGroups.POST("/", groupHttpEndpoints.MakeCreateUserGroupEndpoint())
		userGroupGroups.GET("/groupId", groupHttpEndpoints.MakeGetUserGroupByGroupIdEndpoint())
		userGroupGroups.GET("/userId", groupHttpEndpoints.MakeGetUserGroupByUserIdEndpoint())
		userGroupGroups.DELETE("/", groupHttpEndpoints.MakeDeleteUserGroupByIdEndpoint())
	}
	subjectsGroups := r.Group("/subjects", authMdw.MakeMiddleware())
	{
		subjectsGroups.POST("/", subjectHttpEndpoints.MakeCreateSubjectEndpoint())
		subjectsGroups.GET("/id", subjectHttpEndpoints.MakeGetSubjectByIdEndpoint())
		subjectsGroups.GET("/", subjectHttpEndpoints.MakeListSubjectsEndpoint())
	}
	teacherSubjectGroup := r.Group("/teacherSubject", authMdw.MakeMiddleware())
	{
		teacherSubjectGroup.POST("/", subjectHttpEndpoints.MakeCreateTeacherSubjectEndpoint())
		teacherSubjectGroup.GET("/", subjectHttpEndpoints.MakeListTeacherSubjectsEndpoint())
		teacherSubjectGroup.GET("/id", subjectHttpEndpoints.MakeGetTeacherSubjectByIdEndpoint())
		teacherSubjectGroup.GET("/teacherId", subjectHttpEndpoints.MakeGetTeacherSubjectsByTeacherIdEndpoint())
		teacherSubjectGroup.GET("/subjectId", subjectHttpEndpoints.MakeGetTeacherSubjectsBySubjectIdEndpoint())
		teacherSubjectGroup.GET("/token", subjectHttpEndpoints.MakeGetTeacherSubjectsByTokenEndpoint())
	}
	groupSubjectGroup := r.Group("/groupSubject", authMdw.MakeMiddleware())
	{
		groupSubjectGroup.POST("/", subjectHttpEndpoints.MakeCreateGroupSubjectEndpoint())
		groupSubjectGroup.GET("/", subjectHttpEndpoints.MakeListGroupSubjectsEndpoint())
		groupSubjectGroup.GET("/id", subjectHttpEndpoints.MakeGetGroupSubjectsByIdEndpoint())
		groupSubjectGroup.GET("/teacherSubId", subjectHttpEndpoints.MakeGetGroupSubjectByIdTeacherSubEndpoint())
		groupSubjectGroup.GET("/groupId", subjectHttpEndpoints.MakeGetGroupSubjectByGroupIdEndpoint())
		groupSubjectGroup.GET("/ids", subjectHttpEndpoints.MakeGetGroupSubjectByTeacherGroupIdsEndpoint())
	}
	lessonGroup := r.Group("/lesson", authMdw.MakeMiddleware())
	{
		lessonGroup.POST("/", lessonHttpEndpoints.MakeCreateLessonEndpoint())
		lessonGroup.GET("/", lessonHttpEndpoints.MakeGetLessonByIdEndpoint())
		lessonGroup.GET("/groupSubjectId", lessonHttpEndpoints.MakeListLessonByGroupSubjectIdEndpoint())
		lessonGroup.PUT("/", lessonHttpEndpoints.MakeUpdateLessonEndpoint())
		lessonGroup.DELETE("/", lessonHttpEndpoints.MakeDeleteLessonEndpoint())
		lessonGroup.PUT("/file", lessonHttpEndpoints.MakeUploadFileEndpoint())
	}
	log.Info().Msg("app is running on port:" + port)
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
			return
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
		log.Info().Msg(err.Error())
	}
}
