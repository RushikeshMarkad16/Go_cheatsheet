package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Router init for rest APIs
	r := mux.NewRouter()

	// url/endpoints
	// all request land here
	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/getmethodurl", GetMethodFunc).Methods("GET")
	r.HandleFunc("/postmethodurl", PostMethodFunc).Methods("POST")

	// bind mux instace to http/net pkg
	http.Handle("/", r)

	fmt.Println(" Server running .... on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n ### homepage function called ####\n")
	fmt.Fprintf(w, "This is REST API") // display on web page
}

func GetMethodFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n ### GetMethodFunc called ####\n")
	fmt.Fprintf(w, "This is Get Method") // display on web page
}

func PostMethodFunc(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	req := map[string]interface{}{}
	err := decoder.Decode(&req)
	if err != nil {
		return
	}

	fmt.Printf("\n ### Post method called #### \n Recevied Params : %+v", req) // display on server logs/backend console

	msg := fmt.Sprintf("Hello %s!, This is Post Method.", req["name"])
	fmt.Fprintf(w, msg) // display on web page
}
