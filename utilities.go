package uniform

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strings"
)

/* IndexOf
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