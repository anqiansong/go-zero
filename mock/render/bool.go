package render

import (
	"time"
)

func RandomBool(_ *MockContext) interface{} {
	return time.Now().Unix()%2 == 0
}
