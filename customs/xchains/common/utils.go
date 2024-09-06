package common

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseAttributeValue(value string) any {
	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		return value[1 : len(value)-1]
	}
	if strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]") {
		items := strings.Split(value[1:len(value)-1], ",")
		hexValue := "0x"
		for _, item := range items {
			v, e := strconv.Atoi(item)
			if e == nil {
				hexValue = fmt.Sprintf("%s%x", hexValue, v)
			}
		}
		return hexValue
	}
	return value
}
