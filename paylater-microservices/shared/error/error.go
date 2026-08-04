package errorx

// Common error message constants (stable client-facing text).
const (
	Unauthorized         = "Unauthorized"
	Forbidden            = "Access Denied"
	NotFound             = "Not Found"
	BadRequest           = "Bad Request"
	Internal             = "Internal Server Error"
	ServiceUnavailable   = "Service Unavailable"
	MissingAuthHeader    = "Authorization header missing"
	InvalidToken         = "Invalid Token"
	UnauthorizedInternal = "Unauthorized internal request"
)
