package models

import (
	"time"
)

type Deaths struct {
	UpdatedAt 			time.Time	`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 				Date 		`gorm:"primarykey" csv:"date" json:"date"`
	State 				string		`gorm:"primarykey" csv:"state" json:"state"`
	New 				int 		`csv:"deaths_new" json:"new"`
	BroughtInDead 		int 		`csv:"deaths_bid" json:"bid"`
	NewOnDate 			int 		`csv:"deaths_new_dod" json:"newDoD"` // Deaths based on date of death
	BroughtInDeadOnDate int 		`csv:"deaths_bid_dod" json:"bidDoD"` // Brought In Dead on date of death
	VaxPartial 			int 		`csv:"deaths_pvax" json:"partialVax"`
	VaxFull 			int 		`csv:"deaths_fvax" json:"fullyVax"`
	DayGap 				int 		`csv:"deaths_tat" json:"dayGap"` // Median days between date of death and date of report for all deaths reported on the day
}