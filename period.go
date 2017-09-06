package utils

import (
	"time"
)

// Period struct represents Year and Month Period
type Period struct {
	code uint32
}

// NewPeriod creates and return a Period initialized by period code
func NewPeriod(code uint32) *Period {
	return &Period{code}
}

// Code returns code (201701)
func (p Period) Code() int {
	return int(p.code)
}

// Month returns month (1 - 12)
func (p Period) Month() int {
	return int(p.code % 100)
}

// Year returns year (2000 ..)
func (p Period) Year() int {
	return int(p.code / 100)
}

// IsValid returns period validity based on Year and Month
func (p Period) IsValid() bool {
	return p.Year() > 1900 && p.Month() > 0 && p.Month() <= 12
}

// DateBegin return first day of period
func (p Period) DateBegin() time.Time {
	if p.IsValid() {
		return time.Date(int(p.Year()), time.Month(p.Month()), 1, 0, 0, 0, 0, time.UTC)
	}
	return p.DateZero()
}

// DateEnd return last day of period
func (p Period) DateEnd() time.Time {
	if p.IsValid() {
		var datePeriod = p.DateBegin()
		return datePeriod.AddDate(0, 1, -1)
	}
	return p.DateZero()
}

// Description return string like "January 2006"
func (p Period) Description() string {
	var datePeriod = p.DateBegin()
	return datePeriod.Format("January 2006")
}

// Equals compare Period with other Period by code and return (true x false)
func (p Period) Equals(other Period) bool {
	return (p.code == other.code)
}

// CompareTo compares Period with other Period by code and return (-1 x O x +1)
func (p Period) CompareTo(other Period) int {
	return (int(p.code) - int(other.code))
}

// OpGt compares Period with other Period by code and return greater (true x false)
func (p Period) OpGt(other Period) bool {
	return (p.CompareTo(other) > 0)
}

// OpLt compares Period with other Period by code and return lesser (true x false)
func (p Period) OpLt(other Period) bool {
	return (p.CompareTo(other) < 0)
}

// DateZero return zero time instant, January 1, year 1, 00:00:00 UTC.
func (p Period) DateZero() time.Time {
	return time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
}
