package ocdb

import "fmt"

func (p Path) String() string {
	return fmt.Sprintf("Path{Path: %q, Level0: %q, Level1: %q, Level2: %q, Valid: %v, WildCard: %v}",
		p.path, p.lvl0, p.lvl1, p.lvl2, p.valid, p.wildcard,
	)
}
