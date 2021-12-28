package server

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gin-gonic/gin"
	"github.com/izzudinhafiz/go-mycovidapi/models"
)

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
	//TODO: DSN should come from ENV
	var dsn = fmt.Sprintf(
		"host=%v user=%v password=%v port=%v database=%v sslmode=disable",
		os.Getenv("postgres_host"),
		os.Getenv("postgres_user"),
		os.Getenv("postgres_pass"),
		os.Getenv("postgres_port"),
		os.Getenv("postgres_database"),
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
	s.RegisterPath("/clusters/info", s.GetClusters)
	s.RegisterPath("/clusters/category", s.GetClusters)
	s.RegisterPath("/clusters/name", s.GetClusters)

	s.Router.StaticFile("/", "./public/index.html")
}

func (s *Server) RegisterPath(path string, f func(c *gin.Context), basePath ...string) {
	base := "/api/v1"
	if len(basePath) > 0 {
		base = basePath[0]
	}
	s.Router.GET(base + path, f)
}

func (s *Server) Run() {
	s.Initialize()
	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, s.Router)) //TODO: Change port to SERVER_PORT
}

func (s *Server) GetCountryCases(c *gin.Context) {
	respModel := []*models.Cases{}
	s.getCountry(c, respModel)
}

func (s *Server) GetStateCases(c *gin.Context) {
	respModel := []*models.Cases{}
	s.getState(c, respModel)
}

func (s *Server) GetCountryCasesClusters(c *gin.Context) {
	respModel := []*models.CasesCluster{}
	s.getCountry(c, respModel)
}

func (s *Server) GetCountryCasesAgeCategory(c *gin.Context) {
	respModel := []*models.CasesAgeCategory{}
	s.getCountry(c, respModel)
}

func (s *Server) GetStateCasesAgeCategory(c *gin.Context) {
	respModel := []*models.CasesAgeCategory{}
	s.getState(c, respModel)
}

func (s *Server) GetCountryDeaths(c *gin.Context) {
	respModel := []*models.Deaths{}
	s.getCountry(c, respModel)
}

func (s *Server) GetStateDeaths(c *gin.Context) {
	respModel := []*models.Deaths{}
	s.getState(c, respModel)
}

func (s *Server) GetStateICU(c *gin.Context) {
	respModel := []*models.ICU{}
	s.getState(c, respModel)
}

func (s *Server) GetStateHospital(c *gin.Context) {
	respModel := []*models.Hospital{}
	s.getState(c, respModel)
}

func (s *Server) GetStateQuarantineCentre(c *gin.Context) {
	respModel := []*models.QuarantineCentre{}
	s.getState(c, respModel)
}

//TODO: this prolly has special handling
func (s *Server) GetClusters(c *gin.Context) {
	respModel := []*models.Clusters{}
	s.getCountry(c, respModel)
}