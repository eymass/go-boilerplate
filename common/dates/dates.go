package dates

import "time"

func GetInterval(now time.Time) (time.Time, time.Time) {
	if IsDateAfter17(now) {
		tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 17, 0, 0, now.Nanosecond(), now.Location())
		today := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
		return today, tomorrow
	} else {
		yesterday := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
		yesterday = yesterday.AddDate(0, 0, -1)
		today := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
		return yesterday, today
	}
}

func IsDateAfter17(date time.Time) bool {
	hours, _, _ := date.Clock()
	return hours >= 17
}

func GetDates() (time.Time, time.Time) {
	var cstSh, _ = time.LoadLocation("Asia/Jerusalem")
	now := time.Now().In(cstSh)

	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 17, 0, 0, now.Nanosecond(), now.Location())
	today := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
	return today, tomorrow
}

func GetNowTime() time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Jerusalem")
	return time.Now().In(cstSh)
}

func GetTimeFromNowBackDays(date time.Time, days int) time.Time {
	return date.AddDate(0,0, -days)
}

func GetTimeFromMonthStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 17, 0, 0, 0, date.Location())
}

func IsDateBetweenDates(start time.Time, end time.Time, toCheck time.Time) bool {
	var cstSh, _ = time.LoadLocation("Asia/Jerusalem")
	startInTimeZone := start.In(cstSh)
	endInTimeZone := end.In(cstSh)
	toCheckInTimeZone := toCheck.In(cstSh)
	return toCheckInTimeZone.Before(endInTimeZone) && toCheckInTimeZone.After(startInTimeZone)
}
