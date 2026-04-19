package main

import (
	"context"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
}

func query(username string) {
	var user User
	err := pool.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", username).Scan(&user.Username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}

	fmt.Println(user.Username)
}
