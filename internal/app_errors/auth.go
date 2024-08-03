package app_errors

type InvalidValidateToken struct{}

func (a *InvalidValidateToken) Error() string {
	return "Невалидный токен доступа"
}

type InvalidAccessToken struct{}

func (a *InvalidAccessToken) Error() string {
	return "Недопустимый токен доступа"
}
