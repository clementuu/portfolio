package service_test

import (
	"back/service" // Assuming direct use of the store
	"testing"
)

func TestGetAllCompetences(t *testing.T) {
	competences := service.GetAllCompetences()

	if competences == nil {
		t.Errorf("GetAllCompetences() returned nil; expected a slice of competences")
	}

	if len(competences) != 2 {
		t.Error("GetAllCompetences() should return 2 competences.")
	}
	if competences[0].Name != "Go" {
		t.Errorf("Expected first competence name to be 'Go', got '%s'", competences[0].Name)
	}
}

func TestGetCompetence(t *testing.T) {
	tests := []struct {
		name         string
		competenceID int
		expectError  bool
		expectedName string // To verify if we get the correct competence
	}{
		{
			name:         "Retrieve existing competence (ID 1)",
			competenceID: 1,
			expectError:  false,
			expectedName: "Go", // Assuming this is set in store.NewRAMStore()
		},
		{
			name:         "Retrieve existing competence (ID 2)",
			competenceID: 2,
			expectError:  false,
			expectedName: "Docker", // Assuming this is set in store.NewRAMStore()
		},
		{
			name:         "Retrieve non-existent competence (ID 999)",
			competenceID: 999,
			expectError:  true,
			expectedName: "", // Not applicable for error case
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp, err := service.GetCompetence(tt.competenceID)

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
			}
		})
	}
}
