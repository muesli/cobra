package doc_test

import (
	"bytes"
	"fmt"

	"github.com/muesli/coral"
	"github.com/muesli/coral/doc"
)

func ExampleGenManTree() {
	cmd := &coral.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	coral.CheckErr(doc.GenManTree(cmd, header, "/tmp"))
}

func ExampleGenMan() {
	cmd := &coral.Command{
		Use:   "test",
		Short: "my test program",
	}
	header := &doc.GenManHeader{
		Title:   "MINE",
		Section: "3",
	}
	out := new(bytes.Buffer)
	coral.CheckErr(doc.GenMan(cmd, header, out))
	fmt.Print(out.String())
}
