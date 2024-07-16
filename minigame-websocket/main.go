package main

import (
	"encoding/json"
	"fmt"
	"minigame/infrastructure/db"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	godotenv.Load(".env")

	http.HandleFunc("/", Index)
	http.HandleFunc("/api", IndexApi)
	http.ListenAndServe(":8080", nil)
}

func Index(res http.ResponseWriter, req *http.Request) {
	// res.Write([]byte("Hello, World!"))
	db.Connect()
	templates.ExecuteTemplate(res, "Index", nil)
}

func IndexApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(User{
		Name: "John Doe 2",
		Age:  25,
	})

	fmt.Println("Before async test")
	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(5)
	go AsyncTest(&wg)
	go AsyncTest(&wg)
	go AsyncTest(&wg)
	go AsyncTest(&wg)
	go AsyncTest(&wg)
	wg.Wait()
	fmt.Println("After async test")
}

type User struct {
	Name string
	Age  int
}

func AsyncTest(wg *sync.WaitGroup) {
	defer DoneAsyncTest(wg)

	fmt.Println("Start async test")
	time.Sleep(2 * time.Second)
}

func DoneAsyncTest(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("Done async test")
}
