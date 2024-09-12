package libs

func (e *ErrorMessage) Error() string {
	return e.Message
}
