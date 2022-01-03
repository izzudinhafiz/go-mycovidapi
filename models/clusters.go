package models

import (
	"time"
)

type Clusters struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	State 	string			`gorm:"primarykey" csv:"state" json:"state"`
	Cluster string			`gorm:"primarykey" csv:"cluster" json:"clusterName"`
	District string			`csv:"district" json:"district"`
	DateAnnounced Date		`csv:"date_announced" json:"dateAnnounced"`
	DateLastOnset Date		`csv:"date_last_onset" json:"dateLastOnset"`
	Category string			`csv:"category" json:"category"`
	Status string			`csv:"status" json:"status"`
	NewCases int			`csv:"cases_new" json:"newCases"`
	TotalCases int			`csv:"cases_total" json:"totalCases"`
	ActiveCases int			`csv:"cases_active" json:"activeCases"`
	Tests int				`csv:"tests" json:"testsConducted"`
	ICU int					`csv:"icu" json:"ICU"`
	Deaths int				`csv:"deaths" json:"deaths"`
	Recovered int			`csv:"recovered" json:"recovered"`
	Description string 		`csv:"summary_en" json:"description"`
}

type ClustersList struct {
	UpdatedAt time.Time		`json:"updatedAt"`
	State 	string			`json:"state"`
	Cluster string			`json:"clusterName"`
	District string			`json:"district"`
	DateAnnounced Date		`json:"dateAnnounced"`
	DateLastOnset Date		`json:"dateLastOnset"`
	Category string			`json:"category"`
	Status string			`json:"status"`
	Description string 		`json:"description"`
}