package server

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izzudinhafiz/go-mycovidapi/models"
	log "github.com/sirupsen/logrus"
)

var statesAbrvMap = map[string]string {
	"JHR": "Johor",
	"KDH": "Kedah",
	"KTN": "Kelantan",
	"MLK": "Melaka",
	"NSN": "Negeri Sembilan",
	"PLS": "Perlis",
	"SBH": "Sabah",
	"SWK": "Sarawak",
	"TRG": "Terengganu",
	"SGR": "Selangor",
	"PRK": "Perak",
	"PHG": "Pahang",
	"KUL": "W.P. Kuala Lumpur",
	"LBN": "W.P. Labuan",
	"PJY": "W.P. Putrajaya",
	"PNG": "Pulau Pinang",
}

var (
	ErrorNoState = errors.New("requires at least one state")
	ErrorUnknownState = errors.New("unknown state(s) passed as param")
)

func parseStates(states_str string) ([]string, error) {
	if states_str == "" {
		return []string{}, ErrorNoState
	}

	states_raw := strings.Split(states_str, ",")
	for i, state_raw := range states_raw {
		_, key_exists := statesAbrvMap[strings.ToUpper(state_raw)]
		if !key_exists {
			return []string{}, ErrorUnknownState
		}
		states_raw[i] = strings.ToUpper(state_raw)
	}

	return states_raw, nil
}

type CountryRequests struct {
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate time.Time `form:"end_date" time_format:"2006-01-02"`
}

type StateRequests struct {
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate time.Time `form:"end_date" time_format:"2006-01-02"`
	State string `form:"state_id"`
}

type ValidationError struct {
	ErrorMessage string `json:"msg"`
	Location string `json:"loc"`
	Type string `json:"type"`
}

func (s *Server) getCountry(c *gin.Context, out interface{}) error {
	var query CountryRequests
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.AbortWithStatusJSON(422, ValidationError{
			err.Error(),
			"start_date | end_date",
			"string",
		})
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), err)
		return err
	}

	val := reflect.ValueOf(out).Interface()
	tx := s.DB
	if query.StartDate.IsZero() && query.EndDate.IsZero() {
	} else if !query.StartDate.IsZero() && query.EndDate.IsZero() {
		tx = s.DB.Where("date >= ?", query.StartDate)
	} else if query.StartDate.IsZero() && !query.EndDate.IsZero() {
		tx = s.DB.Where("date <= ?", query.EndDate)
	} else {
		tx = s.DB.Where("date BETWEEN ? AND ?", query.StartDate, query.EndDate)
	}

	tx = tx.Order("date").Find(&val, "state ='MY'")
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return tx.Error
	}
	c.JSON(http.StatusOK, val)
	return nil
}

func (s *Server) getState(c *gin.Context, out interface{}) error {
	var query StateRequests
	c.BindQuery(&query)

	states, err := parseStates(query.State)
	if err != nil {
		log.Errorf("%v?%v: %v", c.Request.URL.Path, c.Request.URL.Query(), err)
		errorMsg := ValidationError{ErrorMessage: "Validation Error",}
		if err == ErrorNoState {
			errorMsg = ValidationError{
				ErrorMessage: "requires at least one state",
				Location: "state_id",
				Type: "string (comma separated)",
			}
		} else if err == ErrorUnknownState{
			errorMsg = ValidationError{
				ErrorMessage: "invalid states in query parameter",
				Location: "state_id",
				Type: "string (comma separated)",
			}
		}
		c.AbortWithStatusJSON(422, errorMsg)
		return err
	}

	val := reflect.ValueOf(out).Interface()
	tx := s.DB
	if query.StartDate.IsZero() && query.EndDate.IsZero() {
	} else if !query.StartDate.IsZero() && query.EndDate.IsZero() {
		tx = s.DB.Where("date >= ?", query.StartDate)
	} else if query.StartDate.IsZero() && !query.EndDate.IsZero() {
		tx = s.DB.Where("date <= ?", query.EndDate)
	} else {
		tx = s.DB.Where("date BETWEEN ? AND ?", query.StartDate, query.EndDate)
	}

	tx = tx.Where("state IN ?", states).Order("date, state").Find(&val)
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return tx.Error
	}
	c.JSON(http.StatusOK, val)
	return nil
}

func (s *Server) getClusterList(c *gin.Context, out interface{}) error {
	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Model(models.Clusters{}).Find(&val)
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return tx.Error
	}
	c.JSON(http.StatusOK, val)
	return nil
}

func (s *Server) getActiveClusters(c *gin.Context, out interface{}) error {
	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Model(models.Clusters{}).Where("status = ?", "active").Find(&val)
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return tx.Error
	}
	c.JSON(http.StatusOK, val)
	return nil
}