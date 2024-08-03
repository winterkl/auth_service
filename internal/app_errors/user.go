package app_errors

import "fmt"

type UserAlreadyExists struct {
	Login string
}

func (e *UserAlreadyExists) Error() string {
	return fmt.Sprintf("Пользователь [LOGIN=%v] уже существует", e.Login)
}

type UserNotFound struct{}

func (e *UserNotFound) Error() string {
	return "Пользователь не найден"
}

type IncorrectAuthData struct{}

func (e *IncorrectAuthData) Error() string {
	return "Неверный логин или пароль"
}
