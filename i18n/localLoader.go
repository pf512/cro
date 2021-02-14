package i18n

type LocaleLoader interface{
	Load(availableLocales map[string]Locale)
}
