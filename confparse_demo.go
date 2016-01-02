package main

import (
	"fmt"
	"github.com/kaniini/go-confparse"
)

func WalkEntry(ceptr *confparse.ConfigEntry, depth int) {
	var doSection = ceptr.Entries != nil
	var depthStr = ""

	for i := 0; i < depth; i++ {
		depthStr += "\t"
	}

	fmt.Printf("%s%s '%s'", depthStr, ceptr.VarName, ceptr.VarData)
	if doSection {
		fmt.Printf(" {")
	} else {
		fmt.Printf(";")
	}
	fmt.Printf("\n")

	if (ceptr.Entries != nil) {
		WalkEntry(ceptr.Entries, depth + 1);
	}

	if doSection {
		fmt.Printf("%s}\n\n", depthStr)
	}
}

func WalkFile(cfptr *confparse.ConfigFile) {
	fmt.Printf("file ==> %s\n", cfptr.Filename)

	for ceptr := cfptr.Entries; ceptr != nil; ceptr = ceptr.Next {
		WalkEntry(ceptr, 0)
	}
}

func main() {
	cfh := confparse.LoadConfigFile("example.conf")

	for cf := cfh; cf != nil; cf = cf.Next {
		WalkFile(cf)
	}
}
