package models

import (
	"time"
)

type Cases struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-" json:"updated_at"`
	Date 	Date 		`gorm:"primarykey" csv:"date" json:"date"`
	State 	string		`gorm:"primarykey" csv:"state" json:"state"`
	New 	  	int		`csv:"cases_new" json:"new_cases"`
	Import 	  	int		`csv:"cases_import" json:"import_cases"`
	Active 	  	int		`csv:"cases_active" json:"active"`
	Recovered 	int		`csv:"cases_recovered" json:"recovered"`
	Clusters  	int		`csv:"cases_cluster" json:"clusters"`
	VaxNone		int		`csv:"cases_unvax" json:"non_vax_cases"`
	VaxPartial 	int		`csv:"casses_pvax" json:"partial_vax_cases"`
	VaxFull		int		`csv:"cases_fvax" json:"fully_vax_cases"`
}

// From cases_malaysia.csv
type CasesCluster struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-"`
	Date 	Date 		`gorm:"primarykey" csv:"date"`
	State 	string		`gorm:"primarykey" csv:"-"`
	Import    int		`csv:"cluster_import"`
	Religious int		`csv:"cluster_religious"`
	Community int		`csv:"cluster_community"`
	HighRisk  int		`csv:"cluster_highRisk"`
	Education int		`csv:"cluster_education"`
	Detention int		`csv:"cluster_detentionCentre"`
	Workplace int		`csv:"cluster_workplace"`
}

type CasesAgeCategory struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-"`
	Date 	Date 		`gorm:"primarykey" csv:"date"`
	State 	string		`gorm:"primarykey" csv:"state"`
	Child 	   int		`csv:"cases_child"`
	Adolescent int		`csv:"cases_adolescent"`
	Adult      int		`csv:"cases_adult"`
	Elderly    int		`csv:"cases_elderly"`
}