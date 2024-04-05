package constant

import "errors"

var (
	ErrAccountExist           = errors.New("account-exist")
	ErrAccountIDExist         = errors.New("account-id-exist")
	ErrAccountNoIcon          = errors.New("account-no-icon")
	ErrAccountTokenExpired    = errors.New("account-token-expired")
	ErrLoginPassWrong         = errors.New("login-pass-wrong")
	ErrLoginEmailNotExist     = errors.New("login-email-not-exist")
	ErrBillKindNotExist       = errors.New("bill-kind-not-exist")
	ErrBillKindFatherNotExist = errors.New("bill-kind-father-not-exist")
	ErrRequestBodyNotValid    = errors.New("request-body-not-valid")
)
