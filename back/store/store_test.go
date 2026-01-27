package store

import (
	"testing"
)

func TestNewRAMStore(t *testing.T) {
	s := NewRAMStore()

	if s == nil {
		t.Error("NewRAMStore() should not return nil")
	}

	if s.Competences == nil {
		t.Error("NewRAMStore() should initialize Competences")
	}

	if s.Projets == nil {
		t.Error("NewRAMStore() should initialize Projets")
	}

	if s.Formations == nil {
		t.Error("NewRAMStore() should initialize Formations")
	}

	if s.Experiences == nil {
		t.Error("NewRAMStore() should initialize Experiences")
	}
}
