package store

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetProjets(t *testing.T) {
	s := NewRAMStore()

	projets := s.GetProjets()

	if !reflect.DeepEqual(projets, projetsList) {
		t.Error("GetProjets() should return the initial list of projets")
	}
}

func TestGetProjetsNames(t *testing.T) {
	s := NewRAMStore()

	miniProjets := s.GetProjetsNames()

	if miniProjets == nil {
		t.Error("GetProjetsNames() should not return nil")
	}

	if len(miniProjets) != len(projetsList) {
		t.Errorf("GetProjetsNames() returned %d mini projets, expected %d", len(miniProjets), len(projetsList))
	}

	for i, mp := range miniProjets {
		if mp.ID != projetsList[i].ID {
			t.Errorf("MiniProjet ID at index %d is %d, expected %d", i, mp.ID, projetsList[i].ID)
		}
		if mp.Name != projetsList[i].Name {
			t.Errorf("MiniProjet Name at index %d is %s, expected %s", i, mp.Name, projetsList[i].Name)
		}
	}
}

func TestGetProjet(t *testing.T) {
	s := NewRAMStore()

	tests := []struct {
		name         string
		projetID     int
		expectError  bool
		expectedName string
	}{
		{
			name:         "Retrieve existing projet (ID 1)",
			projetID:     1,
			expectError:  false,
			expectedName: "Escarcelle",
		},
		{
			name:         "Retrieve existing projet (ID 3)",
			projetID:     3,
			expectError:  false,
			expectedName: "Mission TousAntiGaspi",
		},
		{
			name:         "Retrieve non-existent projet (ID 999)",
			projetID:     999,
			expectError:  true,
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			projet, err := s.GetProjet(tt.projetID)

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
			} else {
				expectedErrorMessage := fmt.Sprintf(ErrNoProjet, tt.projetID)
				if err.Error() != expectedErrorMessage {
					t.Errorf("GetProjet() error message = %s, want %s", err.Error(), expectedErrorMessage)
				}
			}
		})
	}
}
