package model

import (
	"database/sql/driver"
	"time"
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

// MyTime 自定义时间格式
type MyTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *MyTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt MyTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// UnmarshalJSON implements the Json interface
func (nt *MyTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	if err != nil {
		*nt = MyTime{Valid: false}
	} else {
		*nt = MyTime{Time: now, Valid: true}
	}
	return nil
}

// MarshalJSON implements the json Valuer interface
func (nt MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(nt.Time).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

// String implements the interface to string
func (nt MyTime) String() string {
	return time.Time(nt.Time).Format(timeFormart)
}
