package variable

import "fmt"

// endPointString is a type for endpoint string
type endPointString string

// endPointFunc is a type for returning endpoint string with argument
type endPointFunc func(arg string) endPointString

// APIEndpoint is an interface for endpoint
type APIEndpoint interface {
	endPointString | endPointFunc
}

// An enumeration that configures APIEndpoint
var (
	USER_LIST   endPointString = "/user"
	USER_DETAIL endPointFunc   = func(userID string) endPointString {
		return endPointString(fmt.Sprintf("/user/%s", userID))
	}
	ORDER_LIST   endPointString = "/order"
	ORDER_DETAIL endPointFunc   = func(orderID string) endPointString {
		return endPointString(fmt.Sprintf("/order/%s", orderID))
	}
)
