package ocdb

import (
	"fmt"
	"io"
)

func (meta *MetaData) Display(w io.Writer) {
	fmt.Fprintf(w, "Class: %q\nResponsible: %q\nBeamPeriod: %d\nAliRoot Version: %q\nComment: %q\nProperties: %d\n",
		meta.class, meta.resp, meta.beam, meta.vers, meta.comment, len(meta.props.Table()),
	)
	for k, v := range meta.props.Table() {
		fmt.Fprintf(w, "  key: %v\n  val: %v\n", k, v)
	}
}
