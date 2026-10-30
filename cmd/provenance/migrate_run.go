package main

// sedum:anchor:imports
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"encoding/json"},"kwargs":"sha256:7c157007712dce97e7bdea05711c286015939516b6b7cf1c855521dcbfec7213","body":"sha256:872d7c3e9dee608e79a7f8ef57930d24b5e30f2a57d891a52549179ace45946a"}
import "encoding/json"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"errors"},"kwargs":"sha256:6706692f262c062bbf50932bee0f02b40642603740eb68025f5e8c8e9b38a164","body":"sha256:2b718237ba0a56467c8fd494cc131e942ef410e6228c03b644d70476c71862dc"}
import "errors"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"fmt"},"kwargs":"sha256:7a655e1af2cc9525d370bbb5f0f3e58f74a4096c6fcad5ceef81a344aa0b7bc5","body":"sha256:bcbac51883cbf8064e9196b371cb880cee429866652447ad557533d817857609"}
import "fmt"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"io"},"kwargs":"sha256:1b5a7700ea026701c21bc3e692c1c8f225b0bc625a46ed093635fa92642bfb8e","body":"sha256:ef8e8abf4ef5ee40a19dc85b49bf09c5df1fe32abfe834e0ab4c2175baf80ddf"}
import "io"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"os"},"kwargs":"sha256:2f77b65a4d083c55e923a283a73764bc75f6cd93b6f2f757b218e4314ad10ef7","body":"sha256:7c6e309eb2575d4aa451ec2ab83161d13e9ae03bbca3a4d2ecf30ce12adf0905"}
import "os"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"path/filepath"},"kwargs":"sha256:ceb5c163ad6c98f37228a9a5292f2dde3db8d95c7d2274efdf9b2043442b0f52","body":"sha256:a6d196b558491e646be93f56a321f339068c64c76aa9d49bb18046e7f9c777a1"}
import "path/filepath"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"sort"},"kwargs":"sha256:295a0a2bca8e3b29528a0ac162815999c2aada1176d7ef4d7e68bd34d6c70f68","body":"sha256:f197172617ae21ef23b1855b09aafb14571020ae81bff877d63d3ff81f4bf63c"}
import "sort"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"strings"},"kwargs":"sha256:336fa7c27cdfe6f63c395c21f9635ab54c24bf2d7a0e716cc89419039eb10c3d","body":"sha256:a9af3c0bc4bf9342df9d7e89a324523fdba7941a9ca4190270e68ff3c4a25ea6"}
import "strings"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"github.com/spf13/cobra"},"kwargs":"sha256:9258404ca7221fa676e875f27bb07f4bb4bcc2ff0a9ce2c11f16d20f700ed490","body":"sha256:f361c02d0f8e1999cf3fb715393e16310eeec9debd5aaa13936af7617e2b46aa"}
import "github.com/spf13/cobra"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"gopkg.in/yaml.v3"},"kwargs":"sha256:bc86f6ac6092f209c4515e42eb17660ba708dfe70babe755d4a5e593927c898d","body":"sha256:a4214cbff710690b396b73d2b1ccc9aa8b55a807d07f73aa80f10af0c41e13b9"}
import "gopkg.in/yaml.v3"

// /sedum:addImport
// sedum:addImport {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:d6a9443c14d82b7bdf19bb96637f957fd44ad9bfbe8f167c97e6f7699c8b830f","id":{"file":"cmd/provenance/migrate_run.go","path":"github.com/livecodelife/provenance/migrate"},"kwargs":"sha256:978d9f655631d47bb59c0041ea039d62536cfd7729586ee14da3c61a929b489b","body":"sha256:27e5ef58141132dce0f705931f8a8ad38cd4f42a845dbd54b4003e0a84a6ce1c"}
import "github.com/livecodelife/provenance/migrate"

// /sedum:addImport

