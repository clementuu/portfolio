package cv

import (
	"testing"
)

func TestGetFormations(t *testing.T) {
	formations := GetFormations()

	if formations == nil {
		t.Errorf("GetFormations() returned nil; expected a slice of formations")
	}

	if len(formations) != 3 {
		t.Error("GetFormations() should return 3 formations.")
	}

	if formations[0].Intitule != "Master" {
		t.Errorf("Expected first formation name to be 'Master', got '%s'", formations[0].Intitule)
	}

	if formations[1].Periode != "2019 - 2022" {
		t.Errorf("Expected first formation period to be '2019 - 2022', got '%s'", formations[1].Periode)
	}

	if formations[2].Etablissement != "Bonaparte" {
		t.Errorf("Expected first formation etablissement to be 'Bonaparte', got '%s'", formations[2].Etablissement)
	}
}

func TestGetExperiences(t *testing.T) {
	experiences := GetExperiences()

	if experiences == nil {
		t.Errorf("GetExperiences() returned nil; expected a slice of experiences")
	}

	if len(experiences) != 2 {
		t.Error("GetExperiences() should return 2 experiences.")
	}

	if experiences[0].Intitule != "Pizzaiolo" {
		t.Errorf("Expected first experience name to be 'Pizzaiolo', got '%s'", experiences[0].Intitule)
	}

	if len(experiences[1].Taches) != 3 {
		t.Errorf("Expected second experience to have 3 tasks, got %d", len(experiences[1].Taches))
	}

	if experiences[1].Taches[0] != "Fumer un max de clopes" {
		t.Errorf("Expected first experience task to be 'Fumer un max de clopes', got '%s'", experiences[1].Taches[0])
	}
}
