package models

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type Date time.Time

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date(nullTime.Time)
	return
}

func (date Date) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}

func (date Date) GobEncode() ([]byte, error) {
	return time.Time(date).GobEncode()
}

func (date *Date) GobDecode(b []byte) error {
	return (*time.Time)(date).GobDecode(b)
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(date).Format("2006-01-02") + `"`), nil
}

func (date *Date) UnmarshalJSON(b []byte) error {
	return (*time.Time)(date).UnmarshalJSON(b)
}

func (date *Date) UnmarshalCSV(csv string) error {
	x, err := time.Parse("2006-01-02", csv)
	*date = Date(x)
	return err
}

// func (datetime *DateTime) UnmarshalCSV(csv string) (err error) {
// 	datetime.Time, err = time.Parse("2006-01-02", csv)
// 	return err
// }