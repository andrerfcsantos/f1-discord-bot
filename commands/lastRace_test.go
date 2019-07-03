package commands

import (
	"log"
	"testing"
)

func TestAbs(t *testing.T) {
	m, _ := LastRace()
	log.Printf("%v\n", m)
}
