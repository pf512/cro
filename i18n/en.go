package i18n

// English

type en struct {
}

func (e en) languageCode() string {
	return "en"
}

func (e en) Am() string {
	return "AM"
}

func (e en) Pm() string {
	return "PM"
}

func (e en) CommaOnlyInMonthX0() string {
	return ""
}

func (e en) CommaOnlyInYearX0() string {
	return ""
}

func (e en) DayX0() string {
	return ""
}

func (e en) SetPeriodBeforeTime() bool {
	return false
}


func (e en) AtX0SecondsPastTheMinuteGt20() string {
		return ""
	}

func (e en) AtX0MinutesPastTheHourGt20() string {
	return ""
}

func (e en) CommaMonthX0ThroughMonthX1() string {
	return ""
}

func (e en) CommaYearX0ThroughYearX1() string {
	return ""
}

func (e en) Use24HourTimeFormatByDefault() bool {
	return false
}

func (e en) AnErrorOccuredWhenGeneratingTheExpressionD() string {
	return "An error occured when generating the expression description.  Check the cron expression syntax."
}
	
func (e en) EveryMinute() string {
	return "every minute"
}

func (e en) EveryHour() string {
	return "every hour"
}

func (e en) AtSpace() string {
	return "At "
}
	
func (e en) EveryMinuteBetweenX0AndX1() string {
	return "Every minute between %s and %s"
}
	
func (e en) At() string {
	return "At"
}
	
func (e en) SpaceAnd() string {
	return " and"
}
	
func (e en) EverySecond() string {
	return "every second"
}
	
func (e en) EveryX0Seconds() string {
	return "every %s seconds"
}
	
func (e en) SecondsX0ThroughX1PastTheMinute() string {
	return "seconds %s through %s past the minute"
}
	
func (e en) AtX0SecondsPastTheMinute() string {
	return "at %s seconds past the minute"
}
	
func (e en) EveryX0Minutes() string {
	return "every %s minutes"
}
	
func (e en) MinutesX0ThroughX1PastTheHour() string {
	return "minutes %s through %s past the hour"
}
	
func (e en) AtX0MinutesPastTheHour() string {
	return "at %s minutes past the hour"
}
	
func (e en) EveryX0Hours() string {
	return "every %s hours"
}
	
func (e en) BetweenX0AndX1() string {
	return "between %s and %s"
}
	
func (e en) AtX0() string {
	return "at %s"
}
	
func (e en) CommaEveryDay() string {
	return ", every day"
}
	
func (e en) CommaEveryX0DaysOfTheWeek() string {
	return ", every %s days of the week"
}
	
func (e en) CommaX0ThroughX1() string {
	return ", %s through %s"
}
	
func (e en) First() string {
	return "first"
}
	
func (e en) Second() string {
	return "second"
}
	
func (e en) Third() string {
	return "third"
}
	
func (e en) Fourth() string {
	return "fourth"
}
	
func (e en) Fifth() string {
	return "fifth"
}
	
func (e en) CommaOnThe() string {
	return ", on the "
}
	
func (e en) SpaceX0OfTheMonth() string {
	return " %s of the month"
}
	
func (e en) LastDay() string {
	return "the last day"
}
	
func (e en) CommaOnTheLastX0OfTheMonth() string {
	return ", on the last %s of the month"
}
	
func (e en) CommaOnlyOnX0() string {
	return ", only on %s"
}
	
func (e en) CommaAndOnX0() string {
	return ", and on %s"
}
	
func (e en) CommaEveryX0Months() string {
	return ", every %s months"
}
	
func (e en) CommaOnlyInX0() string {
	return ", only in %s"
}
	
func (e en) CommaOnTheLastDayOfTheMonth() string {
	return ", on the last day of the month"
}
	
func (e en) CommaOnTheLastWeekdayOfTheMonth() string {
	return ", on the last weekday of the month"
}
	
func (e en) CommaDaysBeforeTheLastDayOfTheMonth() string {
	return ", %s days before the last day of the month"
}
	
func (e en) FirstWeekday() string {
	return "first weekday"
}
	
func (e en) WeekdayNearestDayX0() string {
	return "weekday nearest day %s"
}
	
func (e en) CommaOnTheX0OfTheMonth() string {
	return ", on the %s of the month"
}
	
func (e en) CommaEveryX0Days() string {
	return ", every %s days"
}
	
func (e en) CommaBetweenDayX0AndX1OfTheMonth() string {
	return ", between day %s and %s of the month"
}
	
func (e en) CommaOnDayX0OfTheMonth() string {
	return ", on day %s of the month"
}
	
func (e en) CommaEveryHour() string {
	return ", every hour"
}
	
func (e en) CommaEveryX0Years() string {
	return ", every %s years"
}
	
func (e en) CommaStartingX0() string {
	return ", starting %s"
}
	
func (e en) DaysOfTheWeek() []string {
	return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
}
	
func (e en) MonthsOfTheYear()[]string {
	return []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
}

