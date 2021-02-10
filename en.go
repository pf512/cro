package main

// English

//import { Locale } from "../locale"

type en struct {
}

func (e en) am() string {
	return "AM"
}

func (e en) pm() string {
	return "PM"
}

func (e en) commaOnlyInMonthX0() string {
	return ""
}

func (e en) commaOnlyInYearX0() string {
	return ""
}

func (e en) dayX0() string {
	return ""
}

func (e en) setPeriodBeforeTime() bool {
	return false
}



func (e en) atX0SecondsPastTheMinuteGt20() string {
		return ""
	}

func (e en) atX0MinutesPastTheHourGt20() string {
	return ""
}

func (e en) commaMonthX0ThroughMonthX1() string {
	return ""
}

func (e en) commaYearX0ThroughYearX1() string {
	return ""
}

func (e en) use24HourTimeFormatByDefault() bool {
	return false
}

func (e en) anErrorOccuredWhenGeneratingTheExpressionD() string {
	return "An error occured when generating the expression description.  Check the cron expression syntax."
}
	
func (e en) everyMinute() string {
	return "every minute"
}

func (e en) everyHour() string {
	return "every hour"
}

func (e en) atSpace() string {
	return "At "
}
	
func (e en) everyMinuteBetweenX0AndX1() string {
	return "Every minute between %s and %s"
}
	
func (e en) at() string {
	return "At"
}
	
func (e en) spaceAnd() string {
	return " and"
}
	
func (e en) everySecond() string {
	return "every second"
}
	
func (e en) everyX0Seconds() string {
	return "every %s seconds"
}
	
func (e en) secondsX0ThroughX1PastTheMinute() string {
	return "seconds %s through %s past the minute"
}
	
func (e en) atX0SecondsPastTheMinute() string {
	return "at %s seconds past the minute"
}
	
func (e en) everyX0Minutes() string {
	return "every %s minutes"
}
	
func (e en) minutesX0ThroughX1PastTheHour() string {
	return "minutes %s through %s past the hour"
}
	
func (e en) atX0MinutesPastTheHour() string {
	return "at %s minutes past the hour"
}
	
func (e en) everyX0Hours() string {
	return "every %s hours"
}
	
func (e en) betweenX0AndX1() string {
	return "between %s and %s"
}
	
func (e en) atX0() string {
	return "at %s"
}
	
func (e en) commaEveryDay() string {
	return ", every day"
}
	
func (e en) commaEveryX0DaysOfTheWeek() string {
	return ", every %s days of the week"
}
	
func (e en) commaX0ThroughX1() string {
	return ", %s through %s"
}
	
func (e en) first() string {
	return "first"
}
	
func (e en) second() string {
	return "second"
}
	
func (e en) third() string {
	return "third"
}
	
func (e en) fourth() string {
	return "fourth"
}
	
func (e en) fifth() string {
	return "fifth"
}
	
func (e en) commaOnThe() string {
	return ", on the "
}
	
func (e en) spaceX0OfTheMonth() string {
	return " %s of the month"
}
	
func (e en) lastDay() string {
	return "the last day"
}
	
func (e en) commaOnTheLastX0OfTheMonth() string {
	return ", on the last %s of the month"
}
	
func (e en) commaOnlyOnX0() string {
	return ", only on %s"
}
	
func (e en) commaAndOnX0() string {
	return ", and on %s"
}
	
func (e en) commaEveryX0Months() string {
	return ", every %s months"
}
	
func (e en) commaOnlyInX0() string {
	return ", only in %s"
}
	
func (e en) commaOnTheLastDayOfTheMonth() string {
	return ", on the last day of the month"
}
	
func (e en) commaOnTheLastWeekdayOfTheMonth() string {
	return ", on the last weekday of the month"
}
	
func (e en) commaDaysBeforeTheLastDayOfTheMonth() string {
	return ", %s days before the last day of the month"
}
	
func (e en) firstWeekday() string {
	return "first weekday"
}
	
func (e en) weekdayNearestDayX0() string {
	return "weekday nearest day %s"
}
	
func (e en) commaOnTheX0OfTheMonth() string {
	return ", on the %s of the month"
}
	
func (e en) commaEveryX0Days() string {
	return ", every %s days"
}
	
func (e en) commaBetweenDayX0AndX1OfTheMonth() string {
	return ", between day %s and %s of the month"
}
	
func (e en) commaOnDayX0OfTheMonth() string {
	return ", on day %s of the month"
}
	
func (e en) commaEveryHour() string {
	return ", every hour"
}
	
func (e en) commaEveryX0Years() string {
	return ", every %s years"
}
	
func (e en) commaStartingX0() string {
	return ", starting %s"
}
	
func (e en) daysOfTheWeek() []string {
	return []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
}
	
func (e en) monthsOfTheYear()[]string {
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

