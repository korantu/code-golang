package kdl

import "testing"

// Must is a temporary replacement instead of not-checked-in tools.
func Must( t * testing.T, _ string, tests ... bool ) {
     for _, result := range tests {
     	 if ! result {
	    t.Fail()
	 }
     }
}