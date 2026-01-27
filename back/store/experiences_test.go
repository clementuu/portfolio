package store

import (
	"reflect"
	"testing"
)

func TestGetExperiences(t *testing.T) {
	s := NewRAMStore()
	
	experiences := s.GetExperiences()

	if !reflect.DeepEqual(experiences, experiencesList) {
		t.Error("GetExperiences() should return the initial list of experiences")
	}
}
