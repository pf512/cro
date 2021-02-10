// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// https://www.typescriptlang.org/play?#code/PTAEHUFMBsGMHsC2lQBd5oBYoCoE8AHSAZVgCcBLA1UABWgEM8BzM+AVwDsATAGiwoBnUENANQAd0gAjQRVSQAUCEmYKsTKGYUAbpGF4OY0BoadYKdJMoL+gzAzIoz3UNEiPOofEVKVqAHSKymAAmkYI7NCuqGqcANag8ABmIjQUXrFOKBJMggBcISGgoAC0oACCoASMFmgY7p7ehCTkVOle4jUMdRLYTqCc8LEZzCZmoNJODPHFZZXVtZYYkAAeRJTInDQS8po+rf40gnjbDKv8LqD2jpbYoACqAEoAMsK7sUmxkGSCc+VVQQuaTwVb1UBrDYULY7PagbgUZLJH6QbYmJAECjuMigZEMVDsJzCFLNXxtajBBCcQQ0MwAUVWDEQNUgADVHBQGNJ3KAALygABEAAkYNAMOB4GRogLFFTBPB3AExcwABT0xnM9zsyhc9wASmCKhwDQ8ZC8iElzhB7Bo3zcZmY7AYzEg-Fg0HUiS58D0Ii8AoZTJZggFSRxAvADlQAHJhAA5SASAVBFQAeW+ZF2gldWkgx1QjgUrmkeFATgtOlGWH0KAQiBhwiudokkuiIgMHBx3RYbC43CCJRKOkc8Ly+VAAG9QABteKQPDjmmUTjMAC64847EQ0h+oAAvnzJ4pB4PiA84+OAAy8Y8n0AAWVTF9AAEYb3fvA86eOAEzvu-gHSAAi44AMz-ieOBCg844ACwQYOABiTwAJLjgArAhJTEBUODjgAbAhe4ANy3iUySSqAKruDQ3BMH6o54IIepHh+coKpASrwKqdF4AaJ57kAA
package main

import (
	"regexp"
	"strconv"
	"strings"
)

type ExpressionDescriptor struct {

	locales map[string]Locale
	specialCharacters []string

	expression string
	expressionParts []string
	options Options
	i18n Locale
}

func GetExpressionDescriptor() ExpressionDescriptor {
	var ed ExpressionDescriptor
	var enLocale enLocaleLoader
	ed.initialize(enLocale)

	return ed
}

func (ed *ExpressionDescriptor) initialize(localesLoader LocaleLoader) {
	ed.specialCharacters = []string{"/", "-", ",", "*"}

	throwExceptionOnParseError := true
	verbose := false
	dayOfWeekStartIndexZero := true
	use24HourTimeFormat := false
	locale := "en"

	var options Options
	options.throwExceptionOnParseError = throwExceptionOnParseError
	options.verbose = verbose
	options.dayOfWeekStartIndexZero = dayOfWeekStartIndexZero
	options.use24HourTimeFormat = use24HourTimeFormat
	options.locale = locale

	ed.options = options

	ed.locales = make(map[string]Locale)

	// Load locales
	localesLoader.load(ed.locales)
}

func (ed *ExpressionDescriptor) constructor(expression string) {
	ed.expression = expression

	if _, ok := ed.locales[ed.options.locale]; ok {
		ed.i18n = ed.locales[ed.options.locale]
	} else {
		// fall back to English
		println("Locale "+ ed.options.locale +" could not be found; falling back to 'en'")
		ed.i18n = ed.locales["en"]
	}

	if ed.options.use24HourTimeFormat == false {
		// if use24HourTimeFormat not specified, set based on locale default
		ed.options.use24HourTimeFormat = ed.i18n.use24HourTimeFormatByDefault()
	}


}

func (ed *ExpressionDescriptor) setVerbose(verbose bool) {
	ed.options.verbose = verbose
}

/**
 * Converts a cron expression into a description a human can read
 * @static
 * @param {string} expression - The cron expression
 * @param {IOptions} [{
 *         throwExceptionOnParseError = true,
 *         casingType = CasingTypeEnum.Sentence,
 *         verbose = false,
 *         dayOfWeekStartIndexZero = true,
 *         use24HourTimeFormat = false,
 *         locale = 'en'
 *     }={}]
 * @returns {string}
 */

