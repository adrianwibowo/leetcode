package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {

	arrStr := []string{"adrian", "adroid", "adrimas"}
	fmt.Println(longestCommonPrefix(arrStr))
}

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""

	if len(strs) > 0 {
		sort.Strings(strs) // sor alphabetically so we can just compare frist and last string
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		fmt.Println(strs)
		for i := 0; i < len(first); i++ {
			if string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				return longestPrefix
			}
		}
	}
	return longestPrefix

}
