package models

import (
	"time"
)

type Deaths struct {
	UpdatedAt 			time.Time	`gorm:"autoUpdateTime" csv:"-"`
	Date 				Date 		`gorm:"primarykey" csv:"date"`
	State 				string		`gorm:"primarykey" csv:"state"`
	New 				int 		`csv:"deaths_new"`
	BroughtInDead 		int 		`csv:"deaths_bid"`
	NewOnDate 			int 		`csv:"deaths_new_dod"` // Deaths based on date of death
	BroughtInDeadOnDate int 		`csv:"deaths_bid_dod"` // Brought In Dead on date of death
	VaxPartial 			int 		`csv:"deaths_pvax"`
	VaxFull 			int 		`csv:"deaths_fvax"`
	DayGap 				int 		`csv:"deaths_tat"` // Median days between date of death and date of report for all deaths reported on the day
}