type Options struct {
	throwExceptionOnParseError bool
	verbose bool
	dayOfWeekStartIndexZero bool
	use24HourTimeFormat bool
	locale string
}

func (ed *ExpressionDescriptor) toString(expression string) string {
	// We take advantage of Destructuring Object Parameters (and defaults) in TS/ES6 and now we will reassemble back to
	// an Options type so we can pass around options with ease.

	if expression == "" {
		return ""
	}

	ed.constructor(expression)

	return ed.getFullDescription()

}


func (ed *ExpressionDescriptor) getFullDescription() string {

	description := ""

	parser := GetCronParser(ed.expression, ed.options.dayOfWeekStartIndexZero)
	ed.expressionParts = parser.parse()

	timeSegment := ed.getTimeOfDayDescription()
	dayOfMonthDesc := ed.getDayOfMonthDescription()
	monthDesc := ed.getMonthDescription()
	dayOfWeekDesc := ed.getDayOfWeekDescription()
	yearDesc := ed.getYearDescription()


	description += timeSegment + dayOfMonthDesc + dayOfWeekDesc + monthDesc + yearDesc
	description = ed.transformVerbosity(description, ed.options.verbose)

	// uppercase first character
	description = strings.ToUpper(description[:1]) + description[1:]


	/*catch (ex) {
		if (!this.options.throwExceptionOnParseError) {
			description = this.i18n.anErrorOccuredWhenGeneratingTheExpressionD();
		} else {
			throw `${ex}`;
		}
	}
	 */
	return description
}

func (ed *ExpressionDescriptor) getTimeOfDayDescription() string {

	secondsExpression := ed.expressionParts[0]
	minuteExpression := ed.expressionParts[1]
	hourExpression := ed.expressionParts[2]

	description := ""

	stringUtilities := GetStringUtilities()
	// handle special cases first
	if !stringUtilities.containsAny(minuteExpression, ed.specialCharacters) && !stringUtilities.containsAny(hourExpression, ed.specialCharacters) && !stringUtilities.containsAny(secondsExpression, ed.specialCharacters) {
		// specific time of day (i.e. 10 14)
		description += ed.i18n.atSpace() + ed.formatTime(hourExpression, minuteExpression, secondsExpression)
	} else if secondsExpression =="" && strings.Index(minuteExpression,"-") > -1 && !(strings.Index(minuteExpression,",") > -1) && !(strings.Index(minuteExpression,"/") > -1) && !stringUtilities.containsAny(hourExpression, ed.specialCharacters) {
		// minute range in single hour (i.e. 0-10 11)
		minuteParts := strings.Split(minuteExpression,"-")
		description += stringUtilities.format(ed.i18n.everyMinuteBetweenX0AndX1(), ed.formatTime(hourExpression, minuteParts[0], ""), ed.formatTime(hourExpression, minuteParts[1], ""))
	} else if secondsExpression == "" && strings.Index(hourExpression,",") > -1 && strings.Index(hourExpression,"-") == -1 && strings.Index(hourExpression,"/") == -1 && !stringUtilities.containsAny(minuteExpression, ed.specialCharacters) {
		// hours list with single minute (i.e. 30 6,14,16)
		hourParts := strings.Split(hourExpression, ",")
		description += ed.i18n.at()

		for i := 0; i < len(hourParts); i++ {
			description += " ";
			description += ed.formatTime(hourParts[i], minuteExpression, "")

			if (i < len(hourParts) - 2) {
			description += ","
			}

			if (i == len(hourParts) - 2) {
			description += ed.i18n.spaceAnd()
			}
		}

	} else {
		// default time description

		secondsDescription := ed.getSecondsDescription()
		minutesDescription := ed.getMinutesDescription()
		hoursDescription := ed.getHoursDescription()

		description += secondsDescription

		if len(description) > 0 && len(minutesDescription) > 0 {
			description += ", "
		}

		description += minutesDescription

		if minutesDescription == hoursDescription {
			return description
		}

		if len(description) > 0 && len(hoursDescription) > 0 {
			description += ", "
		}

		description += hoursDescription

	}

	return description
}


