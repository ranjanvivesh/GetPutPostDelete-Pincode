package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

// Structure
type details struct{
	Pincode  string		`json:pincode`
	Ara	 string		`json:area`
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

}

func serveHome(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("<h1>WELCOME TO PINCODE API</h1>"))
}

func getAllPincodes(w http.ResponseWriter, r http.Request) {
	fmt.Println("Get All of the Pincodes")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Container)//Feeding an Empty data to the variable created above so that we can later feed the desired data

}

func getOnePincode(w http.ResponseWriter, r http.Request) {
	fmt.Println("Get One of the Pincodes")
	w.Header().Set("Content-type", "application/json")	

	
}