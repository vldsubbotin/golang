package main

import (
	"fmt"
)

func main() {

	customers := getCustomers()

	for _, customer := range customers {
		fmt.Println(customer)
	}

}

//func getData() (customers []string) {
//
//	customers = []string{"Vladimir Botin", "Vladimir Subbotin", "Vlad Subbotin"}
//
//	customers = append(customers, "Ben Spain")
//	customers = append(customers, "Aleem Janmohamed")
//	customers = append(customers, "Jamie le Notre")
//	customers = append(customers, "Victor Savkov")
//	customers = append(customers, "P The Admin")
//	customers = append(customers, "Adrian Oprea")
//	customers = append(customers, "Jonathan D")
//
//	for _, customer := range customers {
//		fmt.Println(customer)
//	}
//
//	return customers
//}