func (ed *ExpressionDescriptor) getSecondsDescription() string {
	var stringUtilities StringUtilities

	description := ed.getSegmentDescription(
		ed.expressionParts[0],
		ed.i18n.everySecond(),
		func(s string) string { return s },
		func(s string) string { return stringUtilities.format(ed.i18n.everyX0Seconds(), s)},
		func(s string) string { return ed.i18n.secondsX0ThroughX1PastTheMinute()},
		func(s string) string { if s == "0" {
				return ""
			}

			tmp, _ := strconv.Atoi(s)
			if tmp < 20 {
				return ed.i18n.atX0SecondsPastTheMinute()
			}

			if ed.i18n.atX0SecondsPastTheMinuteGt20() != "" {
				return ed.i18n.atX0SecondsPastTheMinuteGt20()
			}

			return ed.i18n.atX0SecondsPastTheMinute()
		},
	)

	return description
}

func (ed *ExpressionDescriptor) getMinutesDescription() string {
	var stringUtilities StringUtilities

	secondsExpression := ed.expressionParts[0]
	hourExpression := ed.expressionParts[2]

	description := ed.getSegmentDescription(
		ed.expressionParts[1],
		ed.i18n.everyMinute(),
		func(s string) string { return s },
		func(s string) string { return stringUtilities.format(ed.i18n.everyX0Minutes(), s)},
		func(s string) string { return ed.i18n.minutesX0ThroughX1PastTheHour()},
		func(s string) string { if s == "0" && strings.Index(hourExpression,"/") == -1 && secondsExpression == "" {
				return ed.i18n.everyHour()
			}

			tmp, _ := strconv.Atoi(s)
			if tmp < 20 {
				return ed.i18n.atX0MinutesPastTheHour()
			}

			if ed.i18n.atX0MinutesPastTheHourGt20() != "" {
				return ed.i18n.atX0MinutesPastTheHourGt20()
			}

			return ed.i18n.atX0MinutesPastTheHour()
		},
	)

	return description
}

func (ed *ExpressionDescriptor) getHoursDescription() string {
	var stringUtilities StringUtilities

	expression := ed.expressionParts[2]
	description := ed.getSegmentDescription(
		expression,
		ed.i18n.everyHour(),
		func(s string) string { return ed.formatTime(s, "0", "")},
		func(s string) string { return stringUtilities.format(ed.i18n.everyX0Hours(), s)},
		func(s string) string { return ed.i18n.betweenX0AndX1()},
		func(s string) string { return ed.i18n.atX0()},
	)

	return description
}

func (ed *ExpressionDescriptor) getDayOfWeekDescription() string {
	var stringUtilities StringUtilities

	daysOfWeekNames := ed.i18n.daysOfTheWeek()

	description := ""
	if ed.expressionParts[5] == "*" {
		// DOW is specified as * so we will not generate a description and defer to DOM part.
		// Otherwise, we could get a contradiction like "on day 1 of the month, every day"
		// or a dupe description like "every day, every day".
		description = ""
	} else {
		description = ed.getSegmentDescription(
			ed.expressionParts[5],
			ed.i18n.commaEveryDay(),
			func(s string) string {
				exp := s
				if strings.Index(s, "#") > -1 {
					exp = s[:strings.Index(s, "#")]
				} else if strings.Index(s, "L") > -1 {
					exp = strings.Replace(exp, "L", "", -1)
				}
				tmpInt, _ := strconv.Atoi(exp)
				return daysOfWeekNames[tmpInt]
			},
			func(s string) string {
				tmpInt, _ := strconv.Atoi(s)
				if (tmpInt == 1) {
					// rather than "every 1 days" just return empty string
					return ""
				}

				return stringUtilities.format(ed.i18n.commaEveryX0DaysOfTheWeek(), s);

			},
			func(s string) string { return ed.i18n.commaX0ThroughX1()},
			func(s string) string {
				format := ""
				if strings.Index(s,"#") > -1 {
					dayOfWeekOfMonthNumber := s[:(strings.Index(s,"#") + 1)]
					dayOfWeekOfMonthDescription := ""
					switch dayOfWeekOfMonthNumber {
						case "1":
							dayOfWeekOfMonthDescription = ed.i18n.first()
							break
						case "2":
							dayOfWeekOfMonthDescription = ed.i18n.second()
							break
						case "3":
							dayOfWeekOfMonthDescription = ed.i18n.third()
							break
						case "4":
							dayOfWeekOfMonthDescription = ed.i18n.fourth()
							break
						case "5":
							dayOfWeekOfMonthDescription = ed.i18n.fifth()
							break
					}

					format = ed.i18n.commaOnThe() + dayOfWeekOfMonthDescription + ed.i18n.spaceX0OfTheMonth()
				} else if strings.Index(s,"L") > -1 {
					format = ed.i18n.commaOnTheLastX0OfTheMonth()
				} else {
					// If both DOM and DOW are specified, the cron will execute at either time.
					domSpecified := ed.expressionParts[3] != "*"

					if domSpecified {
						format = ed.i18n.commaAndOnX0()
					} else {
						format = ed.i18n.commaOnlyOnX0()
					}
				}
				return format
			},
		)
	}

	return description
}

