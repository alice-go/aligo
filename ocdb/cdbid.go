package ocdb

import "fmt"

func (id ID) Runs() RunRange {
	return id.runs
}
func (id ID) String() string {
	return fmt.Sprintf("AliCDBId{Path: %v, RunRange: %v, Version: 0x%x, SubVersion: 0x%x, Last: %q}",
		id.path, id.runs, id.vers, id.subvers, id.last,
	)
}
