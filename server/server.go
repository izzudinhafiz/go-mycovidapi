package server

import (
	"net/http"
	"path"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gorilla/mux"
	"github.com/izzudinhafiz/go-mycovidapi/models"
)

type Server struct {
	Router *mux.Router
	DB *gorm.DB
}

func New() *Server {
	s := Server{}
	s.Initialize()
	return &s
}

func (s *Server) Initialize() {
	//TODO: DSN should come from ENV
	var dsn = "host=izzudinhafiz.com user=hafizwriter password=pass940203 port=5432 database=test_mycovidapi sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	s.Router = mux.NewRouter()
	s.DB = db

	s.RegisterPath("/country/cases", "GET", s.GetCountryCases)
	s.RegisterPath("/country/cases/clusters", "GET", s.GetCountryCasesClusters)
	s.RegisterPath("/country/cases/age", "GET", s.GetCountryCasesAgeCategory)
	s.RegisterPath("/state/cases", "GET", s.GetStateCases)
	s.RegisterPath("/state/cases/age", "GET", s.GetStateCasesAgeCategory)

	s.RegisterPath("/country/deaths", "GET", s.GetCountryDeaths)
	s.RegisterPath("/state/deaths", "GET", s.GetStateDeaths)

	s.RegisterPath("/state/icu", "GET", s.GetStateICU)
	s.RegisterPath("/state/hospital", "GET", s.GetStateHospital)
	s.RegisterPath("/state/pkrc", "GET", s.GetStateQuarantineCentre)

	// TODO: Need to add these in
	s.RegisterPath("/clusters/info", "GET", s.GetClusters)
	s.RegisterPath("/clusters/category", "GET", s.GetClusters)
	s.RegisterPath("/clusters/name", "GET", s.GetClusters)

	s.RegisterPath("/", "GET", HomeHandler, "")
}

func (s *Server) RegisterPath(path string, method string, f func(w http.ResponseWriter, r *http.Request), basePath ...string) {
	base := "/api/v1"
	if len(basePath) > 0 {
		base = basePath[0]
	}
	s.Router.HandleFunc(base + path, f).Methods(method)
}

func (s *Server) Run() {
	s.Initialize()
	log.Fatal(http.ListenAndServe(":8000", s.Router)) //TODO: Change port to SERVER_PORT
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("./public/index.html")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func (s *Server) GetCountryCases(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Cases{}
	s.getCountry(w, r, cases)
}

func (s *Server) GetStateCases(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Cases{}
	s.getState(w, r, cases)
}

func (s *Server) GetCountryCasesClusters(w http.ResponseWriter, r *http.Request) {
	cases := []*models.CasesCluster{}
	s.getCountry(w, r, cases)
}

func (s *Server) GetCountryCasesAgeCategory(w http.ResponseWriter, r *http.Request) {
	cases := []*models.CasesAgeCategory{}
	s.getCountry(w, r, cases)
}

func (s *Server) GetStateCasesAgeCategory(w http.ResponseWriter, r *http.Request) {
	cases := []*models.CasesAgeCategory{}
	s.getState(w, r, cases)
}

func (s *Server) GetCountryDeaths(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Deaths{}
	s.getCountry(w, r, cases)
}

func (s *Server) GetStateDeaths(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Deaths{}
	s.getState(w, r, cases)
}

func (s *Server) GetStateICU(w http.ResponseWriter, r *http.Request) {
	cases := []*models.ICU{}
	s.getState(w, r, cases)
}

func (s *Server) GetStateHospital(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Hospital{}
	s.getState(w, r, cases)
}

func (s *Server) GetStateQuarantineCentre(w http.ResponseWriter, r *http.Request) {
	cases := []*models.QuarantineCentre{}
	s.getState(w, r, cases)
}

//TODO: this prolly has special handling
func (s *Server) GetClusters(w http.ResponseWriter, r *http.Request) {
	cases := []*models.Clusters{}
	s.getCountry(w, r, cases)
}