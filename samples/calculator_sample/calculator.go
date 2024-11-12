package calculator_sample

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	## Задача № 1
	Написать API для указанных маршрутов(endpoints)
	"/info"   // Информация об API
	"/first"  // Случайное число
	"/second" // Случайное число
	"/add"    // Сумма двух случайных чисел
	"/sub"    // Разность
	"/mul"    // Произведение
	"/div"    // Деление

	*результат вернуть в виде JSON

	"math/rand"
	number := rand.Intn(100)

	// Queries
	GET http://127.0.0.1:1234/first

	GET http://127.0.0.1:1234/second

	GET http://127.0.0.1:1234/add
	GET http://127.0.0.1:1234/sub
	GET http://127.0.0.1:1234/mul
	GET http://127.0.0.1:1234/div
	GET http://127.0.0.1:1234/info
*/

var (
	// Порт запуска приложения
	port       string = `8080`
	calculator CalcData
)

type CalcData struct {
	FirstNumber  float64 `json:"first_number"`
	SecondNumber float64 `json:"second_number"`
	Result       float64 `json:"result"`
	ErrorMessage string  `json:"error_message"`
}

type Request struct {
	CalcData   CalcData `json:"calc_data"`
	LogMessage string   `json:"log_message"`
}

func CalculatorSample() {

	log.Println("Trying to start calc REST API!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()

	//1. Если на вход пришел запрос /GetInfo
	router.HandleFunc("/info", GetInfo).Methods("GET")
	router.HandleFunc("/first", GetFirst).Methods("GET")
	router.HandleFunc("/second", GetSecond).Methods("GET")

	router.HandleFunc("/add", Add).Methods("GET")
	router.HandleFunc("/sub", Sub).Methods("GET")
	router.HandleFunc("/mul", Mul).Methods("GET")
	router.HandleFunc("/div", Div).Methods("GET")

	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetFirst(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get first number API")
	writer.WriteHeader(200) // StatusCode для запроса

	{ //исп блок
		randInt := rand.Intn(100)
		calculator.FirstNumber = float64(randInt)
	}

	req := Request{
		CalcData:   calculator,
		LogMessage: "Заполнено случайное первое число",
	}

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func GetSecond(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get second number API")
	writer.WriteHeader(200) // StatusCode для запроса

	{ //исп блок
		randInt := rand.Intn(100)
		calculator.SecondNumber = float64(randInt)
	}

	req := Request{
		CalcData:   calculator,
		LogMessage: "Заполнено случайное второе число",
	}

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func Div(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Div operations API")
	writer.WriteHeader(200) // StatusCode для запроса

	req := Request{
		LogMessage: "Операция деления",
	}
	{ //исп блок
		if calculator.SecondNumber == 0 {
			req.CalcData.ErrorMessage = "Нельзя делить на ноль"
		} else {
			calculator.Result = float64(calculator.FirstNumber) / float64(calculator.SecondNumber)
		}
	}

	req.CalcData = calculator

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func Mul(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Mult operations API")
	writer.WriteHeader(200) // StatusCode для запроса

	{ //исп блок
		calculator.Result = float64(calculator.FirstNumber) * float64(calculator.SecondNumber)
	}

	req := Request{
		CalcData:   calculator,
		LogMessage: "Операция умножения",
	}

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func Add(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Add operations API")
	writer.WriteHeader(200) // StatusCode для запроса

	{ //исп блок
		calculator.Result = float64(calculator.FirstNumber + calculator.SecondNumber)
	}

	req := Request{
		CalcData:   calculator,
		LogMessage: "Операция сложения",
	}

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func Sub(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Sub operations API")
	writer.WriteHeader(200) // StatusCode для запроса

	{ //исп блок
		calculator.Result = float64(calculator.FirstNumber - calculator.SecondNumber)
	}

	req := Request{
		CalcData:   calculator,
		LogMessage: "Операция вычитания",
	}

	json.NewEncoder(writer).Encode(req) // Сериализация + запись в writer
}

func GetInfo(writer http.ResponseWriter, request *http.Request) {
	// Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get info about API")
	writer.WriteHeader(200) // StatusCode для запроса
	message :=
		`Данный сервер умееет выполнять следующие операции:
"/info"   Информация об API
"/first"  Случайное число
"/second" Случайное число
"/add"    Сумма двух случайных чисел
"/sub"    Разность
"/mul"    Произведение
"/div"    Деление
`
	writer.Write([]byte(message))
	// json.NewEncoder(writer).Encode(message) // Сериализация + запись в writer
}
