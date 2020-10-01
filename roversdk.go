package roversdk

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// IsFileExists check file whether exist
func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	// path/to/whatever exist (true) or false for a directory
	return !info.IsDir()
}

// IsDirExists chek directory whether exist
func IsDirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return false
	}
	// path/to/whatever exists
	return true
	// if _, err := os.Stat("/path/to/whatever"); !os.IsNotExist(err) {
	// 	// path/to/whatever exists
	// }
}

// DumpMapSS dumps map[string]string
func DumpMapSS(m map[string]string) {
	var maxLenKey int
	for k := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range m {
		ks := fmt.Sprintf("%q", k)
		vs := fmt.Sprintf("%q", v)
		fmt.Println(ks + strings.Repeat(" ", maxLenKey-len(k)) + " => " + vs)
	}
}

// Dump dumps variable such as struct type
func Dump(s ...interface{}) {
	spew.Dump(s)
}
