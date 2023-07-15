package main

import (
	"encoding/json"
	"fmt"
	// "os"
	// "github.com/jcelliott/lumber"
)

func main() {
	fmt.Println("Blah!!!")
	dir := "./"

	db, err := New(dir, nil)

	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", 23, "123", "City Frog", Address{"Fity", "Ftate", "Funtry", "123456"}},
		{"John Weak", 24, "1234", "City Deer", Address{"Fitys", "Ftates", "Funtrys", "123456"}},
		{"John Leak", 25, "12345", "City Cock", Address{"Fityst", "Ftatest", "Funtryys", "123456"}},
		{"John Beak", 26, "123456", "City Viper", Address{"Fitysty", "Ftatests", "Funtries", "123456"}},
		{"John Sick", 27, "1234567", "City Crab", Address{"Fitystys", "Ftatestst", "Funtryies", "123456"}},
		{"John Tick", 28, "12345678", "City Lion", Address{"Fitystyst", "Ftateststs", "Funtryyyies", "123456"}},
	}
	for _, test := range employees {
		db.Write("users", test.Name, User{
			Name:    test.Name,
			Age:     test.Age,
			Contact: test.Contact,
			Company: test.Company,
			Address: test.Address,
		})
	}
	records, err := db.ReadAll("Users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)
	allusers := []User{}
	for _, user := range records {
		employeeFound := user{}
		if err := json.Unmarshal([]byte(user), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println(allusers)

	// to delete data form db
	// if err :=db.Delete("user","john")l err!=nil{
	// 	fmt.Println("Error", err)
	// }
	// if err := db.Delete("user","");err!=nil{
	// 	fmt.Println("Error", err)
	// }
}
