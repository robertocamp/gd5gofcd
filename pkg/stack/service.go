package stack

import (
	"github.com/robertocamp/gd5gofcd/api/presenter"
	"github.com/robertocamp/gd5gofcd/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"fmt"
)

//Service is an interface from which our api module can access the stackOverflow API data
// struct data for the API is contained in the "entities" pkg
type Service interface {
	GetPlayerScore(name string) int
	//  returns the data to the presenter layer
  NewStack()   (*[]presenter.Stack, error)
}


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
func (c *Client) NewStack() (*entities.Stack, error) {
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


