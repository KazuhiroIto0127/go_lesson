package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsonString = `
{
	"species": "ハト",
	"description": "岩にとまるのが好き",
	"dimensions": {
		"height": 24,
		"width": 10
	}
}
`

type Dimensions struct {
	Width int `json:"width"`
	Height int `json:"height"`
}

type Data struct {
	Species string `json:"species"`
	Description string `json:"description"`
	Dimensions Dimensions `json:"dimenstions"`
}

func main(){
	data := Data{
		Species: "つばめ",
		Description: "小さい鳥",
		Dimensions: Dimensions{
			Height: 10,
			Width: 20,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
	fmt.Println("JSON data:", string(jsonData))

	var data2 Data
	err = json.Unmarshal([]byte(jsonString), &data2)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	fmt.Println("Decode JSON data:", data2)
	fmt.Println(data2.Description)
}