// sedum:anchor:decls
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"runMigrate"},"kwargs":"sha256:4e5728e04b1794857ca1dbdc8c022568f0835e47b604a5af8087e08175526904","body":"sha256:d70a1e8c6cbd2b42548c444e30a37cd0295272616e14553979fbceed934b425c"}
func runMigrate(out io.Writer, cmd *cobra.Command, args []string) error {
	r, err := openRepo(cmd)
	if err != nil {
		return err
	}
	planOnly, _ := cmd.Flags().GetBool("plan")
	asJSON, _ := cmd.Flags().GetBool("json")
	accept, _ := cmd.Flags().GetBool("accept-suggestions")
	keep, _ := cmd.Flags().GetStringArray("keep-separate")
	fold, _ := cmd.Flags().GetStringArray("fold")
	workspace, _ := cmd.Flags().GetString("workspace")
	if asJSON && !planOnly {
		return errors.New("--json needs --plan")
	}

	dirs := map[string]string{}
	for _, sr := range r.Config.SharedRepos {
		dirs[sr.Name] = sharedCacheDir(sr)
	}
	src, err := migrate.Load(r.Root, migrate.LoadOptions{Workspace: workspace, SharedDirs: dirs})
	if err != nil {
		return err
	}
	if src.Dir != r.Config.Dir {
		return fmt.Errorf("records directory mismatch: %s config says %q, migration reads %q", r.ConfigPath, r.Config.Dir, src.Dir)
	}
	plan, err := migrate.MakePlan(src)
	if err != nil {
		return err
	}
	opts, err := migrateOptions(plan, accept, keep, fold)
	if err != nil {
		return err
	}

	if planOnly {
		view := planView(plan, opts)
		if asJSON {
			enc := json.NewEncoder(out)
			enc.SetIndent("", "  ")
			return enc.Encode(view)
		}
		printPlan(out, view)
		return nil
	}

	if !needsMigration(src) {
		fmt.Fprintf(out, "\n✓ Nothing to migrate: every record in %s is already v2.\n\n", src.Dir)
		if missing, err := missingAttrLines(r.Root, r.Config.Dir); err == nil && len(missing) > 0 {
			fmt.Fprintf(out, "  Hint: add these lines to .gitattributes: %s\n\n", strings.Join(missing, "; "))
		}
		return nil
	}
	if dirty, err := r.git("status", "--porcelain", "--", r.Config.Dir, legacyConfigFile, configFile); err == nil && dirty != "" {
		r.fail("Cannot migrate: %s has uncommitted changes.\n\n%s\n\n  Commit or stash them first, so the migration is one reviewable change.", r.Config.Dir, dirty)
		return errors.New("dirty records directory")
	}

	res, err := migrate.Apply(plan, opts, r.Store)
	if err != nil {
		return fmt.Errorf("migrate: %w (records written so far: %d; inspect with git status and restore with git checkout)", err, res.Records)
	}
	moved := ""
	if r.Legacy {
		if moved, err = moveConfig(r.Root); err != nil {
			return err
		}
	}
	if err := r.writeManifest(); err != nil {
		return err
	}
	removedManifest := false
	if err := os.Remove(filepath.Join(r.Root, legacyManifest)); err == nil {
		removedManifest = true
	} else if !os.IsNotExist(err) {
		return err
	}

	attrStatus, err := ensureGitattributes(r.Root, r.Config.Dir)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "\n✓ Migrated %s to v2: %d records (from %d), %d snapshots\n\n", src.Dir, res.Records, len(src.Records), res.Snapshots)
	if len(res.Removed) > 0 {
		fmt.Fprintf(out, "  Folded into other records (kept as aliases): %d\n", len(res.Removed))
	}
	if len(res.Repinned) > 0 {
		fmt.Fprintf(out, "  Re-pinned to the final version of a folded parent (pins only, no new version): %d\n", len(res.Repinned))
		for _, rp := range res.Repinned {
			fmt.Fprintf(out, "    %s  %s@%d → @%d\n", rp.ID, rp.Parent, rp.From, rp.To)
		}
	}
	if len(res.SharedPins) > 0 {
		fmt.Fprintf(out, "  Shared parents pinned to the synced version: %d\n", len(res.SharedPins))
		for _, sp := range res.SharedPins {
			fmt.Fprintf(out, "    %s  %s:%s@%d\n", sp.ID, sp.Repo, sp.Parent, sp.Version)
		}
	}
	if moved != "" {
		fmt.Fprintf(out, "  Config: %s\n", moved)
	}
	fmt.Fprintf(out, "  Version manifest: %s\n", manifestPath)
	fmt.Fprintf(out, "  Git attributes: .gitattributes %s (marks snapshots, evidence and the manifest as generated)\n", attrStatus)
	if removedManifest {
		fmt.Fprintf(out, "  Removed the v1 hash manifest: %s\n", legacyManifest)
	}
	fmt.Fprintln(out, "  Hooks: run `provenance hooks install` to replace linespec's (they are kept as .bak).")
	fmt.Fprintln(out)
	return nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"migrateOptions"},"kwargs":"sha256:104b725c1a300dc5f9b5411de696e427b9bb53a58292d12a80500c70cc9a5869","body":"sha256:aef792ae4b23558deedc5bae21fa6604ab2f3b05ee1b617317d4aba9064a3c65"}
