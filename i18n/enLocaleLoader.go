package i18n

type EnLocaleLoader struct {
}

func (enl EnLocaleLoader) Load(availableLocales map[string]Locale) {
	var languageImp en
	languageCode := "en"
	availableLocales[languageCode] = languageImp
}


