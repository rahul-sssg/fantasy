package usq_google

import "fmt"

func MapCheck() {

	mapCheck := map[string]string{
		"at_hash":        "dDq05pUC7K8nQxFfDwzapg",
		"aud":            "35858082528-b3ekj6u835sd31q4almrtuslaulpgcrm.apps.googleusercontent.com", //The value of aud in the ID token is equal to one of your app's client IDs
		"azp":            "35858082528-b3ekj6u835sd31q4almrtuslaulpgcrm.apps.googleusercontent.com",
		"email":          "rahulsinha4698@gmail.com",
		"email_verified": "true",
		"exp":            "1.695367363e+09", //The expiry time (exp) of the ID token should be there.
		"family_name":    "Sinha",
		"given_name":     "Rahul",
		"iat":            "1.695363763e+09",
		"iss":            "accounts.google.com", //The value of iss in the ID token is equal to accounts.google.com or https://accounts.google.com.
		"jti":            "bb1a8e58f0683e5c07a5cebefb67400ea0c11d8e",
		"locale":         "en",
		"name":           "Rahul Sinha",
		"nbf":            "1.695363463e+09",
		"picture":        "https://lh3.googleusercontent.com/a/ACg8ocLcdO4hOy9GMuc_VQcBar2-x59xTh1ZdkOVuV9TFMl3kq8=s96-c",
		"sub":            "108434034775072407722", //save this in database to identify user as google_login_id unique for every user
	}
	for key, value := range mapCheck {
		fmt.Println(key, ":= ", value)
	}
}
