package main

import (
	"errors"
	"strconv"
	"strings"
)

func assert(value bool, message string) error {
	if !value {
		return errors.New(message)
	}

	return nil
}

type RangeValidator struct {
	test bool
}

func (rv *RangeValidator) secondRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		second, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(second >= 0 && second <= 59, "seconds part must be >= 0 and <= 59")
		}
	}

	return nil
}

func (rv *RangeValidator) minuteRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		minute, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(minute >= 0 && minute <= 59, "minutes part must be >= 0 and <= 59")
		}
	}

	return nil
}

func (rv *RangeValidator) hourRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		hour, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(hour >= 0 && hour <= 59, "hours part must be >= 0 and <= 23")
		}
	}

	return nil
}

func (rv *RangeValidator) dayOfMonthRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		dayOfMonth, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(dayOfMonth >= 0 && dayOfMonth <= 31, "DOM part must be >= 0 and <= 31")
		}
	}

	return nil
}

func (rv *RangeValidator) monthRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		month, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(month >= 0 && month <= 31, "month part must be >= 0 and <= 12")
		}
	}

	return nil
}

func (rv *RangeValidator) dayOfWeekRange(parse string) error {
	parsed := strings.Split(parse,",")

	for i := 0; i < len(parsed); i++ {
		month, err := strconv.Atoi(parsed[i])
		if err != nil {
			return assert(month >= 0 && month <= 6, "DOW part must be >= 0 and <= 6")
		}
	}

	return nil
}
