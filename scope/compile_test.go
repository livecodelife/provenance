package scope

import (
	"regexp"
	"sync"
	"testing"
)

func TestCompileCachesAndMatchesRegexp(t *testing.T) {
	for _, expr := range []string{`^src/.*\.go$`, `a|b`, `(`, `[`, `^$`, `\p{Foo}`} {
		want, wantErr := regexp.Compile(expr)
		got, err := Compile(expr)
		if (err == nil) != (wantErr == nil) || (err != nil && err.Error() != wantErr.Error()) {
			t.Errorf("%q: error %v, want %v", expr, err, wantErr)
		}
		if err == nil && got.String() != want.String() {
			t.Errorf("%q: compiled %q", expr, got)
		}
		again, err2 := Compile(expr)
		if again != got || err2 != err {
			t.Errorf("%q: second call returned a different result", expr)
		}
	}
}

func TestCompileConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, p := range []string{"src/**", "src/*.go", "re:^x", "a?b"} {
				if _, err := Match("src/a.go", p); err != nil {
					t.Error(err)
				}
				Overlap(p, "src/**/*.go")
			}
		}()
	}
	wg.Wait()
}

// Match must give the same answers as compiling the pattern fresh each time.
func TestMatchUnchangedByCache(t *testing.T) {
	cases := []struct {
		path, pattern string
		want          bool
	}{
		{"src/a.go", "src/**", true},
		{"src/a/b.go", "src/*.go", false},
		{"src/a.go", "src/?.go", true},
		{"src/a.go", "re:^src/", true},
		{"src/a.go", "re:^lib/", false},
		{"src/a.go", "src/a.go", true},
	}
	for round := 0; round < 2; round++ {
		for _, c := range cases {
			if got, err := Match(c.path, c.pattern); err != nil || got != c.want {
				t.Errorf("round %d: Match(%q, %q) = %v, %v; want %v", round, c.path, c.pattern, got, err, c.want)
			}
		}
	}
	if _, err := Match("a", "re:("); err == nil {
		t.Error("invalid regex accepted")
	}
	if _, err := Match("a", "re:("); err == nil {
		t.Error("invalid regex accepted on the second call")
	}
}
