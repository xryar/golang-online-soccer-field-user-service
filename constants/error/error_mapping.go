package error

func ErrMapping(err error) bool {
	allErrors := append(GeneralErrors[:], UserErrors[:]...)

	for _, item := range allErrors {
		if err.Error() == item.Error() {
			return true
		}
	}

	return false
}
