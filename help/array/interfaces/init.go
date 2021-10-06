package interfaces

import (
	"reflect"
)

func Find(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}

func Contains(thingList []interface{}, thing interface{}) bool {
	for _, element := range thingList {
		if thing == element {
			return true
		}
	}
	return false
}

func Filter(thingList []interface{}, cond func(interface{}) bool) []interface{} {
	var result = make([]interface{}, 0)
	for index, element := range thingList {
		if cond(element) {
			result = append(result, thingList[index])
		}
	}
	return result
}

func Appender(thingList []interface{}, thing interface{}) []interface{} {
	return append(thingList, thing)
}
