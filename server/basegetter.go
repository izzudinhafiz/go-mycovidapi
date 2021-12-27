package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"time"

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

func parseStartEndDate(r *http.Request) (models.Date, models.Date, error) {
	null_date := models.Date(time.Time{})
	start_date_str := r.FormValue("start_date")
	end_date_str := r.FormValue("end_date")
	if end_date_str == "" {
		return null_date, null_date, errors.New("requires end date")
	}

	var start_date time.Time
	var end_date time.Time
	var err error

	if start_date_str != "" {
		start_date, err = time.Parse("2006-01-02", start_date_str)
		if err != nil {
			return null_date, null_date, errors.New("invalid start date format")
		}
	}

	end_date, err = time.Parse("2006-01-02", end_date_str)
	if err != nil {
		return null_date, null_date, errors.New("invalid end date format")
	}

	return models.Date(start_date), models.Date(end_date), nil
}

func parseStates(r *http.Request) ([]string, error) {
	states_str := r.FormValue("state")
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

func (s *Server) getCountry(w http.ResponseWriter, r *http.Request, out interface{}) error {
	start_date, end_date, err := parseStartEndDate(r)

	if err != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, err)
		return err
	}
	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Where("date BETWEEN ? AND ?", start_date, end_date).Find(&val, "state ='Malaysia'")
	if tx.Error != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, tx.Error)
		return tx.Error
	}
	jsonResponse, err := json.Marshal(val)
	if err != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	return nil
}

func (s *Server) getState(w http.ResponseWriter, r *http.Request, out interface{}) error {
	start_date, end_date, err := parseStartEndDate(r)

	if err != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, err)
		return err
	}

	states, err := parseStates(r)
	if err != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, err)
		return err
	}

	val := reflect.ValueOf(out).Interface()
	tx := s.DB.Where("date BETWEEN ? AND ?", start_date, end_date).Where("state IN ?", states).Find(&val)
	if tx.Error != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, tx.Error)
		return tx.Error
	}
	jsonResponse, err := json.Marshal(val)
	if err != nil {
		log.Errorf("%v?%v: %v", r.URL.Path,r.URL.RawQuery, err)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

	return nil
}
