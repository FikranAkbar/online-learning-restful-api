package helper

import (
	"fmt"
	"reflect"
	"strings"
)

func StringModel(model any) string {
	builder := strings.Builder{}

	v := reflect.ValueOf(model)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		builder.WriteString(
			fmt.Sprintf("%s -> Value: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface()),
		)
	}

	return builder.String()
}