func (ed *ExpressionDescriptor) getMonthDescription() string {
	var stringUtilities StringUtilities

	monthNames := ed.i18n.monthsOfTheYear()

	description	:= ed.getSegmentDescription(
		ed.expressionParts[4],
		"",
		func(s string) string {
			tmpInt, _ := (strconv.Atoi(s))
			return monthNames[tmpInt - 1]
		},
		func(s string) string {
			tmpInt, _ := strconv.Atoi(s)
			if tmpInt == 1 {
				// rather than "every 1 months" just return empty string
				return ""
			} else {
				return stringUtilities.format(ed.i18n.commaEveryX0Months(), s)
			}
		},
		func(s string) string {
			if ed.i18n.commaMonthX0ThroughMonthX1() != "" {
				return ed.i18n.commaMonthX0ThroughMonthX1()
			}

			return ed.i18n.commaX0ThroughX1()
		},
		func(s string) string {
			if ed.i18n.commaOnlyInMonthX0() != "" {
				return ed.i18n.commaOnlyInMonthX0()
			}
			return ed.i18n.commaOnlyInX0()
		},
	)

	return description
}

func (ed *ExpressionDescriptor) getDayOfMonthDescription() string {
	var stringUtilities StringUtilities

	description := ""
	expression := ed.expressionParts[3]

	switch expression {

		case "L":
			description = ed.i18n.commaOnTheLastDayOfTheMonth()
			break
		case "WL":
		case "LW":
			description = ed.i18n.commaOnTheLastWeekdayOfTheMonth()
			break
		default:
			re := regexp.MustCompile("(\\d{1,2}W)|(W\\d{1,2})")
			weekDayNumberMatches := re.FindAllString(expression, -1) // i.e. 3W or W2
			if len(weekDayNumberMatches) > 0 {
				tmp := strings.ReplaceAll(weekDayNumberMatches[0], "W", "")
				tmpInt, _ := strconv.Atoi(tmp)
				dayNumber := tmpInt

				var dayString string
				if dayNumber == 1 {
					dayString = ed.i18n.firstWeekday()
				} else {
					dayString = stringUtilities.format(ed.i18n.weekdayNearestDayX0(), strconv.Itoa(dayNumber))
				}

				description = stringUtilities.format(ed.i18n.commaOnTheX0OfTheMonth(), dayString)
				break
			} else {
				// Handle "last day offset" (i.e. L-5:  "5 days before the last day of the month")
				re := regexp.MustCompile("L-(\\d{1,2})")
				lastDayOffSetMatches := re.FindAllString(expression, -1)
				if len(lastDayOffSetMatches) > 0 {
					offSetDays := lastDayOffSetMatches[1]
					description = stringUtilities.format(ed.i18n.commaDaysBeforeTheLastDayOfTheMonth(), offSetDays)
					break
				} else if expression == "*" && ed.expressionParts[5] != "*" {
				// * dayOfMonth and dayOfWeek specified so use dayOfWeek verbiage instead
				return ""
				} else {
				description = ed.getSegmentDescription(
					expression,
					ed.i18n.commaEveryDay(),
					func(s string) string {
						if s == "L" {
							return ed.i18n.lastDay()
						} else if ed.i18n.dayX0() != "" {
							return stringUtilities.format(ed.i18n.dayX0(), s)
						}

						return s
					},
					func(s string) string {
						if s == "1" {
							return ed.i18n.commaEveryDay()
						}
						return ed.i18n.commaEveryX0Days()
					},
					func(s string) string {
						return ed.i18n.commaBetweenDayX0AndX1OfTheMonth()
					},
					func(s string) string {
						return ed.i18n.commaOnDayX0OfTheMonth()
					},
				)
			}
		break
		}
	}

	return description
}

