package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// exported var global
type User struct {
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
	Email       string `faker:"email"`
	Mobile      string `faker:"phone_number"`
	UserAddress Address
}

type Address struct {
	Pin     int
	Country string
}

func main() {

	var user User
	err := faker.FakeData(&user)

	if err == nil {
		fmt.Printf("%+v\n", user)
	}
}

// func main() {

//		fmt.Println("Hello")
//		f := faker.FakeData(&User{})
//		fmt.Println(f)
//	}
// type Gondoruwo struct {
// 	Name       string
// 	Locatadata int
// }

// type Sample struct {
// 	ID        int64     `faker:"customIdFaker"`
// 	Gondoruwo Gondoruwo `faker:"gondoruwo"`
// 	Danger    string    `faker:"danger"`
// }

// func CustomGenerator() {
// 	// explicit
// 	faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
// 		return int64(43), nil
// 	})

// 	// functional
// 	faker.AddProvider("danger", func() faker.TaggedFunction {
// 		return func(v reflect.Value) (interface{}, error) {
// 			return "danger-ranger", nil
// 		}
// 	}())

// 	faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
// 		obj := Gondoruwo{
// 			Name:       "Power",
// 			Locatadata: 324,
// 		}
// 		return obj, nil
// 	})

// }
