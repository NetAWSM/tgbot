package e

import "fmt"

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) error { //Вызвается функция и в случае если ошибка не нулевая, то вызывается Wrap
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}
