package main

// sedum:anchor:imports
// sedum:addImport {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/repin_run.go","path":"fmt"}}
import "fmt"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/repin_run.go","path":"io"}}
import "io"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/repin_run.go","path":"time"}}
import "time"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/repin_run.go","path":"github.com/spf13/cobra"}}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","kwargs":{"file":"cmd/provenance/repin_run.go","path":"github.com/livecodelife/provenance/stale"}}
import "github.com/livecodelife/provenance/stale"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"tier":"owned","record":"prov-2026-9bfa6107","gen":"sha256:e937012cd7f04b4dea168e2f24be21b5aadd0783afbe0dda58babc69031e3dbc","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","kwargs":{"code":"func runRepin(out io.Writer, cmd *cobra.Command, args []string) error {\n\tr, err := openRepo(cmd)\n\tif err != nil {\n\t\treturn err\n\t}\n\tid, err := recordArg(cmd, args)\n\tif err != nil {\n\t\treturn err\n\t}\n\trec, _, err := r.resolve(id)\n\tif err != nil {\n\t\treturn err\n\t}\n\treason, _ := cmd.Flags().GetString(\"reason\")\n\tto, _ := cmd.Flags().GetInt(\"to\")\n\tlog := stale.FilePinLog{Root: r.Root, Dir: r.Config.Dir}\n\tshared, _ := loadShared(r)\n\tpin, err := stale.RepinShared(r.Store, shared, log, rec.ID, to, reason, r.author(), time.Now().UTC())\n\tif err != nil {\n\t\tr.fail(\"Cannot repin %s: %v\", rec.ID, err)\n\t\treturn err\n\t}\n\tfmt.Fprintf(out, \"\\n✓ %s now pinned to %s@%d (was %d)\\n\\n  No new version was written. Commit with the tag %s\\n\\n\", rec.ID, pin.Parent, pin.To, pin.From, tag(rec.ID, rec.Version))\n\treturn nil\n}","file":"cmd/provenance/repin_run.go","name":"runRepin"}}
func runRepin(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	id, err := recordArg(cmd, args)
	if err != nil {
		return err
	}
	rec, _, err := r.resolve(id)
	if err != nil {
		return err
	}
	reason, _ := cmd.Flags().GetString("reason")
	to, _ := cmd.Flags().GetInt("to")
	log := stale.FilePinLog{Root: r.Root, Dir: r.Config.Dir}
	shared, _ := loadShared(r)
	pin, err := stale.RepinShared(r.Store, shared, log, rec.ID, to, reason, r.author(), time.Now().UTC())
	if err != nil {
		r.fail("Cannot repin %s: %v", rec.ID, err)
		return err
	}
	fmt.Fprintf(out, "\n✓ %s now pinned to %s@%d (was %d)\n\n  No new version was written. Commit with the tag %s\n\n", rec.ID, pin.Parent, pin.To, pin.From, tag(rec.ID, rec.Version))
	return nil
}

// /sedum:setFunction
