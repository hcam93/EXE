package runners

func CreateRunner(lang string) (Runner, error) {
	runner, ok := supportedLanguages[lang]
	if !ok {
		return nil, &UnsupportedLanguageError{lang: lang}
	}
	return runner, nil
}
