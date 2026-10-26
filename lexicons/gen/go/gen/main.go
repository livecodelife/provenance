// Command gen writes cbor_gen.go for the generated lexicon types. Run it with
// lexicons/gen/regen.sh; it is the CBOR step lexgen leaves to cbor-gen.
package main

import (
	"fmt"
	"os"

	cbg "github.com/whyrusleeping/cbor-gen"

	provenancelex "github.com/livecodelife/provenance/lexicons/gen/go"
)

func main() {
	if err := cbg.WriteMapEncodersToFile("cbor_gen.go", "provenancelex",
		provenancelex.Defs_Content{},
		provenancelex.Defs_Spec{},
		provenancelex.Project{},
		provenancelex.Record{},
		provenancelex.Record_Source{},
		provenancelex.Version{},
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
