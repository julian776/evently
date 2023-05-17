package errors

func Check(err error, handler func(...any) any) any {
	if err != nil {
		return handler()
	}
	return func() {
		// Don't do anything
	}
}
