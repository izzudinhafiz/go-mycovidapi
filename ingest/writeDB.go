package ingest

import (
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/izzudinhafiz/go-mycovidapi/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var statesMap = map[string]string {
	"Johor": "JHR",
	"Kedah": "KDH",
	"Kelantan": "KTN",
	"Melaka": "MLK",
	"Negeri Sembilan": "NSN",
	"Pahang": "PHG",
	"Pulau Pinang": "PNG",
	"Perak": "PRK",
	"Perlis": "PLS",
	"Sabah": "SBH",
	"Sarawak": "SWK",
	"Selangor": "SGR",
	"Terengganu": "TRG",
	"W.P. Kuala Lumpur": "KUL",
	"W.P. Labuan": "LBN",
	"W.P. Putrajaya": "PJY",
}

var statesIDMap = map[int]string {
	1: "JHR",
	2: "KDH",
	3: "KTN",
	4: "MLK",
	5: "NSN",
	6: "PHG",
	7: "PNG",
	8: "PRK",
	9: "PLS",
	10: "SBH",
	11: "SWK",
	12: "SGR",
	13: "TRG",
	14: "KUL",
	15: "LBN",
	16: "PJY",
}


func writeCases(db *gorm.DB) {
	log.Info("Writing Cases")
	cases_state := []*models.Cases{}
	cases_f := openFileWithErrorHandling("cases_state.csv")

	if err := gocsv.UnmarshalFile(cases_f, &cases_state); err != nil {
		 log.Error("Failed to read State Cases from cases_state.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	cases_f.Close()

	for _, row := range cases_state {
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx := db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cases_state, 1000)
	if tx.Error != nil {
		log.Error("Failed to write State Cases to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	cases_malaysia := []*models.Cases{}
	cases_msia_f := openFileWithErrorHandling("cases_malaysia.csv")
	if err := gocsv.UnmarshalFile(cases_msia_f, &cases_malaysia); err != nil {
		log.Error("Failed to read Malaysia Cases from cases_malaysia.csv\n")
		log.Errorf("\t%v\n", err)
		}
	cases_msia_f.Close()

	for _, case_msia := range cases_malaysia {
		case_msia.State = "MY"
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cases_malaysia, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Malaysia Cases to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	cases_clusters := []*models.CasesCluster{}
	cases_clusters_f := openFileWithErrorHandling("cases_malaysia.csv")
	if err := gocsv.UnmarshalFile(cases_clusters_f, &cases_clusters); err != nil {
		log.Error("Failed to read Cases Clusters from cases_malaysia.csv\n")
		log.Errorf("\t%v\n", err)
		}
	cases_clusters_f.Close()

	for _, case_clusters := range cases_clusters {
		case_clusters.State = "MY"
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cases_clusters, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Cases Clusters to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	cases_age := []*models.CasesAgeCategory{}
	cases_age_f := openFileWithErrorHandling("cases_malaysia.csv")
	if err := gocsv.UnmarshalFile(cases_age_f, &cases_age); err != nil {
		log.Error("Failed to read Malaysia Cases Age from cases_malaysia.csv\n")
		log.Errorf("\t%v\n", err)
		}
	cases_age_f.Close()

	for _, case_age := range cases_age {
		case_age.State = "MY"
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cases_age, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Malaysia Cases Age to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	cases_age = []*models.CasesAgeCategory{}
	cases_age_f = openFileWithErrorHandling("cases_state.csv")
	if err := gocsv.UnmarshalFile(cases_age_f, &cases_age); err != nil {
		log.Error("Failed to read State Cases Age from cases_state.csv\n")
		log.Errorf("\t%v\n", err)
		}
	cases_age_f.Close()

	for _, row := range cases_age {
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cases_age, 1000)
	if tx.Error != nil {
		log.Error("Failed to write State Cases Age to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}
}

func writeDeaths(db *gorm.DB) {
	log.Info("Writing Deaths")
	deaths_state := []*models.Deaths{}
	deaths_f := openFileWithErrorHandling("deaths_state.csv")

	if err := gocsv.UnmarshalFile(deaths_f, &deaths_state); err != nil {
		 log.Error("Failed to read State Deaths from deaths_state.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	deaths_f.Close()

	for _, row := range deaths_state {
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx := db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&deaths_state, 1000)
	if tx.Error != nil {
		log.Error("Failed to write State Deaths to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}


	deaths_msia := []*models.Deaths{}
	deaths_msia_f := openFileWithErrorHandling("deaths_malaysia.csv")
	if err := gocsv.UnmarshalFile(deaths_msia_f, &deaths_msia); err != nil {
		log.Error("Failed to read Msia Deaths from deaths_malaysia.csv\n")
		log.Errorf("\t%v\n", err)
		}
	deaths_msia_f.Close()

	for _, death_msia := range deaths_msia {
		death_msia.State = "MY"
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&deaths_msia, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Msia Deaths to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}
}

func writeClusters(db *gorm.DB) {
	parseStateID := func (statesID string) []string {
		statesRaw := strings.Split(statesID, ",")
		var statesStr = make([]string, len(statesRaw))
		for i, states := range statesRaw {
			states = strings.TrimSpace(states)
			stateID, err := strconv.Atoi(states)
			if err == nil {
				statesStr[i] = statesIDMap[stateID]
			}
		}

		return statesStr
	}

	log.Info("Writing Clusters")
	clusters := []*models.Clusters{}
	clusters_f := openFileWithErrorHandling("clusters.csv")

	if err := gocsv.UnmarshalFile(clusters_f, &clusters); err != nil {
		 log.Error("Failed to read Clusters from clusters.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	clusters_f.Close()

	// Unpack clusters row since the states are originally in a list.
	cleanedClusters := []*models.Clusters{}
	for _, row := range clusters {
		states := parseStateID(row.State)
		for _, state := range states {
			newRow := models.Clusters{
				State: state,
				Cluster: row.Cluster,
				District: row.District,
				DateAnnounced: row.DateAnnounced,
				DateLastOnset: row.DateLastOnset,
				Category: row.Category,
				Status: row.Status,
				NewCases: row.NewCases,
				TotalCases: row.TotalCases,
				ActiveCases: row.ActiveCases,
				Tests: row.Tests,
				ICU: row.ICU,
				Deaths: row.Deaths,
				Recovered: row.Recovered,
				Description: row.Description,
			}
			cleanedClusters = append(cleanedClusters, &newRow)
		}
	}

	tx := db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&cleanedClusters, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Clusters to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}
}

func writeHealthcares(db * gorm.DB) {
	log.Info("Writing Healthcares")
	pkrc := []*models.QuarantineCentre{}
	pkrc_f := openFileWithErrorHandling("pkrc.csv")

	if err := gocsv.UnmarshalFile(pkrc_f, &pkrc); err != nil {
		 log.Error("Failed to read PKRC from pkrc.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	pkrc_f.Close()

	for _, row := range pkrc {
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.AdmitNegative = row.AdmitTotal - row.AdmitPositive - row.AdmitSuspected
			row.DischargeNegative = row.DischargeTotal - row.DischargePositive - row.DischargeSuspected
			row.Total = row.Suspected + row.Positive + row.Negative
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx := db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&pkrc, 1000)
	if tx.Error != nil {
		log.Error("Failed to write PKRC to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	hospital := []*models.Hospital{}
	hospital_f := openFileWithErrorHandling("hospital.csv")

	if err := gocsv.UnmarshalFile(hospital_f, &hospital); err != nil {
		 log.Error("Failed to read Hospital from hospital.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	hospital_f.Close()

	for _, row := range hospital {
		row.AdmitNegative = row.AdmitTotal - row.AdmitPositive - row.AdmitSuspected
		row.DischargeNegative = row.DischargeTotal - row.DischargePositive - row.DischargeSuspected
		row.Total = row.Suspected + row.Positive + row.Negative
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&hospital, 1000)
	if tx.Error != nil {
		log.Error("Failed to write Hospital to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}

	icu := []*models.ICU{}
	icu_f := openFileWithErrorHandling("icu.csv")

	if err := gocsv.UnmarshalFile(icu_f, &icu); err != nil {
		 log.Error("Failed to read ICU from icu.csv\n")
		 log.Errorf("\t%v\n", err)
		}
	icu_f.Close()

	for _, row := range icu {
		state_abrv, exists := statesMap[row.State]
		if exists {
			row.State = state_abrv
		} else {
			row.State = ""
		}
	}

	tx = db.Clauses(clause.OnConflict{UpdateAll: true,}).CreateInBatches(&icu, 1000)
	if tx.Error != nil {
		log.Error("Failed to write ICU to DB\n")
		log.Errorf("\t%v\n", tx.Error)
	}
}