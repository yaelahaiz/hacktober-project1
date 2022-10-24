package main

import (
	"fmt"
	"strings"
)

func main() {

	var value = "kdowakdowa"
	path := value
	if hashSpace := strings.Contains(path, " "); hashSpace {
		rmSpaceRight := strings.TrimRight(path, " ")
		replaceDash := strings.ReplaceAll(rmSpaceRight, " ", "-")
		path = strings.ToLower(replaceDash)
	} else {
		path = strings.ToLower(value)
	}

	fmt.Println(path)
}