// migrateOptions turns the keep-separate flags into migrate.Options. Every
// id they name must be an absorbed record in the plan.
func migrateOptions(p *migrate.Plan, accept bool, keep, fold []string) (migrate.Options, error) {
	opts := migrate.Options{AcceptSuggestions: accept, KeepSeparate: map[string]bool{}}
	for _, set := range []struct {
		ids  []string
		keep bool
	}{{keep, true}, {fold, false}} {
		for _, id := range set.ids {
			if _, ok := p.Entries[id]; !ok {
				return opts, fmt.Errorf("%s is not folded into another record by this plan", id)
			}
			if prev, dup := opts.KeepSeparate[id]; dup && prev != set.keep {
				return opts, fmt.Errorf("%s is named by both --keep-separate and --fold", id)
			}
			opts.KeepSeparate[id] = set.keep
		}
	}
	return opts, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"needsMigration"},"kwargs":"sha256:84d53f1d1cf7b22d5440186a701af059fa003b91b8ffae43bf9debb199934a7d","body":"sha256:034bd656df8af150eb7f4fd6f0b55ad518b90e942214588bf6b38b619451d688"}
// needsMigration reports whether any local record is still schema 1.
func needsMigration(src *migrate.Repo) bool {
	for _, s := range src.Records {
		if s.Rec.Schema != 2 {
			return true
		}
	}
	return false
}

