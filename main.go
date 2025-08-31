package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/good-binary/utility/uuid"
)

func main() {
	// Generate a new UUID
	u := uuid.NewUUID()
	fmt.Println("Generated UUID:", u.String())

	// Validate the UUID string
	valid := uuid.Validate(u.String())
	fmt.Println("Is valid?", valid)

	// Parse from string
	parsed, err := uuid.Parse(u.String())
	if err != nil {
		log.Fatal("Parse error:", err)
	}
	fmt.Println("Parsed UUID:", parsed.String())

	// Compare UUIDs
	fmt.Println("Equal?", u.Equal(parsed))

	// Nil UUID
	nilUUID := uuid.Nil()
	fmt.Println("Nil UUID:", nilUUID.String())

	// Marshal to JSON
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Marshal error:", err)
	}
	fmt.Println("UUID as JSON:", string(jsonBytes))

	// Unmarshal from JSON
	var u2 uuid.UUID
	err = json.Unmarshal(jsonBytes, &u2)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}
	fmt.Println("Unmarshaled UUID:", u2.String())
	fmt.Println("Equal after unmarshal?", u.Equal(u2))
}
