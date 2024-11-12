package main

import (
	"fmt"

	"github.com/LittleDrongo/httpserver-lesson/samples"
)

func main() {

	err := samples.MarshallSample()
	if err != nil {
		fmt.Printf("Ошибка сериализации JSON: %v\n", err)
	} else {
		fmt.Println("Сериализация прошла успешно")
	}
}
