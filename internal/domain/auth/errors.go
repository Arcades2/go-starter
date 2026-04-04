package auth

import "app/internal/domain/errors"

var registry = errors.NewRegistry()

var (
	ErrInvalidCredentials    = registry.Register("INVALID_CREDENTIALS", "invalid credentials")
	ErrFailedToGenerateToken = registry.Register("FAILED_GENERATE_TOKEN", "failed to generate token")
	ErrRegisterInvalidInput  = registry.Register("REGISTER_INVALID_INPUT", "invalid input")
	ErrLoginInvalidInput     = registry.Register("LOGIN_INVALID_INPUT", "invalid input")
	ErrHashingPassword       = registry.Register("HASHING_PASSWORD_FAILED", "failed to hash password")
)

func AllAuthErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
