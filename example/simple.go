package main

import (
	"os"
	"time"

	"github.com/ddouglas/go-krpc"
)

func main() {
	conn, err := krpc.NewDefaultConnection()
	if err != nil {
		os.Exit(1)
	}
	sc, err := krpc.NewSpaceCenter(&conn)
	if err != nil {
		os.Exit(1)
	}

	vessel, err := sc.NewVessel()
	if err != nil {
		os.Exit(1)
	}

	sc.Quicksave()

	vessel.SetThrottle(0.75)
	// vessel.ActivateNextStage()
	time.Sleep(10 * time.Second)
	// sc.Quickload()

}
