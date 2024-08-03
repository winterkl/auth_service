package app_errors

type TokenNotFound struct{}

func (e *TokenNotFound) Error() string {
	return "Токен не найден"
}
