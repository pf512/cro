package i18n

type Locale interface{
	// TODO: Circle back and use null/undefined aware types for optionals below in TypeScript 2.0: https://github.com/Microsoft/TypeScript/pull/7140
	// TODO: These locale translations would be a good use for ES6 template strings except we sometimes concatenate multiple transactions together before
	//       doing the actual template replacement.

	SetPeriodBeforeTime() bool
	Pm() string
	Am() string
	Use24HourTimeFormatByDefault() bool
	AnErrorOccuredWhenGeneratingTheExpressionD() string
	EveryMinute() string
	EveryHour() string
	AtSpace() string
	EveryMinuteBetweenX0AndX1() string
	At() string
	SpaceAnd() string
	EverySecond() string
	EveryX0Seconds() string
	SecondsX0ThroughX1PastTheMinute() string
	AtX0SecondsPastTheMinute() string
	AtX0SecondsPastTheMinuteGt20() string // optional
	EveryX0Minutes() string
	MinutesX0ThroughX1PastTheHour() string
	AtX0MinutesPastTheHour() string
	AtX0MinutesPastTheHourGt20() string // optional
	EveryX0Hours() string
	BetweenX0AndX1() string
	AtX0() string
	CommaEveryDay() string
	CommaEveryX0DaysOfTheWeek() string
	CommaX0ThroughX1() string
	CommaMonthX0ThroughMonthX1() string // optional
	CommaYearX0ThroughYearX1() string // optional
	First() string
	Second() string
	Third() string
	Fourth() string
	Fifth() string
	CommaOnThe() string
	SpaceX0OfTheMonth() string
	LastDay() string
	CommaOnTheLastX0OfTheMonth() string
	CommaOnlyOnX0() string
	CommaAndOnX0() string
	CommaEveryX0Months() string
	CommaOnlyInX0() string
	CommaOnlyInMonthX0() string
	CommaOnlyInYearX0() string
	CommaOnTheLastDayOfTheMonth() string
	CommaOnTheLastWeekdayOfTheMonth() string
	CommaDaysBeforeTheLastDayOfTheMonth() string
	FirstWeekday() string
	WeekdayNearestDayX0() string
	CommaOnTheX0OfTheMonth() string
	CommaEveryX0Days() string
	CommaBetweenDayX0AndX1OfTheMonth() string
	CommaOnDayX0OfTheMonth() string
	CommaEveryX0Years() string
	CommaStartingX0() string
	DayX0() string
	DaysOfTheWeek() []string
	MonthsOfTheYear() []string
}

