Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g., /cmd/myapp).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, put that code in the /internal directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.




tests example


// tests

// package main

// import (
//     "net/httapplication
//     "net/http/httptest"
//     "os"
//     "testing"

//     "github.com/stretchr/testify/require"
// )

// // executeRequest, creates a new ResponseRecorder
// // then executes the request by calling ServeHTTP in the router
// // after which the handler writes the response to the response recorder
// // which we can then inspect.
// func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
//     rr := httptest.NewRecorder()
//     s.Router.ServeHTTP(rr, req)

//     return rr
// }

// // checkResponseCode is a simple utility to check the response code
// // of the response
// func checkResponseCode(t *testing.T, expected, actual int) {
//     if expected != actual {
//         t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//     }
// }

// func TestHelloWorld(t *testing.T) {
//     // Create a New Server Struct
//     s := CreateNewServer()
//     // Mount Handlers
//     s.MountHandlers()

//     // Create a New Request
//     req, _ := http.NewRequest("GET", "/", nil)

//     // Execute Request
//     response := executeRequest(req, s)

//     // Check the response code
//     checkResponseCode(t, http.StatusOK, response.Code)

//     // We can use testify/require to assert values, as it is more convenient
//     require.Equal(t, "Hello World!", response.Body.String())
// }

// go test ./... -v -cover -- command to run all the test cases.
