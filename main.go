package main

import (
    "fmt"
    // "errors"
    "strconv"
    // "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

type calResource struct {
	num1  float64 `json:"num1"`
	num2  float64 `json:"num2"`
}

var x, y float64
var err error
var port int = 3000

// Add performs addition of two numbers
func Add(w http.ResponseWriter, r *http.Request) {
  num1 := r.FormValue("num1")
  num2 := r.FormValue("num2")
  x, err := strconv.ParseFloat(num1, 64)
  y, err := strconv.ParseFloat(num2, 64)
  if err != nil {
  		fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": \"%s\"}\n", "Error detected at input.")
      return
  }
  fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": %f }\n", x, y, x + y)
}

// Subtract performs subtraction of two numbers
func Subtract(w http.ResponseWriter, r *http.Request) {
  num1 := r.FormValue("num1")
  num2 := r.FormValue("num2")
  x, err := strconv.ParseFloat(num1, 64)
  y, err := strconv.ParseFloat(num2, 64)
  if err != nil {
  		fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": \"%s\"}\n", "Error detected at input.")
      return
  }
  fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": %f }\n", x, y, x - y)
}

// Multiply performs multiplication of two numbers
func Multiply(w http.ResponseWriter, r *http.Request) {
  num1 := r.FormValue("num1")
  num2 := r.FormValue("num2")
  x, err := strconv.ParseFloat(num1, 64)
  y, err := strconv.ParseFloat(num2, 64)
  if err != nil {
  		fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": \"%s\"}\n", "Error detected at input.")
      return
  }
  fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": %f }\n", x, y, x * y)
}

// Divide performs division of two numbers
func Divide(w http.ResponseWriter, r *http.Request) {
  num1 := r.FormValue("num1")
  num2 := r.FormValue("num2")
  x, err := strconv.ParseFloat(num1, 64)
  y, err := strconv.ParseFloat(num2, 64)
  if err != nil {
  		fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": \"%s\"}\n", "Error detected at input.")
      return
  }
  if y == 0 {
      // fmt.Println("Division by zero is not allowed. We return 0.")
      fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": \"%s\" }\n", x, y, "Can't divide by 0.")
      return
  }
  fmt.Fprintf(w, "{ \"num1\": %f, \"num2\": %f, \"answers\": %f }\n", x, y, x / y)
}

func main() {
  fmt.Println("Starting server on port", port)

	r := chi.NewRouter()

  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)

  r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
  	w.Write([]byte("pong\n"))
  })

  r.Get("/add", Add)
  r.Get("/sub", Subtract)
  r.Get("/mul", Multiply)
  r.Get("/div", Divide)

  lport := fmt.Sprintf(":%d", port)
  http.ListenAndServe(lport, r)

}
