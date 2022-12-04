package krpc_test

import (
	"fmt"
	"testing"

	sturdyengine "github.com/ddouglas/go-krpc"
)

var scc, _ = sturdyengine.NewDefaultConnection()
var sc, e = sturdyengine.NewSpaceCenter(&scc)

func TestInitSpaceCenter(t *testing.T) {
	sc, e = sturdyengine.NewSpaceCenter(&scc)
	if e != nil {
		fmt.Println(e)
		t.FailNow()
	}
}
