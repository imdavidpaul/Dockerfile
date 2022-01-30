package main

import (
	"fmt"
)

func main() {

	customers := GetCustomers()

	for _, customer := range customers {
		//we can access the "customer" variable in this approach
		fmt.Println(customer)
	}
}
