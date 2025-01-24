package main

import "fmt"

type UserDetails struct {
	firstName string
	lastName  string
	age       int
}

type ContactInfo struct {
	phone string
	email string
}

type AddressInfo struct {
	street string
	city   string
}

type User struct {
	userDetails UserDetails
	contactInfo ContactInfo
	addressInfo AddressInfo
}

func main() {
	// Create a new user
	user := User{
		userDetails: UserDetails{
			firstName: "John",
			lastName:  "Doe",
			age:       30,
		},
		contactInfo: ContactInfo{
			phone: "123-456-7890",
			email: "john.doe@example.com",
		},
		addressInfo: AddressInfo{
			street: "123 Main St",
			city:   "Anytown",
		},
	}

	user2 := User{
		userDetails : UserDetails{
			firstName: "Rupesh",
			lastName:  "Bhosale",
			age:       25,
		},
		contactInfo: ContactInfo{
			phone: "8080186885",
			email: "rupesh.bhosale@example.com",
		},
		addressInfo: AddressInfo{
			street: "Near Post Office",
			city:   "Pune",
		},

	}
	// Print the user's details
	fmt.Println("User Details:", user)
	fmt.Println(user.addressInfo.city)
	fmt.Println("User 2 : ",user2)
}
