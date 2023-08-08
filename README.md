# OpenAIAuth
Fetch access tokens for chat.openai.com

## Python version
```py
from OpenAIAuth import Auth0
auth = Auth0(email_address="example@example.com", password="example_password")
access_token = auth.get_access_token()
```

## Go version
```go
package main

import (
	"fmt"
	"os"

	"github.com/xqdoo00o/OpenAIAuth/auth"
)

func main() {
	auth := auth.NewAuthenticator(os.Getenv("OPENAI_EMAIL"), os.Getenv("OPENAI_PASSWORD"), os.Getenv("PROXY"))
	err := auth.Begin()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		return
	}
	token, err := auth.GetAccessToken()
	if err.Error != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		return
	}
	fmt.Println(token)
}
```

## Credits
- @linweiyuan
- @rawandahmad698
- @pengzhile
