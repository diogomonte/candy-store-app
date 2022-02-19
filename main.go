package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s := Store{}
	s.LoadData()
	customers := s.GetSortedCustomers()

	result, err := json.MarshalIndent(&customers, "", " ")
	if err != nil {
		log.Fatalln("Error converting customers to json")
		return
	}

	fmt.Println(string(result))
}
