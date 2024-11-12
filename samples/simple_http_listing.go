package samples

import (
	"fmt"
	"log"
	"net/http"
)

// w - responeWritter (куда писать ответ)
// r - request (откуда брать запрос)
// Обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	const (
		message = `<h1> Hello! im new web-server!</h1>`
	)

	fmt.Fprintf(w, message)

}

// Функция, которая выбирает нужный обработчик, в зависимости от адреса, по которому пришел запрос.
func RequestHandler() {
	const (
		path = `/`
	)
	http.HandleFunc(path, GetGreet)              // Если придёт запрос по адресу "/", то вызывай функцию обработчика GetGreet.
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем веб сервер в ержиме "слушния" на порту 8080
}

func SimpleHttplListen() {

	// Создаём ресурс
	RequestHandler()
}
