// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join
// https://www.typescriptlang.org/play?#code/PTAEHUFMBsGMHsC2lQBd5oBYoCoE8AHSAZVgCcBLA1UABWgEM8BzM+AVwDsATAGiwoBnUENANQAd0gAjQRVSQAUCEmYKsTKGYUAbpGF4OY0BoadYKdJMoL+gzAzIoz3UNEiPOofEVKVqAHSKymAAmkYI7NCuqGqcANag8ABmIjQUXrFOKBJMggBcISGgoAC0oACCoASMFmgY7p7ehCTkVOle4jUMdRLYTqCc8LEZzCZmoNJODPHFZZXVtZYYkAAeRJTInDQS8po+rf40gnjbDKv8LqD2jpbYoACqAEoAMsK7sUmxkGSCc+VVQQuaTwVb1UBrDYULY7PagbgUZLJH6QbYmJAECjuMigZEMVDsJzCFLNXxtajBBCcQQ0MwAUVWDEQNUgADVHBQGNJ3KAALygABEAAkYNAMOB4GRogLFFTBPB3AExcwABT0xnM9zsyhc9wASmCKhwDQ8ZC8iElzhB7Bo3zcZmY7AYzEg-Fg0HUiS58D0Ii8AoZTJZggFSRxAvADlQAHJhAA5SASAVBFQAeW+ZF2gldWkgx1QjgUrmkeFATgtOlGWH0KAQiBhwiudokkuiIgMHBx3RYbC43CCJRKOkc8Ly+VAAG9QABteKQPDjmmUTjMAC64847EQ0h+oAAvnzJ4pB4PiA84+OAAy8Y8n0AAWVTF9AAEYb3fvA86eOAEzvu-gHSAAi44AMz-ieOBCg844ACwQYOABiTwAJLjgArAhJTEBUODjgAbAhe4ANy3iUySSqAKruDQ3BMH6o54IIepHh+coKpASrwKqdF4AaJ57kAA
package cronutils

import (
 "errors"
 "regexp"
 "strconv"
 "strings"
)


/**
 * Parses and normalizes a cron expression
 * @export
 * @class CronParser
 */
type CronParser struct {
 expression              string
 dayOfWeekStartIndexZero bool
}

func (parser *CronParser) constructor(expression string, dayOfWeekStartIndexZero bool) {
 parser.expression = expression
 parser.dayOfWeekStartIndexZero = dayOfWeekStartIndexZero
}

func GetCronParser(expression string, dayOfWeekStartIndexZero bool) CronParser {
 var cp CronParser
 cp.expression = expression
 cp.dayOfWeekStartIndexZero = dayOfWeekStartIndexZero

 return cp
}

/**
 * Parses and normalizes a cron expression into an array of strings
 * @returns {string[]}
 */
func (parser *CronParser) parse() ([]string, error) {

 parsed, err := parser.extractParts(parser.expression)

 if err != nil {
   println("cronParser parse")
   return parsed, err
 }

 parser.normalize(parsed)
 parser.validate(parsed)

 return parsed, err
}

func (parser *CronParser) extractParts(expression string) ([]string, error) {

 if expression == "" {
  return []string{}, errors.New("expression is empty")
 }

 // split on one or more spaces
 //let parsed: string[] = expression.trim().split(/[ ]+/);
 re := regexp.MustCompile("[ ]+")
 parsed := re.Split(expression, -1)

 if len(parsed) < 5 {

  return []string{}, errors.New("expression has only " + strconv.Itoa(len(parsed)) + " part(s), at least 5 parts are required")

 } else if len(parsed) == 5 {

  // 5 part cron so shift array past seconds element
  var newParsed []string
  newParsed = append(newParsed, "")
  newParsed = append(newParsed, parsed...)
  newParsed = append(newParsed, "")

  parsed = newParsed

 } else if len(parsed) == 6 {
  /* We will detect if this 6 part expression has a year specified and if so we will shift the parts and treat the
    		 first part as a minute part rather than a second part.

  		Ways we detect:

  	   	  1. Last part is a literal year (i.e. 2020)
     		  2. 3rd or 5th part is specified as "?" (DOM or DOW)
  */

  re2 := regexp.MustCompile("\\d{4}$")
  isYearWithNoSecondsPart := re2.Match([]byte(parsed[5])) || parsed[4] == "?" || parsed[2] == "?"

  if isYearWithNoSecondsPart {
   var newParsed []string
   newParsed = append(newParsed, "")
   newParsed = append(newParsed, parsed...)

   parsed = newParsed

  } else {
   // seconds provided
   //parsed.push("")
   parsed = append(parsed, "")
  }
 } else if len(parsed) > 7 {

  return []string{}, errors.New(`expression has `+strconv.Itoa(len(parsed))+` parts; too many`)

 }

 return parsed, nil
}


