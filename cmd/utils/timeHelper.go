package utils

import "time"

// Default timezone greece athens
const DEFAULT_TIMEZONE = "Europe/Athens"

// declare location string
var location string

func SetUpTimezone(timezone string) {
	if timezone != "" {
		loc, _ := time.LoadLocation(timezone)
		time.Local = loc
	} else {
		loc, _ := time.LoadLocation(DEFAULT_TIMEZONE)
		time.Local = loc
	}
}

func GetDate() time.Time {
	return time.Now()
}

func IsFriday() bool {
	return GetDate().Weekday() == time.Friday
}

func IsThursday() bool {
	return GetDate().Weekday() == time.Thursday
}

func Is13th() bool {
	return GetDate().Day() == 13
}

func IsAfternoon() bool {
	return GetDate().Hour() >= 16
}

func IsThursdayAfternoon() bool {
	return IsThursday() && IsAfternoon()
}

func IsFridayAfternoon() bool {
	return IsFriday() && IsAfternoon()
}

func IsFriday13th() bool {
	return IsFriday() && Is13th()
}

func IsWeekend() bool {
	return GetDate().Weekday() == time.Saturday || GetDate().Weekday() == time.Sunday
}

func IsDayBeforeChristmas() bool {
	return GetDate().Month() == time.December && GetDate().Day() == 24 && GetDate().Hour() >= 16
}

func IsChristmas() bool {
	return GetDate().Month() == time.December && GetDate().Day() == 25
}

func IsNewYear() bool {
	return (GetDate().Month() == time.December &&
		GetDate().Day() == 31 &&
		GetDate().Hour() >= 16) ||
		(GetDate().Month() == time.January && GetDate().Day() == 1)
}

func IsHolidays() bool {
	return IsDayBeforeChristmas() || IsChristmas() || IsNewYear()
}