func (ed *ExpressionDescriptor) getYearDescription() string {
	var stringUtilities StringUtilities

	description := ed.getSegmentDescription(
		ed.expressionParts[6],
		"",
		func(s string) string {
			re := regexp.MustCompile("^\\d+$")
			if re.MatchString(s) {
				return s //todo
				//return Date(parseInt(s), 1).getFullYear().toString()
			}

			return s
		},
		func(s string) string {
			return stringUtilities.format(ed.i18n.commaEveryX0Years(), s)
		},
		func(s string) string {
			if ed.i18n.commaYearX0ThroughYearX1() != "" {
				return ed.i18n.commaYearX0ThroughYearX1()
			}
			return ed.i18n.commaX0ThroughX1()
		},
		func(s string) string {
			if ed.i18n.commaOnlyInYearX0() != "" {
				return ed.i18n.commaOnlyInYearX0()
			}
 			return ed.i18n.commaOnlyInX0()
		},
	)

	return description
}

type stringFunction func(string) string

func (ed *ExpressionDescriptor) getSegmentDescription(expression string,
	allDescription string,
	getSingleItemDescription stringFunction,
	getIncrementDescriptionFormat stringFunction,
	getRangeDescriptionFormat stringFunction,
	getDescriptionFormat stringFunction,
	) string {
	var stringUtilities StringUtilities

	description := ""
	doesExpressionContainIncrement := strings.Index(expression,"/") > -1
	doesExpressionContainRange := strings.Index(expression, "-") > -1
	doesExpressionContainMultipleValues := strings.Index(expression, ",") > -1

	if expression == "" {
		// Empty
		description = ""
	} else if expression == "*" {
		// * (All)
		description = allDescription
	} else if (!doesExpressionContainIncrement && !doesExpressionContainRange && !doesExpressionContainMultipleValues) {
		// Simple
		description = stringUtilities.format(getDescriptionFormat(expression), getSingleItemDescription(expression))
	} else if (doesExpressionContainMultipleValues) {
		// Multiple Values

		segments := strings.Split(expression, ",")
		descriptionContent := ""

		for i := 0; i < len(segments); i++ {
			if i > 0 && len(segments) > 2 {
				descriptionContent += ","

				if i < len(segments)-1 {
					descriptionContent += " "
				}
			}

			if i > 0 && len(segments) > 1 && (i == len(segments) - 1 || len(segments) == 2) {
				descriptionContent += ed.i18n.spaceAnd()+" "
			}

			if strings.Index(segments[i],"/") > -1 || strings.Index(segments[i],"-") > -1 {
				// Multiple Values with Increment or Range

				isSegmentRangeWithoutIncrement := strings.Index(segments[i],"-") > -1 && strings.Index(segments[i],"/") == -1

				var currentDescriptionContent string
				if isSegmentRangeWithoutIncrement {
					currentDescriptionContent = ed.getSegmentDescription(
						segments[i],
						allDescription,
						getSingleItemDescription,
						getIncrementDescriptionFormat,
						func (s string) string { return ed.i18n.commaX0ThroughX1() },
						getDescriptionFormat,
					)
				} else {
					currentDescriptionContent = ed.getSegmentDescription(
						segments[i],
						allDescription,
						getSingleItemDescription,
						getIncrementDescriptionFormat,
						getRangeDescriptionFormat,
						getDescriptionFormat,
					)
				}

				if isSegmentRangeWithoutIncrement {
					currentDescriptionContent = strings.Replace(currentDescriptionContent, ", ", "", -1)
				}

				descriptionContent += currentDescriptionContent

			} else if !doesExpressionContainIncrement {
				descriptionContent += getSingleItemDescription(segments[i])
			} else {
				descriptionContent += ed.getSegmentDescription(
					segments[i],
					allDescription,
					getSingleItemDescription,
					getIncrementDescriptionFormat,
					getRangeDescriptionFormat,
					getDescriptionFormat,
				)
			}
		}

		if !doesExpressionContainIncrement {
			description = stringUtilities.format(getDescriptionFormat(expression), descriptionContent)
		} else {
			description = descriptionContent
		}
	} else if doesExpressionContainIncrement {
		// Increment

		segments := strings.Split(expression,"/")
		description = stringUtilities.format(getIncrementDescriptionFormat(segments[1]), segments[1])

		if strings.Index(segments[0],"-") > -1 {
		// Range with Increment (ex: 2-59/3 )

		rangeSegmentDescription := ed.generateRangeSegmentDescription(
		segments[0],
		getRangeDescriptionFormat,
		getSingleItemDescription,
		)

		if strings.Index(rangeSegmentDescription,", ") != 0 {
			description += ", "
		}

		description += rangeSegmentDescription
	} else if strings.Index(segments[0],"*") == -1 {
		rangeItemDescription := stringUtilities.format(
		getDescriptionFormat(segments[0]),
		getSingleItemDescription(segments[0]),
		)

		// remove any leading comma
		rangeItemDescription = strings.Replace(rangeItemDescription,", ", "",-1)

		description += stringUtilities.format(ed.i18n.commaStartingX0(), rangeItemDescription)
	}
	} else if doesExpressionContainRange {
		// Range

		description = ed.generateRangeSegmentDescription(
		expression,
		getRangeDescriptionFormat,
		getSingleItemDescription,
		)
	}

	return description
}

