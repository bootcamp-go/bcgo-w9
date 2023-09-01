package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	js := `{"name": "John Doe", "age": 25}`

	var jsonMap map[string]any
	json.Unmarshal([]byte(js), &jsonMap)
	
	fmt.Println(jsonMap)
}