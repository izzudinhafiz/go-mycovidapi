package models

import (
	"time"
)
type QuarantineCentre struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-"`
	Date 	Date 			`gorm:"primarykey" csv:"date"`
	State 	string			`gorm:"primarykey" csv:"state"`
	Beds int				`csv:"beds"`
	AdmitTotal int			`csv:"admitted_total"`
	AdmitSuspected int		`csv:"admitted_pui"`
	AdmitPositive int		`csv:"admitted_covid"`
	AdmitNegative int 		`csv:"-"`
	DischargeTotal int 		`csv:"discharged_total"`
	DischargeSuspected int	`csv:"discharged_pui"`
	DischargePositive int	`csv:"discharged_covid"`
	DischargeNegative int	`csv:"-"`
	Total		int			`csv:"-"`
	Suspected int			`csv:"pkrc_pui"`
	Positive int			`csv:"pkrc_covid"`
	Negative int			`csv:"pkrc_noncovid"`
}

type Hospital struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-"`
	Date 	Date 			`gorm:"primarykey" csv:"date"`
	State 	string			`gorm:"primarykey" csv:"state"`
	Beds int				`csv:"beds"`
	BedsCovid int			`csv:"beds_covid"`
	BedsNonCrit int			`csv:"beds_noncrit"`
	AdmitTotal int			`csv:"admitted_total"`
	AdmitSuspected int		`csv:"admitted_pui"`
	AdmitPositive int		`csv:"admitted_covid"`
	AdmitNegative int		`csv:"-"`
	DischargeTotal int		`csv:"discharged_total"`
	DischargeSuspected int	`csv:"dicsharged_pui"`
	DischargePositive int	`csv:"discharged_covid"`
	DischargeNegative int	`csv:"-"`
	Total int 				`csv:"-"`
	Suspected int			`csv:"hosp_pui"`
	Positive int			`csv:"hosp_covid"`
	Negative int			`csv:"hosp_noncovid"`
}

type ICU struct {
	UpdatedAt time.Time		`gorm:"autoUpdateTime" csv:"-"`
	Date 	Date 			`gorm:"primarykey" csv:"date"`
	State 	string			`gorm:"primarykey" csv:"state"`
	Beds int				`csv:"beds_icu"`
	BedsRep int				`csv:"beds_icu_rep"`
	BedsTotal int			`csv:"beds_icu_total"`
	BedsCovid int			`csv:"beds_icu_covid"`
	Suspected int			`csv:"icu_pui"`
	Positive int			`csv:"icu_covid"`
	Negative int			`csv:"icu_noncovid"`
	Vent int				`csv:"vent"`
	VentPortable int		`csv:"vent_port"`
	VentProbable int		`csv:"vent_pui"`
	VentPositive int		`csv:"vent_covid"`
	VentNegative int		`csv:"vent_noncovid"`
	VentUsed int 			`csv:"vent_used"`
	VentPortableUsed int 	`csv:"vent_port_used"`
}