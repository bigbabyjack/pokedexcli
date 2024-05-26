package urlutils

import (
	"fmt"
	"log"
	"net/url"
)

func get(s string) {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}
