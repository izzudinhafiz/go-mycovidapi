package server

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	docs "github.com/izzudinhafiz/go-mycovidapi/docs"
)

// @title MYCovidAPI
// @version 	1.0
// @description RESTful API for Covid dataset provided by Ministry of Health Malaysia

// @host		mycovidapi.izzudinhafiz.com


type Server struct {
	Router *gin.Engine
	DB *gorm.DB
}

func New() *Server {
	s := Server{}
	s.Initialize()
	return &s
}

func (s *Server) Initialize() {
	var dsn = fmt.Sprintf(
		"host=%v user=%v password=%v port=%v database=%v sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	s.Router = gin.Default()
	s.DB = db


	s.RegisterPath("/country/cases", s.GetCountryCases)
	s.RegisterPath("/country/cases/clusters", s.GetCountryCasesClusters)
	s.RegisterPath("/country/cases/age", s.GetCountryCasesAgeCategory)
	s.RegisterPath("/state/cases", s.GetStateCases)
	s.RegisterPath("/state/cases/age", s.GetStateCasesAgeCategory)

	s.RegisterPath("/country/deaths", s.GetCountryDeaths)
	s.RegisterPath("/state/deaths", s.GetStateDeaths)

	s.RegisterPath("/state/icu", s.GetStateICU)
	s.RegisterPath("/state/hospital", s.GetStateHospital)
	s.RegisterPath("/state/pkrc", s.GetStateQuarantineCentre)

	// TODO: Need to add these in
	s.RegisterPath("/clusters/info", s.GetClustersInfo)
	s.RegisterPath("/clusters/category", s.GetClustersInfo)
	s.RegisterPath("/clusters/", s.GetClusterList)
	s.RegisterPath("/clusters/active", s.GetActiveClusters)

	s.Router.StaticFile("/", "./public/index.html")
	s.Router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) RegisterPath(path string, f func(c *gin.Context), basePath ...string) {
	base := "/api/v1"
	docs.SwaggerInfo.BasePath = base
	if len(basePath) > 0 {
		base = basePath[0]
	}
	s.Router.GET(base + path, f)
}

func (s *Server) Run() {
	s.Initialize()
	envPort, isSet := os.LookupEnv("SERVER_PORT")
	if !isSet {
		envPort = "8000"
	}
	port := ":" + envPort
	log.Fatal(http.ListenAndServe(port, s.Router))
}

