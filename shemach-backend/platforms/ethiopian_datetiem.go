package platforms

import (
	"time"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
)

type DateHandler struct {
}

func NewDate(year uint, month uint8, day uint8, hour uint8, minute uint8, second uint8) *model.Date {
	return &model.Date{
		Years:   year,
		Months:  month,
		Days:    day,
		Hours:   hour,
		Minutes: minute,
	}
}

var (
	dayNames  = []string{"እሁድ", "ሰኞ", "ማክሰኞ", "ረቡዕ", "ሀሙስ", "አርብ", "ቅዳሜ"}
	yearNames = []string{"ማቴዎስ", "ማርቆስ", "ሉቃስ", "ዮሐንስ"}
)

// STARTING_YEAR := 1962
// STARTING_MONTH := 4
// STARTING_DAY := 22
// STARTING_DAY_INDEX := 2
// STARTING_YEAR_TITLE_INDEX := 1
// STARTING_HOUR := 21
// STARTING_MINUTE := 0
// STARTING_SECOND := 0

var SECONDS_IN_PUAGUME_YEAR = 31536000
var SECONDS_IN_YEAR = 31557600
var SECONDS_IN_MONTH = 2592000
var SECONDS_IN_DAYS = 86400
var SECONDS_IN_HOUR = 3600

/*
	UTC is 3 dekikawoch wedehaula from EAT therefore 0 time malet be EAT 3 seat malet new.
*/

func GetCurrentEthiopianTime() *model.Date {
	return UnixToEthiopianDate(int(time.Now().Unix()))
}

func UnixToEthiopianDate(unix int) *model.Date {
	s_year := 1962
	s_day_index := 6
	s_year_index := 1

	starting_secs := 12344400

	a_unix := starting_secs + unix
	a_c_unix := a_unix

	a_years := a_unix / SECONDS_IN_YEAR
	x := (a_years + 1) % 4
	y := (a_years - x)
	a_unix -= y * SECONDS_IN_YEAR
	var puagume_years int
	if (a_unix >= (730 * SECONDS_IN_DAYS)) && (a_unix < (1097 * SECONDS_IN_DAYS)) {
		// the year is in the zemene lukas
		s_year_index = 2
		puagume_years = a_unix / SECONDS_IN_PUAGUME_YEAR
		a_unix -= (puagume_years) * SECONDS_IN_PUAGUME_YEAR
		// SUBTRACTING THE 6 HOUR THAT'S COMING FROM THE ZEMENE MATHEOS
		a_unix -= (12 * SECONDS_IN_HOUR)
	} else if (a_unix >= (1096 * SECONDS_IN_DAYS)) && (a_unix < (1411 * SECONDS_IN_DAYS)) {
		// the days are more than
		s_year_index = 3
		puagume_years = a_unix / SECONDS_IN_PUAGUME_YEAR

		a_unix -= ((puagume_years - 1) * SECONDS_IN_PUAGUME_YEAR) - (SECONDS_IN_DAYS * 366)

		// SUBTRACTING THE 6 HOUR THAT'S COMING FROM THE ZEMENE MATHEOS
		a_unix -= (18 * SECONDS_IN_HOUR)
	} else if (a_unix >= (365 * SECONDS_IN_DAYS)) && a_unix < (1096*SECONDS_IN_DAYS) {
		s_year_index = 1
		puagume_years = a_unix / SECONDS_IN_PUAGUME_YEAR
		a_unix -= (puagume_years) * SECONDS_IN_PUAGUME_YEAR
		// SUBTRACTING THE 6 HOUR THAT'S COMING FROM THE ZEMENE MATHEOS
		a_unix -= (6 * SECONDS_IN_HOUR)
	} else if a_unix < 365 {
		s_year_index = 0
		puagume_years = 0
	}
	a_years = puagume_years + y

	a_months := a_unix / SECONDS_IN_MONTH

	a_unix -= (a_months * SECONDS_IN_MONTH)
	a_days := a_unix / SECONDS_IN_DAYS
	a_unix -= (a_days * SECONDS_IN_DAYS)
	a_hours := a_unix / SECONDS_IN_HOUR
	a_unix -= (a_hours * SECONDS_IN_HOUR)
	a_minutes := a_unix / 60
	a_unix -= (a_minutes * 60)
	a_seconds := a_unix

	// s_year_index = (s_year_index + ((puagume_years + y)) % 4
	s_day_index = ((s_day_index) + ((a_c_unix) / (SECONDS_IN_DAYS) % 7)) % 7

	hours := a_unix / SECONDS_IN_HOUR
	a_unix -= (hours * SECONDS_IN_HOUR)
	minutes := a_unix / 60
	a_unix -= (minutes * 60)

	if s_year_index == 2 && a_months == 12 && (a_days > 6) {
		a_years += 1
		a_days -= 6
		s_year_index = ((s_year_index + 1) % 4)
	} else if s_year_index == 2 && a_months == 12 && (a_days <= 6) {
		a_months += 1
	} else if s_year_index != 2 && a_months == 12 && a_days > 5 {
		a_years += 1
		a_days -= 5
		s_year_index = ((s_year_index + 1) % 4)
	} else if s_year_index != 2 && a_months == 12 && (a_days <= 5) {
		a_months += 1
	}

	s_year += a_years
	return &model.Date{
		Years:    uint(s_year),
		Months:   uint8(a_months),
		Days:     uint8(a_days),
		Hours:    uint8(a_hours),
		Minutes:  uint8(a_minutes),
		Seconds:  uint8(a_seconds),
		DayName:  dayNames[s_day_index],
		YearName: yearNames[s_year_index],
	}
}

func GetAgeUsingBirthDate(date *model.Date) (year int, month int) {
	currentDate := GetCurrentEthiopianTime()
	year = 0
	month = 0
	year = int(currentDate.Years - date.Years)
	month = int(currentDate.Months - date.Months)
	if month < 0 && currentDate.Months <= 12 {
		month = 12 + month
		year -= 1
	} else if month < 0 && currentDate.Months == 13 {
		if date.Days >= currentDate.Days {
			month = 12 + (month + 1)
			year -= 1
		} else {
			month = 12 + (month - 1)
			year -= 1
		}
	}
	return year, month
}
