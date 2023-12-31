package ecode

// MovePermanently - moved permanently.
func MovePermanently(url string) error {
	return Newf(301, "moved permanently", "location: %s", url)
}

// Found - found.
func Found(url string) error {
	return Newf(302, "found", "location: %s", url)
}

// BadRequest - bad request error.
func BadRequest(reason, message string) error {
	return Newf(400, reason, message)
}

// Unauthorized - unauthorized error.
func Unauthorized(reason, message string) error {
	return Newf(401, reason, message)
}

// Forbidden - forbidden error.
func Forbidden(reason, message string) error {
	return Newf(403, reason, message)
}

// NotFound - not found error.
func NotFound(reason, message string) error {
	return Newf(404, reason, message)
}

// MethodNotAllowed - method not allowed error.
func MethodNotAllowed(reason, message string) error {
	return Newf(405, reason, message)
}

// RequestTimeout - request timeout error.
func RequestTimeout(reason, message string) error {
	return Newf(408, reason, message)
}

// Conflict - conflict error.
func Conflict(reason, message string) error {
	return Newf(409, reason, message)
}

// PreconditionFailed - precondition failed error.
func PreconditionFailed(reason, message string) error {
	return Newf(412, reason, message)
}

// RequestEntityTooLarge - request entity too large error.
func RequestEntityTooLarge(reason, message string) error {
	return Newf(413, reason, message)
}

// RateLimitExceeded - rate limit exceeded error.
func RateLimitExceeded(reason, message string) error {
	return Newf(429, reason, message)
}

// InternalServer - internal server error.
func InternalServer(reason, message string) error {
	return Newf(500, reason, message)
}

// NotImplemented - not implemented error.
func NotImplemented(reason, message string) error {
	return Newf(501, reason, message)
}

// Unavailable - unavailable error.
func Unavailable(reason, message string) error {
	return Newf(503, reason, message)
}

// GatewayTimeout - gateway timeout error.
func GatewayTimeout(reason, message string) error {
	return Newf(504, reason, message)
}
