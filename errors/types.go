package errors

func BadRequest(reason, format string, args ...any) *Error {
	return New(400, reason, format, args...)
}

func Unauthorized(reason, format string, args ...any) *Error {
	return New(401, reason, format, args...)
}

func Forbidden(reason, format string, args ...any) *Error {
	return New(403, reason, format, args...)
}

func NotFound(reason, format string, args ...any) *Error {
	return New(404, reason, format, args...)
}

func MethodNotAllowed(reason, format string, args ...any) *Error {
	return New(405, reason, format, args...)
}

func TooManyRequests(reason, format string, args ...any) *Error {
	return New(429, reason, format, args...)
}

func Internal(reason, format string, args ...any) *Error {
	return New(500, reason, format, args...)
}

func NotImplemented(reason, format string, args ...any) *Error {
	return New(501, reason, format, args...)
}

func ServiceUnavailable(reason, format string, args ...any) *Error {
	return New(503, reason, format, args...)
}

func GatewayTimeout(reason, format string, args ...any) *Error {
	return New(504, reason, format, args...)
}
