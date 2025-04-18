package main

import (
	"fmt"

	"github.com/smtx/smtv/utils"

	arg "github.com/alexflint/go-arg"
)

// Command line arguments parser
var CmdArgs struct {
	Files   []string `arg:"positional,required" help:"List of files to check. Supports glob patterns."`
	Verbose bool     `arg:"-v,--verbose" help:"Enable detailed output for debugging or analysis."`
}

func main() {
	arg.MustParse(&CmdArgs)

	// @TODO improve matching of files
	files := utils.GetFilesToCheck(CmdArgs.Files)

	for _, file := range files {
		fmt.Println(file)
	}
}