func (ed *ExpressionDescriptor) formatTime(hourExpression string, minuteExpression string, secondExpression string) string {
	hour,_ := strconv.Atoi(hourExpression)
	period := ""
	setPeriodBeforeTime := false
	if !ed.options.use24HourTimeFormat {
		setPeriodBeforeTime = ed.i18n.setPeriodBeforeTime()
		if setPeriodBeforeTime {
			period = ed.getPeriod(hour)+" "
		} else {
			period = " " + ed.getPeriod(hour)
		}
	}

	if hour > 12 {
		hour -= 12
	}

	if hour == 0 {
		hour = 12
	}

	//minute := minuteExpression
	second := ""
	if secondExpression != "" {
		second = ":"+("00" + secondExpression)[len(secondExpression):]
	}


	var tmpString string

	hourString := strconv.Itoa(hour)
	if setPeriodBeforeTime {
		tmpString = period
	}

	tmpString += ("00" + hourString)[len(hourString):]+":"+("00" + minuteExpression)[len(minuteExpression):]+second
	if !setPeriodBeforeTime {
		tmpString += period
	}

	return tmpString

}

func (ed *ExpressionDescriptor)  transformVerbosity(description string, useVerboseFormat bool) string {
	if !useVerboseFormat {
		description = strings.Replace(description,", "+ed.i18n.everyMinute(), "", -1)
		description = strings.Replace(description,", "+ed.i18n.everyHour(), "", -1)
		description = strings.Replace(description, ed.i18n.commaEveryDay(), "", -1)

		re := regexp.MustCompile(`\, ?$`)
		re.ReplaceAllString(description, "")

		/*
		description = strings.Replace(description.replace(new RegExp(`, ${this.i18n.everyMinute()}`, "g"), "")
		description = description.replace(new RegExp(`, ${this.i18n.everyHour()}`, "g"), "")
		description = description.replace(new RegExp(this.i18n.commaEveryDay(), "g"), "")
		description = description.replace(/\, ?$/, "")
		 */
	}
	return description
}

func (ed *ExpressionDescriptor) getPeriod(hour int) string {

	if hour >= 12 {
		return ed.i18n.pm()
	}

	return ed.i18n.am()

}


func (ed *ExpressionDescriptor) generateRangeSegmentDescription(
rangeExpression string,
getRangeDescriptionFormat stringFunction,
getSingleItemDescription stringFunction,
) string {
	var stringUtilities StringUtilities

	description	:= ""
	rangeSegments := strings.Split(rangeExpression,"-")
	rangeSegment1Description := getSingleItemDescription(rangeSegments[0])
	rangeSegment2Description := getSingleItemDescription(rangeSegments[1])
	rangeSegment2Description = strings.Replace(rangeSegment2Description,":00", ":59",-1)
	rangeDescriptionFormat := getRangeDescriptionFormat(rangeExpression)
	description += stringUtilities.format(rangeDescriptionFormat, rangeSegment1Description, rangeSegment2Description)

	return description
}
