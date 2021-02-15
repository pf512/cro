package cronutils

import (
	"testing"
)

func TestEverySecond(t *testing.T) {

	testString := "* * * * * *"
	expected := "Every second"

	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}

}

func TestEveryMinute(t *testing.T) {

	testString := "* * * * *"
	expected := "Every minute"

	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourVerbose(t *testing.T) {

	testString := "0 * * * *"
	expected := "Every hour, every day"

	SetVerboseOn()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteVerbose(t *testing.T) {

	testString := "* * * * *"
	expected := "Every minute, every hour, every day"

	SetVerboseOn()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}



func TestEveryMinute2(t *testing.T) {

	testString := "*/1 * * * *"
	expected := "Every minute"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEvery5Minutes(t *testing.T) {

	testString := "*/5 * * * *"
	expected := "Every 5 minutes"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinute3(t *testing.T) {

	testString := "0 0/1 * * * ?"
	expected := "Every minute"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestEveryHour(t *testing.T) {

	testString := "0 0 * * * ?"
	expected := "Every hour"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHour2(t *testing.T) {

	testString := "0 0 0/1 * * ?"
	expected := "Every hour"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteMarch(t *testing.T) {

	testString := "* * * 3 *"
	expected := "Every minute, only in March"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteInMarchAndJune(t *testing.T) {

	testString := "* * * 3,6 *"
	expected := "Every minute, only in March and June"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestEverySecond2013(t *testing.T) {

	testString := "* * * * * * 2013"
	expected := "Every second, only in 2013"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteOnly2013(t *testing.T) {

	testString := "* * * * * 2013"
	expected := "Every minute, only in 2013"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteIn2013And2014(t *testing.T) {

	testString := "* * * * * 2013,2014"
	expected := "Every minute, only in 2013 and 2014"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEvery45Seconds(t *testing.T) {

	testString := "*/45 * * * * *"
	expected := "Every 45 seconds"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEvery5Minutes2(t *testing.T) {

	testString := "*/5 * * * *"
	expected := "Every 5 minutes"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestEvery10Minutes(t *testing.T) {

	testString := "*/10 * * * *"
	expected := "Every 10 minutes"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEvery5Minutes3(t *testing.T) {

	testString := "0 */5 * * * *"
	expected := "Every 5 minutes"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourMisc(t *testing.T) {

	testString := "0 9-17 * * *"
	expected := "Every hour, between 09:00 AM and 05:59 PM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourTuesThruSat(t *testing.T) {

	testString := "0 * ? * 2/1 *"
	expected := "Every hour, Tuesday through Saturday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourMonThruSat(t *testing.T) {

	testString := "0 * ? * 1/1"
	expected := "Every hour, Monday through Saturday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourTuesThruSat2(t *testing.T) {

	testString := "0 * ? * 2/1"
	expected := "Every hour, Tuesday through Saturday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestSpecificHourMinuteWedThruSat(t *testing.T) {

	testString := "0 52 13 ? * 3/1"
	expected := "At 01:52 PM, Wednesday through Saturday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestSpecificHourMonThuFri(t *testing.T) {

	testString := "0 23 ? * MON-FRI"
	expected := "At 11:00 PM, Monday through Friday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestSpecificHourMinMonThuFri(t *testing.T) {

	testString := "30 11 * * 1-5"
	expected := "At 11:30 AM, Monday through Friday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteBetween(t *testing.T) {

	testString := "0-10 11 * * *"
	expected := "Every minute between 11:00 AM and 11:10 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestSpecificTimeJanThruMar(t *testing.T) {

	testString := "23 12 * Jan-Mar *"
	expected := "At 12:23 PM, January through March"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestSpecificTimeJanThruFeb(t *testing.T) {

	testString := "23 12 * JAN-FEB *"
	expected := "At 12:23 PM, January through February"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestOneMinutePast(t *testing.T) {

	testString := "1 1,3-4 * * *"
	expected := "At 1 minutes past the hour, at 01:00 AM and 03:00 AM through 04:59 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestZeroMinutesPastEvery4Hours(t *testing.T) {

	testString := "* 0 */4 * * *"
	expected := "Every second, at 0 minutes past the hour, every 4 hours"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEvery10Seconds2(t *testing.T) {

	testString := "*/10 0 * * * *"
	expected := "Every 10 seconds, at 0 minutes past the hour"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEverySecondBetween12(t *testing.T) {

	testString := "* 0 0 * * *"
	expected := "Every second, at 0 minutes past the hour, between 12:00 AM and 12:59 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteBetween12(t *testing.T) {

	testString := "* 0 * * *"
	expected := "Every minute, between 12:00 AM and 12:59 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEverySecondZeroMinutesPast(t *testing.T) {

	testString := "* 0 * * * *"
	expected := "Every second, at 0 minutes past the hour"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt1130(t *testing.T) {

	testString := "30 11 * * *"
	expected := "At 11:30 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt1223OnlyOnSunday(t *testing.T) {

	testString := "23 12 * * SUN"
	expected := "At 12:23 PM, only on Sunday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt020230(t *testing.T) {

	testString := "30 02 14 * * *"
	expected := "At 02:02:30 PM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt6AM(t *testing.T) {

	testString := "0 0 6 1/1 * ?"
	expected := "At 06:00 AM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt5MinutesPast(t *testing.T) {

	testString := "0 5 0/1 * * ?"
	expected := "At 5 minutes past the hour"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt946OnMonday(t *testing.T) {

	testString := "46 9 * * 1"
	expected := "At 09:46 AM, only on Monday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt946OnlyOnSunday(t *testing.T) {

	testString := "46 9 * * 7"
	expected := `At 09:46 AM, only on Sunday`

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt1223OnDay15(t *testing.T) {

	testString := "23 12 15 * *"
	expected := "At 12:23 PM, on day 15 of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt1223InJanuary(t *testing.T) {

	testString := "23 12 * JAN *"
	expected := "At 12:23 PM, only in January"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt1223InJanuary2(t *testing.T) {

	testString := "23 12 ? JAN *"
	expected := "At 12:23 PM, only in January"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt7(t *testing.T) {

	testString := "0 7 * * *"
	expected := `At 07:00 AM`

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt230And430(t *testing.T) {

	testString := "30 14,16 * * *"
	expected := "At 02:30 PM and 04:30 PM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestAt630And230And430(t *testing.T) {

	testString := "30 6,14,16 * * *"
	expected := "At 06:30 AM, 02:30 PM and 04:30 PM"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestEveryHourOnDay31AndOnMonday(t *testing.T) {

	testString := "0 * 31 * 1"
	expected := "Every hour, on day 31 of the month, and on Monday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

// ===

func TestEveryMinuteLastWeekday(t *testing.T) {

	testString :="* * LW * *"
	expected := "Every minute, on the last weekday of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteLastWeekday2(t *testing.T) {

	testString := "* * WL * *"
	expected := "Every minute, on the last weekday of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteFirstWeekday(t *testing.T) {

	testString :="* * 1W * *"
	expected := "Every minute, on the first weekday of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteWeekdayNearest(t *testing.T) {

	testString :="* * 13W * *"
	expected := "Every minute, on the weekday nearest day 13 of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteFirstWeekday2(t *testing.T) {

	testString :="* * W1 * *"
	expected := "Every minute, on the first weekday of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteWeekdayNearestDay5(t *testing.T) {

	testString :="* * 5W * *"
	expected := "Every minute, on the weekday nearest day 5 of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryMinuteOnWeekdayNearestDay5(t *testing.T) {

	testString :="* * W5 * *"
	expected := "Every minute, on the weekday nearest day 5 of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}


func TestEveryMinuteOnLastThursdayOfMonth(t *testing.T) {

	testString :="* * * * 4L"
	expected := "Every minute, on the last Thursday of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryFiveMinutesLastDayOnlyInJan(t *testing.T) {

	testString := "*/5 * L JAN *"
	expected := "Every 5 minutes, on the last day of the month, only in January"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt8OnDay15AndLastDayOfMonth(t *testing.T) {

	testString := "0 20 15,L * *"
	expected := "At 08:00 PM, on day 15 and the last day of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt8OnDay1Thru10And20(t *testing.T) {

	testString := "0 20 1-10,20-L * *"
	expected := "At 08:00 PM, on day 1 through 10 and 20 through the last day of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt1015OnSaturday(t *testing.T) {

	testString := "0 15 10 * * L"
	expected := "At 10:15 AM, only on Saturday"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt1015OnLastDayOfMonth(t *testing.T) {

	testString := "0 15 10 L * *"
	expected := "At 10:15 AM, on the last day of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt125DaysBeforeLastDayOfTheMonth(t *testing.T) {

	testString := "0 0 0 L-5 * ?"
	expected := "At 12:00 AM, 5 days before the last day of the month"

	SetVerboseOff()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestAt1223OnSecondSunday(t *testing.T) {

	testString := "23 12 * * 1#2"
	expected := "At 12:23 PM, on the second Sunday of the month"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEverySecondEvery2DaysOfTheWeekMThruF(t *testing.T) {

	testString := "* * * ? * 2-6/2"
	expected := "Every second, every 2 days of the week, Monday through Friday"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEverySecondOnlyOnSat(t *testing.T) {

	testString := "* * * ? * 7"
	expected := "Every second, only on Saturday"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEverySecondOnlySunMonTueWedThu(t *testing.T) {

	testString := "* * * ? * 1,2,3,4,5"
	expected := "Every second, only on Sunday, Monday, Tuesday, Wednesday, and Thursday"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourSundayThroughSat(t *testing.T) {

	testString := "0 * ? * 1/1"
	expected := "Every hour, Sunday through Saturday"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}

func TestEveryHourMondayThruSat(t *testing.T) {

	testString := "0 * ? * 2/1"
	expected := "Every hour, Monday through Saturday"

	SetVerboseOff()
	SetDayOfWeekStartIndexZeroFalse()
	received := ToString(testString)
	if received != expected {
		t.Fatalf(`ToString(%s) = %s, received = %s`, testString, expected, received)
	}
}
