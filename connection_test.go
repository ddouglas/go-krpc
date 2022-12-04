package krpc_test

import (
	"fmt"
	"testing"

	sturdyengine "github.com/ddouglas/go-krpc"
)

func TestInitilization(t *testing.T) {
	_, e := sturdyengine.NewConnection("testing", "50000")
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}

func TestDefault(t *testing.T) {
	_, e := sturdyengine.NewDefaultConnection()
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}

func TestClose(t *testing.T) {
	c, e := sturdyengine.NewDefaultConnection()
	if e != nil {
		t.FailNow()
	}
	e = c.Close()
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}
