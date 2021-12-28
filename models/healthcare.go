package models

import (
	"time"
)
type QuarantineCentre struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 			`gorm:"primarykey" csv:"date" json:"date"`
	State 	string			`gorm:"primarykey" csv:"state" json:"state"`
	Beds int				`csv:"beds" json:"beds"`
	AdmitTotal int			`csv:"admitted_total" json:"admittedTotal"`
	AdmitSuspected int		`csv:"admitted_pui" json:"admittedPUI"`
	AdmitPositive int		`csv:"admitted_covid" json:"admittedPositive"`
	AdmitNegative int 		`csv:"-" json:"admittedNegative"`
	DischargeTotal int 		`csv:"discharged_total" json:"dischargedTotal"`
	DischargeSuspected int	`csv:"discharged_pui" json:"dischargedPUI"`
	DischargePositive int	`csv:"discharged_covid" json:"dischargedPositive"`
	DischargeNegative int	`csv:"-" json:"dischargedNegative"`
	Total		int			`csv:"-" json:"stockTotal"`
	Suspected int			`csv:"pkrc_pui" json:"stockPUI"`
	Positive int			`csv:"pkrc_covid" json:"stockPositive"`
	Negative int			`csv:"pkrc_noncovid" json:"stockNegative"`
}

type Hospital struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 			`gorm:"primarykey" csv:"date" json:"date"`
	State 	string			`gorm:"primarykey" csv:"state" json:"state"`
	Beds int				`csv:"beds" json:"beds"`
	BedsCovid int			`csv:"beds_covid" json:"bedsCovid"`
	BedsNonCrit int			`csv:"beds_noncrit" json:"bedsNonCrit"`
	AdmitTotal int			`csv:"admitted_total" json:"admittedTotal"`
	AdmitSuspected int		`csv:"admitted_pui" json:"admittedSuspected"`
	AdmitPositive int		`csv:"admitted_covid" json:"admittedPositive"`
	AdmitNegative int		`csv:"-" json:"admittedNegative"`
	DischargeTotal int		`csv:"discharged_total" json:"dischargedTotal"`
	DischargeSuspected int	`csv:"dicsharged_pui" json:"dischargedSuspected"`
	DischargePositive int	`csv:"discharged_covid" json:"dischargedPositive"`
	DischargeNegative int	`csv:"-" json:"dischargedNegative"`
	Total int 				`csv:"-" json:"stockTotal"`
	Suspected int			`csv:"hosp_pui" json:"stockPUI"`
	Positive int			`csv:"hosp_covid" json:"stockPositive"`
	Negative int			`csv:"hosp_noncovid" json:"stockNegative"`
}

type ICU struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-" json:"updatedAt"`
	Date 	Date 			`gorm:"primarykey" csv:"date" json:"date"`
	State 	string			`gorm:"primarykey" csv:"state" json:"state"`
	Beds int				`csv:"beds_icu" json:"beds"`
	BedsRep int				`csv:"beds_icu_rep" json:"bedsReserved"`
	BedsTotal int			`csv:"beds_icu_total" json:"bedsCapacityTotal"`
	BedsCovid int			`csv:"beds_icu_covid" json:"bedsForCovid"`
	Suspected int			`csv:"icu_pui" json:"stockPUI"`
	Positive int			`csv:"icu_covid" json:"stockPositive"`
	Negative int			`csv:"icu_noncovid" json:"stockNegative"`
	Vent int				`csv:"vent" json:"ventilators"`
	VentPortable int		`csv:"vent_port" json:"ventilatorsPortable"`
	VentProbable int		`csv:"vent_pui" json:"ventilatorsPUI"`
	VentPositive int		`csv:"vent_covid" json:"ventilatorsPositive"`
	VentNegative int		`csv:"vent_noncovid" json:"ventilatorsNegative"`
	VentUsed int 			`csv:"vent_used" json:"ventilatorsUsed"`
	VentPortableUsed int 	`csv:"vent_port_used" json:"ventilatorsPortableUsed"`
}