package main

import "fmt"

type MissingError struct{}

func (m *MissingError) Error() string {
	return "missing error"
}

type WrongError struct{}

func (w *WrongError) Error() string {
	return "wrong error"
}

func main() {
	Do("")
	Do("hello")
	Do("hello world")
}

func Do(str string) {
	switch err := Print(str); err.(type) {
	case *MissingError:
		// обработка ошибки
		fmt.Println(err)
	case *WrongError:
		// обработка ошибки
		fmt.Println(err)
	default:
		fmt.Println("!!!")
	}
}

func Print(str string) error {
	strLen := len(str)
	if strLen == 0 {
		return new(MissingError)
	}
	if strLen < 10 {
		return new(WrongError)
	}
	fmt.Print(str)
	return nil
}