// /sedum:setFunction
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/migrate_run.go","name":"planJSON"},"kwargs":"sha256:dfd9e8739bf2e6e06d5bac522b5a2b2a2bdfefb1784e83f782a057664d7b29e5","body":"sha256:c48baf23d4b3ba5bae2e29ae21a52ff72ed5b4e766cc7955173225a89ca52edd"}
// planJSON is what `migrate --plan --json` prints.
type planJSON struct {
	Repo           string              `json:"repo"`
	RecordsBefore  int                 `json:"records_before"`
	RecordsAfter   int                 `json:"records_after"`
	Chains         map[string][]string `json:"chains"`
	Entries        []entryJSON         `json:"entries"`
	Suggestions    []suggestionJSON    `json:"suggestions"`
	KeptDeprecated []keptJSON          `json:"kept_deprecated"`
	StatusMappings []map[string]any    `json:"status_mappings"`
	Warnings       []map[string]any    `json:"warnings"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/migrate_run.go","name":"entryJSON"},"kwargs":"sha256:cb6c421c01acbfb972ba5a6a5d232cfccaaf74a94d3fc997ecdaf83efbd1eb0c","body":"sha256:5cc17cc4fedb93f229ce8204321ef19741ba3e7b2ae7e2572813e6572a5df4e3"}
type entryJSON struct {
	ID      string `json:"id"`
	Target  string `json:"target"`
	Rule    string `json:"rule"`
	Suggest string `json:"suggest,omitempty"`
	Kept    bool   `json:"kept"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/migrate_run.go","name":"suggestionJSON"},"kwargs":"sha256:0abdbd07e5c6e712e778d2a4f7afc76fd251b9c5fe7e0d3a6f4d69a947dcfe11","body":"sha256:7d59ed9244e89348592c70bed406a6962f9950980f304ffd6b100551f3939133"}
type suggestionJSON struct {
	ID     string `json:"id"`
	Target string `json:"target"`
	Rule   string `json:"rule"`
	Why    string `json:"why"`
}

// /sedum:setType
// sedum:setType {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:ebb283283a745fdc09f0294b6e79bd879a830829cb964550ec6b4432fee5bb72","id":{"file":"cmd/provenance/migrate_run.go","name":"keptJSON"},"kwargs":"sha256:67fc7646179a4e63743e3b571b8449d881a9538755d11d4fd8f027e409e70683","body":"sha256:bdca0e6caa67cbb7e2b116c10b3f867ced0f57a8cca4a35faad6892fc011a46a"}
type keptJSON struct {
	ID         string `json:"id"`
	Extends    string `json:"extends"`
	Implements string `json:"implements,omitempty"`
}

// /sedum:setType
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"planView"},"kwargs":"sha256:753e4b910581ce96c96d36a2e955f260cf831bee4884f52ce038b40886edaa66","body":"sha256:86a66604ebf2375255be92699a81b3c33443d31ab69cf282e3e5b7f6cd73d592"}
// planView is the plan under opts: a kept-separate id leaves its chain and
// stands on its own.
func planView(p *migrate.Plan, opts migrate.Options) planJSON {
	v := planJSON{
		Repo: p.Repo.Name, RecordsBefore: len(p.Repo.Records), RecordsAfter: p.RecordsAfter(opts),
		Chains: map[string][]string{}, Entries: []entryJSON{}, Suggestions: []suggestionJSON{},
		KeptDeprecated: []keptJSON{}, StatusMappings: []map[string]any{}, Warnings: []map[string]any{},
	}
	for s, ids := range p.Versions {
		var chain []string
		for _, id := range ids {
			if p.SurvivorWith(id, opts) == s {
				chain = append(chain, id)
			}
		}
		if len(chain) > 1 {
			v.Chains[s] = chain
		}
	}
	ids := make([]string, 0, len(p.Entries))
	for id := range p.Entries {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		for _, e := range p.Entries[id] {
			v.Entries = append(v.Entries, entryJSON{ID: id, Target: e.Target, Rule: e.Rule, Suggest: e.Suggest, Kept: p.SurvivorWith(id, opts) == id})
		}
	}
	for _, s := range p.Suggestions {
		v.Suggestions = append(v.Suggestions, suggestionJSON{ID: s.ID, Target: s.Target, Rule: s.Rule, Why: s.Why})
	}
	for _, k := range p.KeptDeprecated {
		v.KeptDeprecated = append(v.KeptDeprecated, keptJSON{ID: k.ID, Extends: k.Extends, Implements: k.AsImprint.Implements})
	}
	for _, w := range p.StatusMappings(opts) {
		v.StatusMappings = append(v.StatusMappings, warningMap(w))
	}
	for _, w := range p.Warnings {
		v.Warnings = append(v.Warnings, warningMap(w))
	}
	return v
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"warningMap"},"kwargs":"sha256:8e3b61b9ca45efeeb9b6dd09642f76ce419fd3d6aa50d135c9e7acb7e8a48f84","body":"sha256:06e23e16d703e0627b9f6615af0ca91c091e290597ceae4d3dc05fe81285f89e"}
func warningMap(w migrate.Warning) map[string]any {
	m := map[string]any{"kind": w.Kind}
	for k, val := range w.Fields {
		m[k] = val
	}
	return m
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"printPlan"},"kwargs":"sha256:fe30ebce47fd8d92c9eef2b13d38a85b3764e318f28d854a4465f86dccca83ed","body":"sha256:b2201823bce6287046b820c0ecb4520316b020b99b5d7e9b8b4546ce6d64b151"}
func printPlan(out io.Writer, v planJSON) {
	fmt.Fprintf(out, "\nMigration plan for %s: %d records → %d\n", v.Repo, v.RecordsBefore, v.RecordsAfter)

	survivors := make([]string, 0, len(v.Chains))
	for s := range v.Chains {
		survivors = append(survivors, s)
	}
	sort.Strings(survivors)
	fmt.Fprintf(out, "\nVersion chains (%d), oldest version first:\n", len(survivors))
	for _, s := range survivors {
		fmt.Fprintf(out, "  %s  ← %s\n", s, strings.Join(v.Chains[s], ", "))
	}
	if len(v.Suggestions) > 0 {
		fmt.Fprintf(out, "\nSuggested to keep separate (%d); accept with --accept-suggestions or --keep-separate <id>:\n", len(v.Suggestions))
		for _, s := range v.Suggestions {
			fmt.Fprintf(out, "  %s → %s (%s): %s\n", s.ID, s.Target, s.Rule, s.Why)
		}
	}
	if len(v.KeptDeprecated) > 0 {
		fmt.Fprintf(out, "\nDeprecated bugs kept as deprecated imprints (%d):\n", len(v.KeptDeprecated))
		for _, k := range v.KeptDeprecated {
			fmt.Fprintf(out, "  %s (extends %s)\n", k.ID, k.Extends)
		}
	}
	if len(v.StatusMappings) > 0 {
		fmt.Fprintf(out, "\nStatus mappings to review (%d):\n", len(v.StatusMappings))
		for _, m := range v.StatusMappings {
			fmt.Fprintf(out, "  %v: %v → %v\n", m["id"], m["from"], m["to"])
		}
	}
	if len(v.Warnings) > 0 {
		fmt.Fprintf(out, "\nWarnings (%d):\n", len(v.Warnings))
		for _, w := range v.Warnings {
			fmt.Fprintf(out, "  %s\n", formatWarning(w))
		}
	}
	fmt.Fprintln(out, "\nNothing was written. Run `provenance migrate` with the same flags to apply.")
	fmt.Fprintln(out)
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"formatWarning"},"kwargs":"sha256:e767b1ecc5d257b102b9ee56cb10b4ef3a3860d4c535fc7208403a4a55a33434","body":"sha256:a522dacd61807f7c965d33a191e7f4ee02f8b10b62a1d3a2a74a1e4eccbe40e8"}
func formatWarning(w map[string]any) string {
	keys := make([]string, 0, len(w))
	for k := range w {
		if k != "kind" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	parts := []string{fmt.Sprint(w["kind"])}
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%v", k, w[k]))
	}
	return strings.Join(parts, " ")
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"moveConfig"},"kwargs":"sha256:c3f61413a0df7f3a5b4b70e286b77732a62fbbb1f8bee12c87687461199c7fe6","body":"sha256:9265a0bbb11c8547e8b2d64cb65fb630f451e29078b816f47286830e04dd6a41"}
// moveConfig moves the provenance block of .linespec.yml to .provenance.yml,
// dedented to the top level and otherwise byte for byte, and removes it from
// .linespec.yml, deleting that file when nothing else is left in it.
func moveConfig(root string) (string, error) {
	legacy := filepath.Join(root, legacyConfigFile)
	data, err := os.ReadFile(legacy)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(filepath.Join(root, configFile)); err == nil {
		return "", fmt.Errorf("%s already exists; move the provenance block from %s by hand", configFile, legacyConfigFile)
	}
	lines := strings.SplitAfter(string(data), "\n")
	start := -1
	for i, l := range lines {
		if strings.TrimRight(l, " \r\n") == "provenance:" {
			start = i
			break
		}
	}
	if start < 0 {
		return "", fmt.Errorf("%s has no block-style provenance: key to move", legacyConfigFile)
	}
	end := start + 1
	for end < len(lines) {
		l := lines[end]
		trimmed := strings.TrimSpace(l)
		if trimmed != "" && !strings.HasPrefix(trimmed, "#") && l[0] != ' ' && l[0] != '\t' {
			break
		}
		end++
	}
	// Trailing blank and comment lines belong to whatever follows the block.
	for end > start+1 {
		trimmed := strings.TrimSpace(lines[end-1])
		if trimmed != "" && !(strings.HasPrefix(trimmed, "#") && lines[end-1][0] == '#') {
			break
		}
		end--
	}
	block := lines[start+1 : end]
	indent := ""
	for _, l := range block {
		if strings.TrimSpace(l) != "" {
			indent = l[:len(l)-len(strings.TrimLeft(l, " \t"))]
			break
		}
	}
	var moved strings.Builder
	for _, l := range block {
		moved.WriteString(strings.TrimPrefix(l, indent))
	}
	text, err := rewriteExcludePaths(moved.String())
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(root, configFile), []byte(text), 0o644); err != nil {
		return "", err
	}
	rest := strings.Join(append(append([]string{}, lines[:start]...), lines[end:]...), "")
	if strings.TrimSpace(stripComments(rest)) == "" {
		if err := os.Remove(legacy); err != nil {
			return "", err
		}
		return fmt.Sprintf("moved to %s; %s removed (it held only the provenance block)", configFile, legacyConfigFile), nil
	}
	if err := os.WriteFile(legacy, []byte(rest), 0o644); err != nil {
		return "", err
	}
	return fmt.Sprintf("provenance block moved from %s to %s", legacyConfigFile, configFile), nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"rewriteExcludePaths"},"kwargs":"sha256:2cc307af2217d0ee117eb6190715d3e32c21a3822e1cfbd48d89aed5ca0d7b02","body":"sha256:fa7763677403ed5d79f64eac062d9fcdc9e7baaf148b8515668ac4609b1c5e86"}
