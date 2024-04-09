package checkers

const MAX_CWS_LIMIT = 180

func IsWithinCwsLimit(phrase string) bool {
	return len(phrase) <= MAX_CWS_LIMIT
}
