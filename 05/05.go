package p5

import (
	"log"
	"strings"
)

// ReplaceSpace 替换空格为%20
func ReplaceSpace(s string) string {
	r := strings.Replace(s, " ", "%20", -1)
	log.Println(r)
	return r
}
