package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var servername string

type info struct {
	ImageVersion string    `json:"image_version"`
	Message      string    `json:"message"`
	Server       string    `json:"server"`
	DateTime     time.Time `json:"datetime"`
}

func main() {
	var err error
	servername, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods(http.MethodGet)
	fmt.Println("Done, listen port: 8080")
	http.ListenAndServe(":8080", r)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	info := info{
		ImageVersion: "v2.0.0",
		Message:      "Hello World!",
		Server:       servername,
		DateTime:     time.Now(),
	}
	buffer, err := json.Marshal(info)
	if err != nil {
		fmt.Print(err.Error())
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer)
	fmt.Println(string(buffer))
}
