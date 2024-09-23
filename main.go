package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Structure
type details struct{
	Pincode  string		`json:pincode`
	Area	 string		`json:area`
	District string 	`json:district`
	State	 string		`json:state`
}

//Fake Database
var Container []details

//Middleware,helper

func (c *details) IsEmpty() bool {
	return c.Pincode == "" && c.Pincode == ""
}


func main() {
	fmt.Println("PINCODE API")
	r := mux.NewRouter()

	//seeding the details

	Container = append(Container, details{Pincode: "835219",Area: "Chakla",District: "Ranchi",State: "Jharkhand"})
	Container = append(Container, details{Pincode: "832001",Area: "Ranchi",District: "Ranchi",State: "Jharkhand"})

	//routing

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/pincodes",getAllPincodes).Methods("GET")
	r.HandleFunc("/pincodes/{pin}",getOnePincode).Methods("GET")
	r.HandleFunc("/pincodes",createOnePincode).Methods("POST")
	r.HandleFunc("/pincodes",updateOnePincode).Methods("PUT")
	r.HandleFunc("/pincodes/{pin}",deleteOnePincode).Methods("DELETE")

	//listening to port

	log.Fatal(http.ListenAndServe(":8000",r))
	
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO PINCODE API</h1>"))
}

func getAllPincodes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All of the Pincodes")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Container)//Feeding an Empty data to the variable created above so that we can later feed the desired data

}

func getOnePincode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One of the Pincodes")
	w.Header().Set("Content-type", "application/json")	

	params := mux.Vars(r) //getting the pin

	//loop and matching the pincode

	for _,content := range Container{
		if content.Pincode== params["pin"] {
			json.NewEncoder(w).Encode(Container)
			return
		}
	}
	json.NewEncoder(w).Encode("Please Enter Valid Pincode Details")
	return

}

func createOnePincode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One of the Pincodes")
	w.Header().Set("Content-type", "application/json")

	// If the body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Enter Valid Pincode Details")
	}

	// If the data send is and empty set i.e. {}

	var Container details
	_ = json.NewDecoder(r.Body).Decode(&Container)

	if Container.IsEmpty() {
		json.NewEncoder(w).Encode("Please Enter Valid Pincode Details")
		return
	}
}

func updateOnePincode(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Get One of the Pincodes")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)//get the picode from the request using r

	

	for index,content  := range Container{
		if content.Pincode == params["pin"] {
			Container = append(Container[:index], Container[index+1:]...)
			var content details
			_ = json.NewDecoder(r.Body).Decode(&content)
			content.Pincode = params["pin"]
			Container = append(Container, content)
			json.NewEncoder(w).Encode(content)
			return
		}
	}

}

func deleteOnePincode(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Get One of the Pincodes")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r) //getting the pin

	for index, content := range Container {
		if content.Pincode == params["id"] { //matching the id
			Container = append(Container[:index], Container[index+1:]...)
			break
		}
	}
}
