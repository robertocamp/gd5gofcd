# gd5gofcd --"go fiber clean design"

# links
## fiber clean design recipe:
- https://github.com/gofiber/recipes/tree/master/clean-architecture
## setup
`git clone`
`cd gd5gofcd`
`go mod init github.com/robertocamp/gd5gofcd`
## mongodb on mac
- https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-os-x/

## adding endpoints to the API:
1. edit routes.go and add the endpoint and the handler
  + if the endpoint is not part of `book_handler.go`, add a new routes file in app.go
  + `routes.StackRouter(api stackService)`
  + new routes will require a new service



## stack JSON API
- https://api.stackexchange.com/2.2/search?order=desc&sort=activity&intitle=perl&site=stackoverflow
- https://api.stackexchange.com/2.2/search?order=desc&sort=activity&intitle=fiber&site=stackoverflow
### stack API integration to gd5gofcd
  1. working client with struct and json repsonse: /Users/robert/go/src/github.com/learning-go/stack/stack.go
  2. run the `stack.go` program and then copy the returned string as a file in the "json" folder
  3. convert this into an endpoint in the progam: "/stack"
## API testing
### run local: 
- `go run app.go`
- `http://localhost:3000/api/stack`
- look for services running on port 3000: `lsof -nP -iTCP -sTCP:LISTEN | grep 3000`

### unit tests:
https://assertible.com/blog/testing-and-validating-api-responses-with-json-schema
https://stackoverflow.com/questions/8292050/is-there-any-publicly-accessible-json-data-source-to-test-with-real-world-data
http://www.jsontest.com/
https://www.tumblr.com/docs/en/api/v1
https://api.stackexchange.com/2.2/search?order=desc&sort=activity&intitle=perl&site=stackoverflow
https://httpbin.org/
