package common

func PrintRecoverFromPanic(isPanic ...*bool) {
	if err := recover(); err != nil {
		if len(isPanic) != 0 {
			*isPanic[0] = true
		}
	}
}
