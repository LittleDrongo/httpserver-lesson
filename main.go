package main

import (
	"fmt"

	"github.com/LittleDrongo/httpserver-lesson/samples"
)

func main() {

}

func marshalSamples() {

	err := samples.MarshallSample()
	if err != nil {
		fmt.Printf("Ошибка сериализации JSON: %v\n", err)
	} else {
		fmt.Println("1. Сериализация прошла успешно")
	}

	err = samples.UnmarshallSample()

	if err != nil {
		fmt.Printf("Ошибка де-сериализации JSON: %v\n", err)
	} else {
		fmt.Println("2. Де-сериализация прошла успешно")
	}
}
