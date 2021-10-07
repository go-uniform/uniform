package uniform

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strings"
)

// get the index of a string inside of a string array
func IndexOf(haystack []string, needle string, caseSensitive bool) int {
	if haystack == nil {
		panic("specify an array to search through")
	}
	if needle == "" {
		panic("specify a string to search for")
	}
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
func Contains(haystack []string, needle string, caseSensitive bool) bool {
	if haystack == nil {
		panic("specify an array to search through")
	}
	if needle == "" {
		panic("specify a string to search for")
	}
	return IndexOf(haystack, needle, caseSensitive) != -1
}

// trim the filterItems from the items array
func Filter(items []string, filterItems []string) []string {
	if filterItems == nil || len(filterItems) <= 0 {
		return items
	}
	newItems := make([]string, 0)
	for _, item := range items {
		if Contains(filterItems, item, false) {
			continue
		}
		newItems = append(newItems, item)
	}
	return newItems
}

// generate a sha512 hash for the given value/salt combo
func Hash(value interface{}, salt string) string {
	encoder := sha512.New()
	encoder.Write([]byte(fmt.Sprintf(`%s%v`, salt, value)))
	return base64.StdEncoding.EncodeToString(encoder.Sum(nil))
}