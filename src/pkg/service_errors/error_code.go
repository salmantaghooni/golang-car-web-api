package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"

	//OTP
	OTPExists   = "otp_exists"
	OTPUsed     = "otp_used"
	OTPNotValid = "otp_not_valid"

	// User
	EmailExists      = "Email exists"
	UsernameExists   = "Username exists"
	PermissionDenied = "Permission denied"

	// DB
	RecordNotFound = "record not found"
)
