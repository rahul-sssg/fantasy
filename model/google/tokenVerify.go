package usq_google

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/idtoken"
)

var token string          // this comes from your web or mobile app maybe
const googleClientId = "" // from credentials in the Google dev console

func ValidateToken(googleSignInToken string) {
	token = googleSignInToken

	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println("rr: ", err)
		// handle error, stop execution
	}

	payload, err := tokenValidator.Validate(context.Background(), token, googleClientId)
	if err != nil {
		fmt.Println("rr2: ", err)
		// handle error, stop execution
	}
	fmt.Println("payload: ", payload.Audience)
	fmt.Println("payload: ", payload.Expires)
	fmt.Println("payload: ", payload.IssuedAt)
	fmt.Println("payload: ", payload.Issuer)
	fmt.Println("payload: ", payload.Subject)
	fmt.Println("payload: ", payload.Claims)

	email := payload.Claims["email"]
	name := payload.Claims["name"]
	is_email_verified := payload.Claims["email_verified"]
	expires_at := payload.Claims["exp"]
	iat_at := payload.Claims["iat"]

	fmt.Println("rr expire: ", time.Now(), "  ", expires_at, " ", iat_at)

	family_name := payload.Claims["family_name"]
	given_name := payload.Claims["given_name"]
	url := payload.Claims["picture"]

	fmt.Println(email, ":= ", name, " ", family_name, " ", given_name, " ", url)
	fmt.Println("is_email_verified", ":=", is_email_verified)
}
