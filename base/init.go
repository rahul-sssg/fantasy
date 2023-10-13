package base

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()                     // Load will read your env file(s) and load them into ENV for this process.
	fmt.Println(os.Getenv("TEST_CRED")) // example test line
}
