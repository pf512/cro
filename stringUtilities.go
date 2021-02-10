package main

import "strings"

type StringUtilities struct {
}

func GetStringUtilities() StringUtilities {
	var su StringUtilities
	return su
}

/**
 * Takes a string with '%s' placeholders and replaces them with provided values.
 * Works like sprintf in C or string.Format in C#.
 * @static
 * @param {string} template - The string template with enclosed %s replacements
 * @param {...string[]} values - The ordered replacement text
 * @returns {string}
 */
func (su *StringUtilities) format(template string, values ...string) string {

	for _, replacementString := range values {
		template = strings.Replace(template, "%s", replacementString,1)
	}

	return template
}

/**
 *
 * Given a string and an array of search strings, determines if the string
 * contains any value from the array.
 * @static
 * @param {string} text - The string to search
 * @param {string[]} searchStrings - The array of values to search for
 * @returns {boolean}
 */
func (su *StringUtilities) containsAny(text string, searchStrings []string) bool {

	for _, searchValue := range searchStrings {
		if strings.Contains(text, searchValue) {
			return true
		}
	}

	return false
}
