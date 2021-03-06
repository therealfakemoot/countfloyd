package main

import (
	"context"
	"os"
	"path"

	"github.com/Laughs-In-Flowers/flip"
	"github.com/Laughs-In-Flowers/log"
)

var (
	versionPackage string = path.Base(os.Args[0])
	versionTag     string = "No Tag"
	versionHash    string = "No Hash"
	versionDate    string = "No Date"
	C              *flip.Commander
	L              log.Logger
	currentDir     string
)

func main() {
	ctx := context.Background()
	C.Execute(ctx, os.Args)
	os.Exit(0)
}

func init() {
	currentDir, _ = os.Getwd()
	n := path.Base(os.Args[0])
	log.SetFormatter("countfloyd_text", log.MakeTextFormatter(n))
	C = flip.BaseWithVersion(versionPackage, versionTag, versionHash, versionDate)
	L = log.New(os.Stdout, log.LInfo, log.DefaultNullFormatter())
	flip.RegisterGroup("top", -1, TopCommand())
	flip.RegisterGroup("control", 1, StartCommand(), StopCommand(), StatusCommand(), QueryCommand())
	flip.RegisterGroup("action", 2, PopulateCommand(), ApplyCommand())
}