func (parser *CronParser) normalize(expressionParts []string) error {

 // Convert ? to * for DOM and DOW
 expressionParts[3] = strings.Replace(expressionParts[3], "?", "*", -1)
 expressionParts[5] = strings.Replace(expressionParts[5], "?", "*", -1)

 // Convert ? to * for Hour. ? isn't valid for hour position but we can work around it.
 expressionParts[2] = strings.Replace(expressionParts[2], "?", "*", -1)

 // Convert 0/, 1/ to */
 //if (expressionParts[0].indexOf("0/") == 0) {
 if strings.Index(expressionParts[0],"0/") == 0 {
  // Seconds
  expressionParts[0] = strings.Replace(expressionParts[0],"0/", "*/", -1)
 }

 if strings.Index(expressionParts[1],"0/") == 0 {
  // Minutes
  expressionParts[1] = strings.Replace(expressionParts[1],"0/", "*/", -1)
 }

 if strings.Index(expressionParts[2],"0/") == 0 {
  // Hours
  expressionParts[2] = strings.Replace(expressionParts[2],"0/", "*/", -1)
 }

 if strings.Index(expressionParts[3],"1/") == 0 {
  // DOM
  expressionParts[3] = strings.Replace(expressionParts[3],"1/", "*/", -1)
 }

 if strings.Index(expressionParts[4],"1/") == 0 {
  // Month
  expressionParts[4] = strings.Replace(expressionParts[4],"1/", "*/", -1)
 }

 if strings.Index(expressionParts[6],"1/") == 0 {
  // Years
  expressionParts[6] = strings.Replace(expressionParts[6],"1/", "*/", -1)
 }

 // Adjust DOW based on dayOfWeekStartIndexZero option
 // Normalized DOW: 0=Sunday/6=Saturday

 re2 := regexp.MustCompile(`(^\d)|([^#/\s]\d)`)
 // skip anything preceeded by # or /
 t := re2.FindAllString(expressionParts[5], -1)

 index :=0
 for _,_ = range t {

    re3 := regexp.MustCompile(`\D`)
    dowDigits := string(re3.ReplaceAll([]byte(t[index]), []byte("")))
    dowDigitsAdjusted := dowDigits

    //dayOfWeekStartIndexZero := true
    if parser.dayOfWeekStartIndexZero {
     // "7" also means Sunday so we will convert to "0" to normalize it
     if dowDigits == "7" {
      dowDigitsAdjusted = "0"
     }
    } else {
     // If dayOfWeekStartIndexZero==false, Sunday is specified as 1 and Saturday is specified as 7.
     // To normalize, we will shift the  DOW number down so that 1 becomes 0, 2 becomes 1, and so on.
     intToString, _ := strconv.Atoi(dowDigits)
     dowDigitsAdjusted = strconv.Itoa(intToString - 1)
    }

    expressionParts[5] = strings.Replace(expressionParts[5], dowDigits, dowDigitsAdjusted, 1)

    index++
 }


 // Convert DOW 'L' to '6' (Saturday)
 if expressionParts[5] == "L" {
  expressionParts[5] = "6"
 }

 // Convert DOM '?' to '*'
 if expressionParts[3] == "?" {
  expressionParts[3] = "*"
 }

 if strings.Index(expressionParts[3],"W") > -1 && (strings.Index(expressionParts[3],",") > -1 || strings.Index(expressionParts[3],"-") > -1) {
  return errors.New("the 'W' character can be specified only when the day-of-month is a single day, not a range or list of days")
 }

 // Convert DOW SUN-SAT format to 0-6 format
 days := map[string]int{
  "SUN": 0,
  "MON": 1,
  "TUE": 2,
  "WED": 3,
  "THU": 4,
  "FRI": 5,
  "SAT": 6,
 }

 for day := range days {
  expressionParts[5] = strings.ReplaceAll(expressionParts[5], day, strconv.Itoa(days[day]))
 }

 months := map[string]int{
  "JAN": 1,
  "FEB": 2,
  "MAR": 3,
  "APR": 4,
  "MAY": 5,
  "JUN": 6,
  "JUL": 7,
  "AUG": 8,
  "SEP": 9,
  "OCT": 10,
  "NOV": 11,
  "DEC": 12,
 }


 for month := range months {
  re := regexp.MustCompile(`(?i)`+month)
  expressionParts[4] = re.ReplaceAllString(expressionParts[4], strconv.Itoa(months[month]))
 }

 // Convert 0 second to (empty)
 if expressionParts[0] == "0" {
  expressionParts[0] = ""
 }

 // If time increment or * (every) is specified for seconds or minutes and hours part is single item, make it a "self-range" so
 // the expression can be interpreted as an increment / range.  This will allow us to easily interpret an hour part as 'between' a second or minute duration.
 //     For example:
 //     0-20/3 9 * * * => 0-20/3 9-9 * * * (9 => 9-9) => Every 3 minutes, minutes 0 through 20 past the hour, between 09:00 AM and 09:59 AM
 //     */5 3 * * * => */5 3-3 * * * (3 => 3-3) => Every 5 minutes, between 03:00 AM and 03:59 AM

 re3 := regexp.MustCompile(`\*|\-|\,|\/`)
 re4 := regexp.MustCompile(`\*|\/`)

 if !re3.MatchString(expressionParts[2]) && (re4.MatchString(expressionParts[1]) || re4.MatchString(expressionParts[0])) {
  expressionParts[2] += "-"+ expressionParts[2]
 }

 // Loop through all parts and apply global normalization
 for i := 0; i < len(expressionParts); i++ {
  // ignore empty characters around comma
  // if nothing left, set it to *
  if strings.Index(expressionParts[i],",") != -1 {
   tmp1 := strings.Split(expressionParts[i], ",")
   var tmp2 []string
   for i :=0; i < len(tmp1); i++ {
    if tmp1[i] !="" {
     tmp2 = append(tmp2, tmp1[i])
     break
    }
   }

   if len(tmp2) == 0 {
    expressionParts[i] = "*"
   } else {
    expressionParts[i] = strings.Join(tmp1, ",")
   }

  }

  // convert all '*/1' to '*'
  if expressionParts[i] == "*/1" {
   expressionParts[i] = "*"
  }

  /* Convert Month,DOW,Year step values with a starting value (i.e. not '*') to range expressions.
     This allows us to reuse the range expression handling for step values.

  	 For example:
  	 - month part '3/2' will be converted to '3-12/2' (every 2 months between March and December)
  	 - DOW part '3/2' will be converted to '3-6/2' (every 2 days between Tuesday and Saturday)
  */

  re5 := regexp.MustCompile(`^\*|\-|\,`)
  //if strings.Index(expressionParts[i],"/") > -1 && !/^\*|\-|\,/.test(expressionParts[i])) {
  if strings.Index(expressionParts[i],"/") > -1 && !re5.MatchString(expressionParts[i]) {
   stepRangeThrough := ""
   switch i {
   case 4:
    stepRangeThrough = "12"
    break
   case 5:
    stepRangeThrough = "6"
    break
   case 6:
    stepRangeThrough = "9999"
    break
   default:
    stepRangeThrough = ""
    break
   }

   if stepRangeThrough != "" {
    parts := strings.Split(expressionParts[i],"/")
    expressionParts[i] = parts[0]+"-"+stepRangeThrough +"/"+parts[1]
   }

  }

 }

 return nil
}


