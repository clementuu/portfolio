package cv

import (
	"testing"
)

func TestGetCV(t *testing.T) {
	cv := GetCV()
	if cv.Formations == nil || cv.Experiences == nil {
		t.Errorf("GetCV() returned nil; expected a slice of formations and experiences")
	}

	if len(cv.Formations) != 3 {
		t.Error("GetCV() should return 3 formations.")
	}

	if cv.Formations[0].Intitule != "Master" {
		t.Errorf("Expected first formation name to be 'Master', got '%s'", cv.Formations[0].Intitule)
	}

	if cv.Formations[1].Periode != "2019 - 2022" {
		t.Errorf("Expected first formation period to be '2019 - 2022', got '%s'", cv.Formations[1].Periode)
	}

	if cv.Formations[2].Etablissement != "Bonaparte" {
		t.Errorf("Expected first formation etablissement to be 'Bonaparte', got '%s'", cv.Formations[2].Etablissement)
	}

	if len(cv.Experiences) != 2 {
		t.Error("GetCV() should return 2 experiences.")
	}

	if cv.Experiences[0].Intitule != "Pizzaiolo" {
		t.Errorf("Expected first experience name to be 'Pizzaiolo', got '%s'", cv.Experiences[0].Intitule)
	}

	if len(cv.Experiences[1].Taches) != 3 {
		t.Errorf("Expected second experience to have 3 tasks, got %d", len(cv.Experiences[1].Taches))
	}

	if cv.Experiences[1].Taches[0] != "Fumer un max de clopes" {
		t.Errorf("Expected first experience task to be 'Fumer un max de clopes', got '%s'", cv.Experiences[1].Taches[0])
	}
}
