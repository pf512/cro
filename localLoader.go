package main

type LocaleLoader interface{
	load(availableLocales map[string]Locale)
}
