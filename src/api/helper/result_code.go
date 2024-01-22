package helper

const (
	Success         int = 0
	ValidationError int = 40001
	AuthError       int = 40101
	ForbiddenError  int = 40301
	NotFoundError   int = 40401
	LimiterError    int = 42901
	OtpLimiterError int = 42902
	CustomRecovery  int = 50001
	InternalError   int = 50002
)
