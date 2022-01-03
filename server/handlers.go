package server

import (
	"github.com/gin-gonic/gin"
	"github.com/izzudinhafiz/go-mycovidapi/models"
)

// GetCountryCases godoc
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

// GetStateCases godoc
// @summary      List Cases for State
// @description  Gets a list of Cases for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         cases
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Cases
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/cases [GET]
func (s *Server) GetStateCases(c *gin.Context) {
	respModel := []*models.Cases{}
	s.getState(c, respModel)
}

// GetCountryCasesClusters godoc
// @summary      List Cases based on Cluster category for Country
// @description  Gets a list of Cases based on cluster category for Malaysia. If no parameters are passed, will return all available data
// @tags         cases
// @produce      json
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.CasesCluster
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /country/cases/clusters [GET]
func (s *Server) GetCountryCasesClusters(c *gin.Context) {
	respModel := []*models.CasesCluster{}
	s.getCountry(c, respModel)
}

// GetCountryCasesAgeCategory godoc
// @summary      List Cases based on Age category for Country
// @description  Gets a list of Cases based on age category for Malaysia. If no parameters are passed, will return all available data
// @tags         cases
// @produce      json
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.CasesAgeCategory
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /country/cases/age [GET]
func (s *Server) GetCountryCasesAgeCategory(c *gin.Context) {
	respModel := []*models.CasesAgeCategory{}
	s.getCountry(c, respModel)
}

// GetStateCasesAgeCategory godoc
// @summary      List Cases based on Age category for State
// @description  Gets a list of Cases based on age category for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         cases
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.CasesAgeCategory
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/cases/age [GET]
func (s *Server) GetStateCasesAgeCategory(c *gin.Context) {
	respModel := []*models.CasesAgeCategory{}
	s.getState(c, respModel)
}

// GetCountryDeaths godoc
// @summary      List Deaths for Country
// @description  Gets a list of Deaths for Malaysia. If no parameters are passed, will return all available data
// @tags         deaths
// @produce      json
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Deaths
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /country/deaths [GET]
func (s *Server) GetCountryDeaths(c *gin.Context) {
	respModel := []*models.Deaths{}
	s.getCountry(c, respModel)
}

// GetStateDeaths godoc
// @summary      List Deaths for State
// @description  Gets a list of Deaths for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         deaths
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Deaths
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/deaths [GET]
func (s *Server) GetStateDeaths(c *gin.Context) {
	respModel := []*models.Deaths{}
	s.getState(c, respModel)
}

// GetStateICU godoc
// @summary      List ICU utilization for State
// @description  Gets a list of ICU utilization for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         healthcare
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.ICU
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/icu [GET]
func (s *Server) GetStateICU(c *gin.Context) {
	respModel := []*models.ICU{}
	s.getState(c, respModel)
}

// GetStateHospital godoc
// @summary      List Hospital utilization for State
// @description  Gets a list of Hospital utilization for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         healthcare
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Hospital
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/hospital [GET]
func (s *Server) GetStateHospital(c *gin.Context) {
	respModel := []*models.Hospital{}
	s.getState(c, respModel)
}

// GetStateQuarantineCentre godoc
// @summary      List Quarantine Centre utilization for State
// @description  Gets a list of Quarantine Centre (PKRC) utilization for states in Malaysia. If no date parameters are passed, will return all available data
// @tags         healthcare
// @produce      json
// @param		 state_id query string true "State abbreviation. e.g. Pahang=PHG. Can be chained, e.g. state_id=PHG,KUL,PNG"
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.QuarantineCentre
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /state/pkrc [GET]
func (s *Server) GetStateQuarantineCentre(c *gin.Context) {
	respModel := []*models.QuarantineCentre{}
	s.getState(c, respModel)
}

//TODO: this prolly has special handling
// GetClustersInfo godoc
// @summary      List Cluster information
// @description  Gets a list of cluster information.
// @tags         clusters
// @produce      json
// @param        start_date query string false "YYYY-MM-DD. If omitted, will return data from start of records"
// @param        end_date query string false "YYYY-MM-DD. If omitted, will return data to end of records"
// @success      200  {array}   models.Clusters
// @failure      422  {object}  ValidationError "Invalid parameters passed"
// @router       /clusters/info [GET]
func (s *Server) GetClustersInfo(c *gin.Context) {
	respModel := []*models.Clusters{}
	s.getClusterList(c, respModel)
}

// GetClustersInfo godoc
// @summary      List all cluster
// @description  Gets a list of cluster.
// @tags         clusters
// @produce      json
// @success      200  {array}   models.ClustersList
// @router       /clusters [GET]
func (s *Server) GetClusterList(c *gin.Context) {
	respModel := []*models.ClustersList{}
	s.getClusterList(c, respModel)
}

// GetClustersInfo godoc
// @summary      List all active cluster
// @description  Gets a list of active cluster.
// @tags         clusters
// @produce      json
// @success      200  {array}   models.Clusters
// @router       /clusters/active [GET]
func (s *Server) GetActiveClusters(c *gin.Context) {
	respModel := []*models.Clusters{}
	s.getActiveClusters(c, respModel)
}