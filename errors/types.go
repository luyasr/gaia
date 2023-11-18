package errors

func BadRequest(reason, message string) *Error {
	return New(400, reason, message)
}

func Unauthorized(reason, message string) *Error {
	return New(401, reason, message)
}

func Forbidden(reason, message string) *Error {
	return New(403, reason, message)
}

func NotFound(reason, message string) *Error {
	return New(404, reason, message)
}

func MethodNotAllowed(reason, message string) *Error {
	return New(405, reason, message)
}

func TooManyRequests(reason, message string) *Error {
	return New(429, reason, message)
}

func Internal(reason, message string) *Error {
	return New(500, reason, message)
}

func NotImplemented(reason, message string) *Error {
	return New(501, reason, message)
}

func ServiceUnavailable(reason, message string) *Error {
	return New(503, reason, message)
}

func GatewayTimeout(reason, message string) *Error {
	return New(504, reason, message)
}
