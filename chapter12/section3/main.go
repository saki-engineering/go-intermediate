package main

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientID := "[myClientID]"
	idToken := `[myIDToken]`

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
