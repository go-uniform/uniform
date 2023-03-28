package uniform

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"strings"
)

var ParseRequest = func(data []byte) (Request, error) {
	var request Request

	if err := bson.Unmarshal(data, &request); err != nil {
		return request, err
	}

	switch value := request.Model.(type) {
	case primitive.Binary:
		request.Model = value.Data
		break
	}

	return request, nil
}

/*
	IndexOf

Get the index of a string inside of a string array
*/
var IndexOf = func(haystack []string, needle string, caseSensitive bool) int {
	if caseSensitive {
		for i, item := range haystack {
			if item == needle {
				return i
			}
		}
	} else {
		lowNeedle := strings.ToLower(needle)
		for i, item := range haystack {
			if strings.ToLower(item) == lowNeedle {
				return i
			}
		}
	}
	return -1
}

// see if a string array contains a given string
var Contains = func(haystack []string, needle string, caseSensitive bool) bool {
	return IndexOf(haystack, needle, caseSensitive) != -1
}

// trim the filterItems from the items array
var Filter = func(items, filters []string, caseSensitive bool) []string {
	if filters == nil || len(filters) <= 0 {
		return items
	}

	filteredItems := make([]string, 0)
	for _, item := range items {
		if Contains(filters, item, caseSensitive) {
			continue
		}
		filteredItems = append(filteredItems, item)
	}

	return filteredItems
}

// private function required for circular reference
func isEmpty(object interface{}) bool {
	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return isEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// determine if a value is empty or not
var IsEmpty = func(value interface{}) bool {
	if value == nil || value == "" {
		return true
	}

	stringValue := strings.TrimSpace(fmt.Sprint(value))
	if stringValue == "" || stringValue == "nil" || stringValue == "<nil>" {
		return true
	}

	if strings.HasPrefix(stringValue, "0001-01-01") {
		return true
	}

	return isEmpty(value)
}

// generate a sha512 hash for the given value/salt combo
var Hash = func(value interface{}, salt string) string {
	concatenatedData := []byte(fmt.Sprintf(`%s%v`, salt, value))

	encoder := sha512.New()
	_, err := encoder.Write(concatenatedData)
	if err != nil {
		panic(err)
	}
	hashedData := encoder.Sum(nil)

	base64EncodedString := base64.StdEncoding.EncodeToString(hashedData)
	return base64EncodedString
}
