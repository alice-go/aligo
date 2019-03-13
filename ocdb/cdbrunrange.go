package ocdb

import "fmt"

func (rr RunRange) String() string {
	return fmt.Sprintf("RunRange{First: %d, Last: %d}", rr.First, rr.Last)
}
