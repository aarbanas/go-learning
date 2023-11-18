package structs

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	// this column has the struct tags
	UserID    string    `json:"userID"`
	IsActive  bool      `json:"isActive"`
	LastLogin time.Time `json:"lastLogin"`
	UserType  int       `json:"userType"`
}

func ExecuteJsonEncoding() {
	user := User{UserID: "123", IsActive: true, UserType: 2, LastLogin: time.Now()}

	jsonResponse, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonResponse))
}
