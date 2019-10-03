package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func mainhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main %s!", r.URL.Path[1:])
}

func sleepyhandler(w http.ResponseWriter, r *http.Request) {

	url := r.URL

	fmt.Println("Params were", url.Query())

	sleepParam := url.Query().Get("t")

	if sleepParam == "" {
		return
	}

	isleepParam, err := strconv.ParseFloat(sleepParam, 64)

	if err != nil {
		fmt.Printf("error %s !", err)
		fmt.Fprintf(w, "error %s !", err)
	}

	time.Sleep(time.Duration(isleepParam) * time.Second)

	fmt.Fprintf(w, "sleepy sleepy sleep %s !", sleepParam)
}

func main() {

	http.HandleFunc("/sleepy", sleepyhandler)
	http.HandleFunc("/", mainhandler)

	fmt.Println(http.ListenAndServe(":8080", nil))
}
