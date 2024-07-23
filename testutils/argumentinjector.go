package testutils

import (
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/mock"
)

// NewArgumentInjector create a mock.MatchedBy able to set a value when the target func call the stubby.
// It can be used return a value when the target call a function with a pointer parameter.
func NewArgumentInjector(argument interface{}) interface{} {
	return mock.MatchedBy(func(param interface{}) bool {
		ptrParam := &param
		jsonData, _ := json.Marshal(argument)
		return json.Unmarshal(jsonData, &ptrParam) == nil
	})
}
