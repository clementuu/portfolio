package store

import (
	"reflect"
	"testing"
)

func TestGetFormations(t *testing.T) {
	s := NewRAMStore()
	
	formations := s.GetFormations()

	if !reflect.DeepEqual(formations, formationsList) {
		t.Error("GetFormations() should return the initial list of formations")
	}
}
