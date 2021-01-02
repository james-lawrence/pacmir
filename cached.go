package pacmir

import (
	"log"
	"sync"
	"time"

	paconf "github.com/Morganamilo/go-pacmanconf"
	"github.com/pkg/errors"
)

// NewCachedConfig returns a caching wrapper around
// the pacman configuration.
func NewCachedConfig(path string) *CachedConfig {
	return &CachedConfig{
		path:       path,
		m:          &sync.RWMutex{},
		stalestamp: time.Now(),
	}
}

// CachedConfig a caching wrapper around a pacman configuration.
type CachedConfig struct {
	path       string
	m          *sync.RWMutex
	stalestamp time.Time
	config     *paconf.Config
}

// Current returns the current pacman configuration
func (t *CachedConfig) Current() *paconf.Config {
	t.refresh()
	t.m.RLock()
	defer t.m.RUnlock()

	return t.config
}

// Mirrors of the given repository.
func (t *CachedConfig) Mirrors(repository string) []string {
	t.refresh()
	t.m.RLock()
	defer t.m.RUnlock()

	if t.config == nil {
		return []string{}
	}

	if repo := t.config.Repository(repository); repo != nil {
		return repo.Servers
	}

	return []string{}
}

func (t *CachedConfig) refresh() {
	var (
		err    error
		config *paconf.Config
	)

	t.m.RLock()
	required := t.config == nil || t.stalestamp.After(time.Now())
	t.m.RUnlock()

	if !required {
		return
	}

	// refresh config
	t.m.Lock()
	defer t.m.Unlock()

	// check again
	if !(t.config == nil || t.stalestamp.After(time.Now())) {
		return
	}

	if config, _, err = paconf.ParseFile(t.path); err != nil {
		log.Println(errors.Wrap(err, "failed to parse config"))
		t.reset(nil)
		return
	}

	t.reset(config)
}

func (t *CachedConfig) reset(c *paconf.Config) {
	t.config = c
	t.stalestamp = time.Now().Add(time.Minute)
}
