package i18n

type EsLocaleLoader struct {
}

func (esl EsLocaleLoader) Load(availableLocales map[string]Locale) {
	var languageImp es
	languageCode := "es"
	availableLocales[languageCode] = languageImp
}
