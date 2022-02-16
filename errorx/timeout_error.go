package errorx

type TimeoutError struct {
	error
}

func (e *TimeoutError) Timeout() bool {
	return true
}

func (e *TimeoutError) Temporary() bool {
	return true
}

func (e *TimeoutError) Error() string {
	return "context deadline exceeded (Client.Timeout exceeded while awaiting headers)"
}
