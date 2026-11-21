package lint

import (
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/livecodelife/provenance/record"
)

// memoStore answers Versions and Snapshot from memory after the first read.
// One lint run reads each snapshot several times (the version rules, then the
// manifest check), and a snapshot is a YAML parse plus a content hash. The
// store is not expected to change during a run, which holds for lint: it
// writes nothing. Errors are remembered too, so a failing read fails the same
// way every time it is asked.
type memoStore struct {
	record.Store
	mu    sync.Mutex
	vers  map[string]*versionsEntry
	snaps map[snapKey]*snapEntry
}

type snapKey struct {
	id      string
	version int
}

type versionsEntry struct {
	once sync.Once
	vs   []int
	err  error
}

type snapEntry struct {
	once sync.Once
	s    *record.Snapshot
	err  error
}

func newMemoStore(s record.Store) *memoStore {
	return &memoStore{Store: s, vers: map[string]*versionsEntry{}, snaps: map[snapKey]*snapEntry{}}
}

func (m *memoStore) Versions(id string) ([]int, error) {
	m.mu.Lock()
	e := m.vers[id]
	if e == nil {
		e = &versionsEntry{}
		m.vers[id] = e
	}
	m.mu.Unlock()
	e.once.Do(func() { e.vs, e.err = m.Store.Versions(id) })
	if e.err != nil {
		return nil, e.err
	}
	return append([]int(nil), e.vs...), nil
}

func (m *memoStore) Snapshot(id string, version int) (*record.Snapshot, error) {
	k := snapKey{id, version}
	m.mu.Lock()
	e := m.snaps[k]
	if e == nil {
		e = &snapEntry{}
		m.snaps[k] = e
	}
	m.mu.Unlock()
	e.once.Do(func() { e.s, e.err = m.Store.Snapshot(id, version) })
	return e.s, e.err
}

// parallelism is how many files are read at once. Only the files store is
// read in parallel: any other store may not allow concurrent calls.
func parallelism(s record.Store) int {
	if _, ok := s.(*record.FilesStore); !ok {
		return 1
	}
	n := runtime.GOMAXPROCS(0)
	if n > 8 {
		n = 8
	}
	if n < 1 {
		n = 1
	}
	return n
}

// forEach calls fn(i) for every i in [0, n) on up to workers goroutines and
// returns once all are done. fn must only write to its own slot.
func forEach(n, workers int, fn func(i int)) {
	if workers <= 1 || n <= 1 {
		for i := 0; i < n; i++ {
			fn(i)
		}
		return
	}
	if workers > n {
		workers = n
	}
	var wg sync.WaitGroup
	next := make(chan int)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range next {
				fn(i)
			}
		}()
	}
	for i := 0; i < n; i++ {
		next <- i
	}
	close(next)
	wg.Wait()
}

// fileList is every regular file under the lint root by root-relative path,
// in the order filepath.Walk visits them. Scope regexes are matched against
// it; the tree is walked once per run instead of once per pattern.
type fileList struct {
	once  sync.Once
	files []string
}

func (l *linter) fileNames() []string {
	l.walked.once.Do(func() {
		root := l.opts.Root
		if root == "" {
			root = "."
		}
		_ = filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			rel := p
			if root != "." {
				if r, e := filepath.Rel(root, p); e == nil {
					rel = r
				}
			}
			l.walked.files = append(l.walked.files, rel)
			return nil
		})
	})
	return l.walked.files
}
