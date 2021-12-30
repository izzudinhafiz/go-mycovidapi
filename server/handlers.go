package server

import (
	"github.com/gin-gonic/gin"
	"github.com/izzudinhafiz/go-mycovidapi/models"
)

// GetCountry Cases godoc
// @summary      List Cases for Country
// @description  Gets a list of Cases for Malaysia. If no parameters are passed, will return all available data
// @tags         cases
// @produce      json
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Cases
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /country/cases [GET]
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