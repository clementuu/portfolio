package store

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCompetences(t *testing.T) {
	s := NewRAMStore()

	competences := s.GetCompetences()

	if !reflect.DeepEqual(competences, competencesList) {
		t.Error("GetCompetences() should return the initial list of competences")
	}
}

func TestGetCompetence(t *testing.T) {
	s := NewRAMStore()

	tests := []struct {
		name         string
		competenceID int
		expectError  bool
		expectedName string
	}{
		{
			name:         "Retrieve existing competence (ID 1)",
			competenceID: 1,
			expectError:  false,
			expectedName: "Go",
		},
		{
			name:         "Retrieve existing competence (ID 10)",
			competenceID: 10,
			expectError:  false,
			expectedName: "Rebol",
		},
		{
			name:         "Retrieve non-existent competence (ID 999)",
			competenceID: 999,
			expectError:  true,
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp, err := s.GetCompetence(tt.competenceID)

			if (err != nil) != tt.expectError {
				t.Errorf("GetCompetence() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				if comp.ID != tt.competenceID {
					t.Errorf("GetCompetence() got competence with ID %d, want %d", comp.ID, tt.competenceID)
				}
				if comp.Name != tt.expectedName {
					t.Errorf("GetCompetence() got competence name '%s', want '%s'", comp.Name, tt.expectedName)
				}
			} else {
				expectedErrorMessage := fmt.Sprintf(ErrNoCompetence, tt.competenceID)
				if err.Error() != expectedErrorMessage {
					t.Errorf("GetCompetence() error message = %s, want %s", err.Error(), expectedErrorMessage)
				}
			}
		})
	}
}