func (parser *CronParser) validate(parsed []string) {
   parser.assertNoInvalidCharacters("DOW", parsed[5])
   parser.assertNoInvalidCharacters("DOM", parsed[3])
   parser.validateRange(parsed)
}

func (parser *CronParser)  validateRange(parsed []string) error {

 var rv RangeValidator

 err1 := rv.secondRange(parsed[0])
 if err1 != nil {
  return err1
 }

 err2 := rv.minuteRange(parsed[1])
 if err2 != nil {
  return err2
 }

 err3 := rv.hourRange(parsed[2])
 if err3 != nil {
  return err3
 }

 err4 := rv.dayOfMonthRange(parsed[3])
 if err4 != nil {
  return err4
 }

 err5 := rv.monthRange(parsed[4])
 if err5 != nil {
  return err5
 }

 err6 := rv.dayOfWeekRange(parsed[5])
 if err6 != nil {
  return err6
 }

 return nil
}


func (parser *CronParser) assertNoInvalidCharacters(partDescription string, expression string) error {

     // No characters other than 'L' or 'W' should remain after normalization
     re := regexp.MustCompile("[A-KM-VX-Z]+/gi")
     invalidChars := re.Match([]byte(expression))

     if invalidChars {
         return errors.New(partDescription + " part contains invalid values: " + string(re.Find([]byte(expression))))
     }

     return nil

}