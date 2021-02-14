package i18n

// Spanish

type es struct {
}

func (e es) languageCode() string {
	return "es"
}


func (e es) Am() string {
	return "AM"
}

func (e es) Pm() string {
	return "PM"
}

func (e es) CommaOnlyInMonthX0() string {
	return ""
}

func (e es) CommaOnlyInYearX0() string {
	return ""
}

func (e es) DayX0() string {
	return ""
}

func (e es) SetPeriodBeforeTime() bool {
	return false
}

func (e es) AtX0SecondsPastTheMinuteGt20() string {
return ""
}

func (e es) AtX0MinutesPastTheHourGt20() string {
return ""
}

func (e es) CommaMonthX0ThroughMonthX1() string {
return ""
}

func (e es) CommaYearX0ThroughYearX1() string {
return ""
}

func (e es) Use24HourTimeFormatByDefault() bool {
return false
}

func (e es) AnErrorOccuredWhenGeneratingTheExpressionD() string {
return "Ocurrió un error mientras se generaba la descripción de la expresión. Revise la sintaxis de la expresión de cron."
}

func (e es) At() string {
return "A las"
}

func (e es) AtSpace() string {
return "A las "
}

func (e es) AtX0() string {
return "a las %s"
}

func (e es) AtX0MinutesPastTheHour() string {
return "a los %s minutos de la hora"
}

func (e es) AtX0SecondsPastTheMinute() string {
return "a los %s segundos del minuto"
}

func (e es) BetweenX0AndX1() string {
return "entre las %s y las %s"
}

func (e es) CommaBetweenDayX0AndX1OfTheMonth() string {
return ", entre los días %s y %s del mes"
}

func (e es) CommaEveryDay() string {
return ", cada día"
}

func (e es) CommaEveryX0Days() string {
return ", cada %s días"
}

func (e es) CommaEveryX0DaysOfTheWeek() string {
return ", cada %s días de la semana"
}

func (e es) CommaEveryX0Months() string {
return ", cada %s meses"
}

func (e es) CommaOnDayX0OfTheMonth() string {
return ", el día %s del mes"
}

func (e es) CommaOnlyInX0() string {
return ", sólo en %s"
}

func (e es) CommaOnlyOnX0() string {
return ", sólo el %s"
}

func (e es) CommaAndOnX0() string {
return ", y el %s"
}

func (e es) CommaOnThe() string {
return ", en el "
}

func (e es) CommaOnTheLastDayOfTheMonth() string {
return ", en el último día del mes"
}

func (e es) CommaOnTheLastWeekdayOfTheMonth() string {
return ", en el último día de la semana del mes"
}

func (e es) CommaDaysBeforeTheLastDayOfTheMonth() string {
return ", %s días antes del último día del mes"
}

func (e es) CommaOnTheLastX0OfTheMonth() string {
return ", en el último %s del mes"
}

func (e es) CommaOnTheX0OfTheMonth() string {
return ", en el %s del mes"
}

func (e es) CommaX0ThroughX1() string {
return ", de %s a %s"
}

func (e es) EveryHour() string {
return "cada hora"
}

func (e es) EveryMinute() string {
return "cada minuto"
}

func (e es) EveryMinuteBetweenX0AndX1() string {
return "cada minuto entre las %s y las %s"
}

func (e es) EverySecond() string {
return "cada segundo"
}

func (e es) EveryX0Hours() string {
return "cada %s horas"
}

func (e es) EveryX0Minutes() string {
return "cada %s minutos"
}

func (e es) EveryX0Seconds() string {
return "cada %s segundos"
}

func (e es) Fifth() string {
return "quinto"
}

func (e es) First() string {
return "primero"
}

func (e es) FirstWeekday() string {
return "primer día de la semana"
}

func (e es) Fourth() string {
return "cuarto"
}

func (e es) MinutesX0ThroughX1PastTheHour() string {
return "del minuto %s al %s pasada la hora"
}

func (e es) Second() string {
return "segundo"
}

func (e es) SecondsX0ThroughX1PastTheMinute() string {
return "En los segundos %s al %s de cada minuto"
}

func (e es) SpaceAnd() string {
return " y"
}

func (e es) SpaceX0OfTheMonth() string {
return " %s del mes"
}

func (e es) LastDay() string {
return "el último día"
}

func (e es) Third() string {
return "tercer"
}

func (e es) WeekdayNearestDayX0() string {
return "día de la semana más próximo al %s"
}

func (e es) CommaEveryX0Years() string {
return ", cada %s años"
}

func (e es) CommaStartingX0() string {
return ", comenzando %s"
}

func (e es) DaysOfTheWeek() []string {
return []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"}
}

func (e es) MonthsOfTheYear() []string {
return []string{
"enero",
"febrero",
"marzo",
"abril",
"mayo",
"junio",
"julio",
"agosto",
"septiembre",
"octubre",
"noviembre",
"diciembre",
}

}
