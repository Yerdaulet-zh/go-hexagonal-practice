package main

import (
	"encoding/json"
	"fmt"
)

type GeoLocation struct {
	City string `json:"city5"` // Struct tags define the JSON keys
}

type UserSession struct {
	Location *GeoLocation
}

func main() {
	// --- 1. TRANSFORMATION: DOMAIN TO DATABASE (MARSHAL) ---
	// Imagine your Service layer gives you this:
	domainLocation := GeoLocation{City: "Almaty"}

	// We convert the struct into a slice of bytes (JSON)
	// This is what gorm.io/datatypes.JSON stores in Postgres
	jsonData, _ := json.Marshal(domainLocation)

	fmt.Printf("Database View (JSON Bytes): %s\n", string(jsonData))

	// --- 2. TRANSFORMATION: DATABASE TO DOMAIN (UNMARSHAL) ---
	// Now imagine GORM gives you those bytes back:
	rawFromDB := []byte(`{"city5":"Almaty"}`)

	// We need a place to "pour" the data into
	var decodedLocation GeoLocation

	// Unmarshal takes the bytes and fills the struct fields
	err := json.Unmarshal(rawFromDB, &decodedLocation)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Now we can put it back into our Domain Session
	session := UserSession{
		Location: &decodedLocation,
	}

	fmt.Printf("Domain View (Go Struct): %+v\n", session.Location.City)
}
