package main

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientId := "[myClientID]"
	id_token := `[myIDToken]`

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	payload, err := tokenValidator.Validate(context.Background(), id_token, googleClientId)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
