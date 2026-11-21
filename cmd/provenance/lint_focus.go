package main

import "github.com/livecodelife/provenance/lint"

// lintRepoFocus is lintRepo limited to the issues about the records in focus
// (all of them when focus is nil); see lint.Options.Focus.
func lintRepoFocus(r *repo, enforcement string, focus map[string]bool) (*lint.Result, error) {
	if enforcement == "" {
		enforcement = r.Config.Enforcement
	}
	if enforcement == "" {
		enforcement = "none"
	}
	shared, _ := loadShared(r)
	return lint.Lint(r.Store, lint.Options{
		Shared:       shared,
		Root:         r.Root,
		Enforcement:  enforcement,
		ExcludePaths: r.Config.ExcludePaths,
		SharedRepos:  sharedNames(r),
		ManifestPath: manifestPath,
		Focus:        focus,
	})
}
