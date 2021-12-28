package server

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func parseStates(states_str string) ([]string, error) {
	if states_str == "" {
		return []string{}, errors.New("requires at least one state")
	}

	states_raw := strings.Split(states_str, ",")
	states := []string{}
	for _, state_raw := range states_raw {
		state, key_exists := statesAbrvMap[state_raw]
		if key_exists {
			states = append(states, state)
		} else {
			return []string{}, errors.New("unknown state(s) passed as param")
		}
	}

	return states, nil
}

type CountryRequests struct {
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate time.Time `form:"end_date" time_format:"2006-01-02"`
}

type StateRequests struct {
	StartDate time.Time `form:"start_date" time_format:"2006-01-02"`
	EndDate time.Time `form:"end_date" time_format:"2006-01-02"`
	State string
}

func (s *Server) getCountry(c *gin.Context, out interface{}) error {
	var query CountryRequests
	c.BindQuery(&query)

	//TODO: Handle different combination of start_date, end_date
	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Where("date BETWEEN ? AND ?", query.StartDate, query.EndDate).Find(&val, "state ='MY'")
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		return tx.Error
	}
	c.JSON(200, val)
	return nil
}

func (s *Server) getState(c *gin.Context, out interface{}) error {
	var query StateRequests
	c.BindQuery(&query)

	states, err := parseStates(query.State)
	if err != nil {
		log.Errorf("%v?%v: %v", c.Request.URL.Path, c.Request.URL.Query(), err)
		return err
	}

	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Where("date BETWEEN ? AND ?", query.StartDate, query.EndDate).Where("state IN ?", states).Find(&val)
	if tx.Error != nil {
		log.Errorf("%v?%v: %v",c.Request.URL.Path, c.Request.URL.Query(), tx.Error)
		return tx.Error
	}
	c.JSON(200, val)
	return nil
}
