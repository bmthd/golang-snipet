package structure

import "fmt"

// endPointString is a type for endpoint string
type endPointString string

func (e endPointString) IsFunc() bool {
	return false
}

// endPointFunc is a type for returning endpoint string with argument
type endPointFunc func(arg string) endPointString

func (e endPointFunc) IsFunc() bool {
	return true
}

// APIEndpoint is an interface for endpoint
type APIEndpoint interface {
	endPointString | endPointFunc
	// IsFunc returns true if the endpoint is a function
	IsFunc() bool
}

// var (
// 	USER_LIST   = EndPointString("/user")
// 	USER_DETAIL = EndPointFunc(
// 		func(userID string) string {
// 			return fmt.Sprintf("/user/%s", userID)
// 		})
// )

// var END_POINT = map[string]interface{}{
// 	"USER_LIST": EndPointString("/user"),
// 	"USER_DETAIL": EndPointFunc(func(userID string) EndPointString {
// 		return EndPointString(fmt.Sprintf("/user/%s", userID))
// 	}),
// }

var END_POINT = struct {
	USER_LIST   endPointString
	USER_DETAIL endPointFunc
}{
	USER_LIST: endPointString("/user"),
	USER_DETAIL: endPointFunc(func(userID string) endPointString {
		return endPointString(fmt.Sprintf("/user/%s", userID))
	}),
}
