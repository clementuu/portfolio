package projet

import (
	"strings"
	"testing"
)

func TestGetAllProjets(t *testing.T) {
	projets := GetAll()

	if projets == nil {
		t.Errorf("GetAllProjets() returned nil; expected a slice of projets")
	}

	if len(projets) != 2 {
		t.Error("GetAllProjets() should return 2 projets.")
	}

	if projets[0].Name != "Maquette 2cv" {
		t.Errorf("Expected first projet name to be 'Maquette de 2cv', got '%s'", projets[0].Name)
	}

	if !strings.Contains(projets[1].Resume, "Laisser pousser les cheveux") {
		t.Errorf("Expected second projet resume to contain 'Laisser pousser les cheveux', got '%s'", projets[1].Resume)
	}
}

func TestGetProjetsNames(t *testing.T) {
	miniProjets := GetNames()

	if miniProjets == nil {
		t.Errorf("GetProjetsNames() returned nil; expected a slice of mini projets")
	}

	if len(miniProjets) != 2 {
		t.Error("GetProjetsNames() should return 2 mini projets.")
	}

	if miniProjets[0].Name != "Maquette 2cv" {
		t.Errorf("Expected first mini projet name to be 'Maquette de 2cv', got '%s'", miniProjets[0].Name)
	}

	if miniProjets[1].ID != 2 {
		t.Errorf("Expected second mini projet ID to be 2, got %d", miniProjets[1].ID)
	}
}

func TestGetProjet(t *testing.T) {
	knownID := 1
	nonExistentID := 999

	tests := []struct {
		name         string
		projetID     int
		expectError  bool
		expectedName string
	}{
		{
			name:         "Retrieve existing projet (ID 1)",
			projetID:     knownID,
			expectError:  false,
			expectedName: "Maquette 2cv",
		},
		{
			name:         "Retrieve non-existent projet (ID 999)",
			projetID:     nonExistentID,
			expectError:  true,
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projet, err := Get(tt.projetID)

			if (err != nil) != tt.expectError {
				t.Errorf("GetProjet() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				if projet.ID != tt.projetID {
					t.Errorf("GetProjet() got projet with ID %d, want %d", projet.ID, tt.projetID)
				}
				if projet.Name != tt.expectedName {
					t.Errorf("GetProjet() got projet name '%s', want '%s'", projet.Name, tt.expectedName)
				}
			}
		})
	}
}
