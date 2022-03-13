package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func getCustomers() (customers []Customer) {

	vladimir := Customer{FirstName: "Vladimir", LastName: "Subbotin"}

	customers = append(customers, vladimir,
		Customer{FirstName: "Anya", LastName: "Kakulkina"},
		Customer{FirstName: "Mark", LastName: "Richardson"},
		Customer{FirstName: "David", LastName: "Blah"},
		Customer{FirstName: "Pol", LastName: "Bobotin"},
		Customer{FirstName: "Scott", LastName: "Merphy"},
		Customer{FirstName: "Daniel", LastName: "Sorokin"},
	)

	return customers

}
