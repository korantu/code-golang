package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"dots"
)

var visual = flag.Bool("visual", false, "Output visual information")
var updated = flag.Bool("v2", true, "Use second version of the scanner")


var stats = 0
func count_parenthis(path string) int {
	result := 0
	is_comment := false
	is_string := false
	do_escape := false
	if bytes, err := ioutil.ReadFile(path); err == nil {
		for _, byte := range bytes {
			// String
			switch {
			case is_comment: // Do nothing
			case do_escape:
				do_escape = false
				continue
			case byte == '\\': // Not the best way...
				do_escape = true
			case byte == '"':
				is_string = ! is_string
			}

			// Comment
			switch {
			case is_string: // Do nothing
			case byte == '\n':
				is_comment = false
			case byte == ';':
				is_comment = true
			}

			// Counting
			switch {
			case is_string:
			case is_comment: // Do nothing
			case byte == '(':
				stats++
				result++
			case byte == ')':
				result--
			}
		}
	}
	return result
}


type ANoter struct { Sexps int }
func (a *ANoter) Up() { }
func (a *ANoter) Down() { a.Sexps+=1 }
func (a *ANoter) Token([]byte){}

func main() {
	where := flag.String("place", ".", "Where to look")
	flag.Parse()

	print("<<" + *where + ">>\n")

	noter := &ANoter{}
	byte_count := 0

	old_walker := func(path string, info os.FileInfo, _ error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(path, ".el") {
				count := count_parenthis(path)
				if *visual || count != 0 {
					fmt.Println(path, ":", count)
				}
			
		}
		return nil
	};

	 new_walker := func(path string, info os.FileInfo, _ error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(path, ".el") {
				if bytes, err := ioutil.ReadFile(path); err == nil {
					dots.ElispScanner(bytes).Scan( noter)
				byte_count += len(bytes)
				}
		}
		return nil
	}

	if ( *updated) {
		filepath.Walk(*where, new_walker)
		fmt.Println("Total ", noter, " sexps, new way; ", byte_count, " bytes.")
	} else {
		filepath.Walk(*where, old_walker)
		fmt.Println("Total ", stats, " sexps, old way.")
	}
}
