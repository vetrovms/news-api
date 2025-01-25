package errors

const (
	ResourceNotFound    = "resource not found"
	ServiceNotAvailable = "service not available"
	WrongFileFormat     = "wrong file format, allowed: jpeg, png"
	WrongJWT            = "missing or malformed JWT"
	ExpiredJWT          = "Invalid or expired JWT"
)
