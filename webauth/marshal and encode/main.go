package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	/*p1 := person{First: "James"}
	p2 := person{First: "Don"}
	xp := []person{p1, p2}

	marshaled, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(marshaled))

	xp2 := []person{}
	err = json.Unmarshal(marshaled, &xp2)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(xp2)
	*/

	http.HandleFunc("/encode", encodeF)
	http.HandleFunc("/decode", decodeF)
	http.ListenAndServe(":8080", nil)
}
func encodeF(w http.ResponseWriter, r *http.Request) {
	p1 := person{First: "James"}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
func decodeF(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println(err)
	}
	log.Println(p1)
}
