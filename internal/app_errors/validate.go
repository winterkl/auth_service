package app_errors

import "fmt"

type IsRequired struct {
	Field string
}

func (e *IsRequired) Error() string {
	return fmt.Sprintf("Поле %s обязательно для заполнения", e.Field)
}
