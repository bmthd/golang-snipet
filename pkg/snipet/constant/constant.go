// Description: This package is for constant values.
package constant

// Status is a struct for status code and name
type Status struct {
	code uint8
	name string
}

var (
	OK = Status{0, "OK"}
	NG = Status{1, "NG"}
)

// Code returns status code
func (s Status) Code() uint8 {
	return s.code
}

// Name returns status name
func (s Status) Name() string {
	return s.name
}

// GetStatusList returns status list
func GetStatusList() []Status {
	list := make([]Status, 0)
	list = append(list, OK)
	list = append(list, NG)
	return list
}

// GetNameByCode returns status name by code
func GetNameByCode(code uint8) string {
	for _, status := range GetStatusList() {
		if status.Code() == code {
			return status.Name()
		}
	}
	return ""
}
