package pmapi

import "errors"

var (
	ErrNoConnection       = errors.New("no internet connection")
	ErrUnauthorized       = errors.New("API client is unauthorized")
	ErrUpgradeApplication = errors.New("application upgrade required")

	ErrBad2FACode         = errors.New("incorrect 2FA code")
	ErrBad2FACodeTryAgain = errors.New("incorrect 2FA code: please try again")

	ErrPaidPlanRequired = errors.New("paid subscription plan is required")
	ErrPasswordWrong    = errors.New("wrong password")
)

// ErrUnprocessableEntity ...
type ErrUnprocessableEntity struct {
	OriginalError error
}

func IsUnprocessableEntity(err error) bool {
	_, ok := err.(ErrUnprocessableEntity)
	return ok
}

func (err ErrUnprocessableEntity) Error() string {
	return err.OriginalError.Error()
}

// ErrBadRequest ...
type ErrBadRequest struct {
	OriginalError error
}

func IsBadRequest(err error) bool {
	_, ok := err.(ErrBadRequest)
	return ok
}

func (err ErrBadRequest) Error() string {
	return err.OriginalError.Error()
}

// ErrAuthFailed ...
type ErrAuthFailed struct {
	OriginalError error
}

func IsFailedAuth(err error) bool {
	_, ok := err.(ErrAuthFailed)
	return ok
}

func (err ErrAuthFailed) Error() string {
	return err.OriginalError.Error()
}

// ErrUnlockFailed ...
type ErrUnlockFailed struct {
	OriginalError error
}

func IsFailedUnlock(err error) bool {
	_, ok := err.(ErrUnlockFailed)
	return ok
}

func (err ErrUnlockFailed) Error() string {
	return err.OriginalError.Error()
}
