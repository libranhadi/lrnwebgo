package test

import (
	"testing"
	"fmt"
	"encoding/json"
	"os"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	ProfileUrl string `json:"profile_url"`
	Hobbies []string
	Addresses []Address
}

type Address struct {
	Street string
	Country string
	PostalCode string
}

func TestEncodeJson(t *testing.T)  {
	data := map[string]interface{}{
		"Firstname": "John",
		"Lastname":  "Kyu",
		"profile_url":  "https://example.jpg",
		"age": 2,
		"Hobbies" : []string{"Golf", "Football"},
		"Addresses" : []Address {
			{
				Street : "Nakama Street",
				Country : "Wano",
				PostalCode : "1071",
			},
			{
				Street : "Fishman Street",
				Country : "Sabaody",
				PostalCode : "487",
			},
		},
		
	}

	user, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(user)
}

func TestDecodeJson(t *testing.T)  {
	var jsonBlob = []byte(`[
			{
				"age":10,"Firstname":"A","Lastname":"Ji", "Hobbies": ["Golf","Football", "Coding"], 
				"Addresses": [
									{
										"Street": "Nakama Street",
										"Country": "Wano",
										"PostalCode": "1071"
									},
									{
										"Street": "Fishman Street",
										"Country": "Sabaody",
										"PostalCode": "487"
									}
							],
				"profile_url":  "https://example.jpg"
			},
			{"age":10,"Firstname":"D","Lastname":"Law", "Hobbies": ["Swim","Fishing"]}
		]`)

	var users []User
	err := json.Unmarshal(jsonBlob, &users)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range users {
		fmt.Println("Index : ", i , v.Firstname)
		fmt.Println("Index : ", i ,  v.Lastname)
		fmt.Println("Index : ", i, v.ProfileUrl)
		fmt.Println("Index : ", i , v.Age)
		for _, a := range v.Hobbies {
			fmt.Println("Index : ", i , "HOBBY : " , a)
		}
		for _, b := range v.Addresses {
			fmt.Println("Index : ", i , "Country : " , b.Country)
			fmt.Println("Index : ", i , "Street : " , b.Street)
			fmt.Println("Index : ", i , "PostalCode : " , b.PostalCode)
		}
	}
}

func TestEncodeJsonWithMap(t *testing.T)  {
	products := map[string]interface{}{
		"id" : 10,
		"name" : 'D',
		"price" : 109.00,
	}

	data, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(data)
	
}


func TestDecodeJsonWithMap(t *testing.T)  {
	jsonStringProduct := []byte(`{"id":10,"name":"D","price":"10.00"}`)
	var products map[string]interface{}

	err := json.Unmarshal(jsonStringProduct, &products);
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range products {
		fmt.Println(v);
	}
}


func TestDecoder(t *testing.T)  {
	reader, _ := os.Open("User.json");

	user := &User{}
	data := json.NewDecoder(reader);
	data.Decode(user);
	fmt.Println(user)
}


func TestEncoder(t *testing.T)  {
	writer, _ := os.Create("output.json");
	data := json.NewEncoder(writer);
	user := User{	
		Firstname: "John",
		Lastname:  "Kyu",
		ProfileUrl:  "https://example.jpg",
		Age: 2,
	}

	data.Encode(user);

	fmt.Println(user)
}
