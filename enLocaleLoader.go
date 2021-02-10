package main

import (
	"fmt"
	"reflect"
)

type enLocaleLoader struct {
	language string `default:"en"`
}

func (enl enLocaleLoader) load(availableLocales map[string]Locale) {
	var englishImp en
	availableLocales[default_tag(enLocaleLoader{})] = englishImp
}

func default_tag(l enLocaleLoader) string {

	// TypeOf returns type of
	// interface value passed to it
	typ := reflect.TypeOf(l)

	// checking if null string
	if l.language == "" {

		// returns the struct field
		// with the given parameter "language"
		f, _ := typ.FieldByName("language")

		// returns the value associated
		// with key in the tag string
		// and returns empty string if
		// no such key in tag
		l.language = f.Tag.Get("default")
	}

	return fmt.Sprintf("%s", l.language)
}

