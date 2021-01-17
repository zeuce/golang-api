package main

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
)
 
var port string = ":" + strconv.Itoa(GetConfig().Port)

//Check ... Check for errors
func Check(e error){
	if(e != nil){
		fmt.Println(e)
	}
}

func main() {
	if(GetConfig().SSL) {
		fmt.Printf("Listening on https://localhost%s\n", port)
		log.Fatal(http.ListenAndServeTLS(port, GetConfig().CertFile, GetConfig().KeyFile, GetRouter()))
	} else {
		fmt.Printf("Listening on http://localhost%s\n", port)
		log.Fatal(http.ListenAndServe(port, GetRouter()))
	}
}