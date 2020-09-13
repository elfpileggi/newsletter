package usecase_test

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestStore(t *testing.T) {
	g := Goblin(t)
	g.Describe("Message Store", func() {
		g.It("Should throw if has no arguments", func() {
		})
	})
}
