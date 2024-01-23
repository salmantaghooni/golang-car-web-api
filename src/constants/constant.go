package constants

// User
const (
	AdminRoleName   = "admin"
	DefaultRoleName = "default"
	DefaultUserName = "admin"
	DefaultPassword = "password"
	DefaultEmail    = "salman.taghooni@gmail.com"
)

// Redis
const RedisOTPDefaultKey = "otp_service"

// Claims
const (
	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserId"
	FirstNameKey           string = "FirstName"
	LastNameKey            string = "LastName"
	UsernameKey            string = "Username"
	EmailKey               string = "Email"
	MobileNumberKey        string = "MobileNumber"
	RolesKey               string = "Roles"
	ExpireTimeKey          string = "Exp"
)
