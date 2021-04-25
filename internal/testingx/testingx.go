package testingx

import (
	"io"
	"log"
	"testing"

	"github.com/franela/goblin"
)

func Init(t *testing.T) *goblin.G {
	log.SetOutput(io.Discard)
	return goblin.Goblin(t)
}
