package models

import (
	"time"
)

type Cases struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 		`gorm:"primarykey" csv:"date" json:"date"`
	State 	string		`gorm:"primarykey" csv:"state" json:"state"`
	New 	  	int		`csv:"cases_new" json:"newCases"`
	Import 	  	int		`csv:"cases_import" json:"importCases"`
	Active 	  	int		`csv:"cases_active" json:"activeCases"`
	Recovered 	int		`csv:"cases_recovered" json:"recovered"`
	Clusters  	int		`csv:"cases_cluster" json:"clusterCases"`
	VaxNone		int		`csv:"cases_unvax" json:"nonVaxCases"`
	VaxPartial 	int		`csv:"casses_pvax" json:"partialVaxCases"`
	VaxFull		int		`csv:"cases_fvax" json:"fullyVaxCases"`
}

// From cases_malaysia.csv
type CasesCluster struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 		`gorm:"primarykey" csv:"date" json:"date"`
	State 	string		`gorm:"primarykey" csv:"-" json:"state"`
	Import    int		`csv:"cluster_import" json:"import"`
	Religious int		`csv:"cluster_religious" json:"religious"`
	Community int		`csv:"cluster_community" json:"community"`
	HighRisk  int		`csv:"cluster_highRisk" json:"highRisk"`
	Education int		`csv:"cluster_education" json:"education"`
	Detention int		`csv:"cluster_detentionCentre" json:"detention"`
	Workplace int		`csv:"cluster_workplace" json:"workplace"`
}

type CasesAgeCategory struct {
	UpdatedAt time.Time	`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 		`gorm:"primarykey" csv:"date" json:"date"`
	State 	string		`gorm:"primarykey" csv:"state" json:"state"`
	Child 	   int		`csv:"cases_child" json:"children"`
	Adolescent int		`csv:"cases_adolescent" json:"adolescent"`
	Adult      int		`csv:"cases_adult" json:"adult"`
	Elderly    int		`csv:"cases_elderly" json:"elderly"`
}