package stack

import (
        // "context"
        // "encoding/json"
        "fmt"
        "github.com/gofiber/fiber/v2"
        "net/http"
        "time"
				"strings"
				"strconv"
				// "github.com/robertocamp/gd5gofcd/api/presenter"
				"github.com/robertocamp/gd5gofcd/pkg/entities"
)

// smoke test Hello endpoint
/* func GetStack(c *fiber.Ctx)  error {
        return c.SendString("Hello, World 👋!")
} */

const URL string = "https://api.stackexchange.com/2.2/search?order=desc&sort=activity&intitle=fiber&site=stackoverflow"


// custom unmarshaller for time format
// type EpochConversion string
// func(ec *EpochConversion) UnmarshalJSON(data []byte) error {
// 	t := strings.Trim(string(data), `"`)
// 	sec, err := strconv.ParseInt(t, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	epochTime := time.Unix(sec, 0).Format(time.RFC3339)
// 	// epochTime := time.Unix(sec, 0).Format(time.ANSIC)
// 	*ec = EpochConversion(epochTime)
// 	return nil
// }



// new implementation of GetFullStack
func (c *Client) getStack(f *fiber.Ctx) (*entities.Stack, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?order=desc&sort=activity&intitle=fiber&site=stackoverflow", c.baseURL), nil)
	if err != nil {
		return nil, err
	}
  req.Header.Set("Content-Type", "application/json; charset=utf-8")
  res := entities.Stack{}
	if err := c.sendRequest(req, &res); err !=nil {
		return nil, err
	}
   // swap example code return with endpoint code
	return &res, nil
}

// func GetStack(c *fiber.Ctx)  error {
//         return c.SendString("Hello, World 👋!")
// }
