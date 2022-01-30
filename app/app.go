package main

import "fmt"

func main() {

	customer := getData(1)
	fmt.Println(customer)
}
func getData(customerID int) (customer string) {

	var firstName = "Arun"
	lastname := "Reddy"
	customer = firstName + " " + lastname
	return customer
}
