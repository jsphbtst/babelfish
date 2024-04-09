package checkers

func IsSupportedLanguage(language string) bool {
	supported := map[string]bool{
		"english":    true,
		"tagalog":    true,
		"ilocano":    false, // TODO
		"bisaya":     false, // TODO
		"german":     true,
		"japanese":   true,
		"nihongo":    true,
		"french":     true,
		"spanish":    true,
		"castellano": true,
	}

	isSupported, ok := supported[language]
	if !ok {
		return false
	}

	return isSupported
}
