package utils

import (
	"math/rand"
)

func ShouldIDeploy() bool {
	return !IsFriday() &&
		!IsWeekend() &&
		!IsHolidays() &&
		!IsAfternoon()
}

func GetRandom(reasons []string) string {
	return reasons[rand.Intn(len(reasons))]
}

func GetReason() []string {
	if IsDayBeforeChristmas() {
		return ReasonsForDayBeforeChristmas
	}

	if IsChristmas() {
		return ReasonsForChristmas
	}

	if IsNewYear() {
		return ReasonsNewYear
	}

	if IsFriday13th() {
		return ReasonsForFriday13th
	}

	if IsFridayAfternoon() {
		return ReasonsForFridayAfternoon
	}

	if IsFriday() {
		return ReasonsToNotDeploy
	}

	if IsThursdayAfternoon() {
		return ReasonsForThursdayAfternoon
	}

	if IsWeekend() {
		return ReasonsForWeekend
	}

	if IsAfternoon() {
		return ReasonsForAfternoon
	}

	return ReasonsToDeploy
}