// rewriteExcludePaths fixes the moved config's exclude_paths for the
// migrated repo: .linespec.yml becomes .provenance.yml (once), and the v1
// hash manifest, which migrate deletes, is dropped, and .gitattributes (which migrate writes) is added when missing. Only that key's lines are
// rewritten, in the list's own style; everything else stays byte for byte.
func rewriteExcludePaths(text string) (string, error) {
	var doc yaml.Node
	if err := yaml.Unmarshal([]byte(text), &doc); err != nil {
		return "", fmt.Errorf("%s: %w", configFile, err)
	}
	if len(doc.Content) == 0 || doc.Content[0].Kind != yaml.MappingNode {
		return text, nil
	}
	root := doc.Content[0]
	for i := 0; i+1 < len(root.Content); i += 2 {
		key, seq := root.Content[i], root.Content[i+1]
		if key.Value != "exclude_paths" || seq.Kind != yaml.SequenceNode {
			continue
		}
		var items []*yaml.Node
		changed, seen := false, map[string]bool{}
		for _, it := range seq.Content {
			v := it.Value
			switch v {
			case legacyConfigFile:
				v, changed = configFile, true
			case legacyManifest:
				changed = true
				continue
			}
			if seen[v] {
				changed = true
				continue
			}
			seen[v] = true
			n := *it
			n.Value = v
			items = append(items, &n)
		}
		if !seen[".gitattributes"] {
			n := yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: ".gitattributes"}
			if len(items) > 0 {
				n.Style = items[len(items)-1].Style
			}
			items = append(items, &n)
			changed = true
		}
		if !changed {
			return text, nil
		}
		next := len(text)
		lines := strings.SplitAfter(text, "\n")
		if i+2 < len(root.Content) {
			next = offsetOfLine(lines, root.Content[i+2].Line)
		}
		start := offsetOfLine(lines, key.Line)
		newSeq := *seq
		newSeq.Content = items
		out, err := yaml.Marshal(&yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{key, &newSeq}})
		if err != nil {
			return "", err
		}
		// Keep the blank and comment lines that end the old span: they belong
		// to the key that follows.
		span := strings.SplitAfter(text[start:next], "\n")
		cut := len(span)
		for cut > 1 {
			l := strings.TrimSpace(span[cut-1])
			if l != "" && !strings.HasPrefix(l, "#") {
				break
			}
			cut--
		}
		keep := strings.Join(span[cut:], "")
		return text[:start] + string(out) + keep + text[next:], nil
	}
	return text, nil
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"offsetOfLine"},"kwargs":"sha256:501cdb2bc1e6bfa01055bde71ce5972e4e62cdb123981201b03257ae0abac708","body":"sha256:73b6f4ee94404ecc66fd4d3c4a1745cd0349e4fb4a096072200ff4555fb4d479"}
// offsetOfLine is the byte offset where 1-based line n starts.
func offsetOfLine(lines []string, n int) int {
	off := 0
	for i := 0; i < n-1 && i < len(lines); i++ {
		off += len(lines[i])
	}
	return off
}

// /sedum:setFunction
// sedum:setFunction {"v":2,"tier":"owned","record":"prov-2026-b89af444","gen":"sha256:ead87ea88eac0f9b9f7dc9d559ecf4a821e98b5b4efff40d12e933cf9d469b28","pkg":"sha256:6589611adfb6618a7920e43f951a79901428192ea95c1084fde9f222aecd9537","id":{"file":"cmd/provenance/migrate_run.go","name":"stripComments"},"kwargs":"sha256:1ad32a8c3e72d43be3aa0415f35201bdcc194e935d7842b89768fda85acb407c","body":"sha256:529d2ed76dff5b1e7a59fe9406878b19cd4e910af94fd4ab06c7718e103d3dcb"}
func stripComments(s string) string {
	var b strings.Builder
	for _, l := range strings.Split(s, "\n") {
		if !strings.HasPrefix(strings.TrimSpace(l), "#") {
			b.WriteString(l)
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// /sedum:setFunction
