package models

import (
	"time"
)

type Clusters struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-"`
	State 	string			`gorm:"primarykey" csv:"state"`
	Cluster string			`gorm:"primarykey" csv:"cluster"`
	District string			`csv:"district"`
	DateAnnounced Date		`csv:"date_announced"`
	DateLastOnset Date		`csv:"date_last_onset"`
	Category string			`csv:"category"`
	Status string			`csv:"status"`
	NewCases int			`csv:"cases_new"`
	TotalCases int			`csv:"cases_total"`
	ActiveCases int			`csv:"cases_active"`
	Tests int				`csv:"tests"`
	ICU int					`csv:"icu"`
	Deaths int				`csv:"deaths"`
	Recovered int			`csv:"recovered"`
	Description string 		`csv:"summary_en"`
}