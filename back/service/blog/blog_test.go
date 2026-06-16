package blog

import (
	"testing"
)

func TestGetArticles(t *testing.T) {
	articles := GetArticles()

	if articles == nil {
		t.Errorf("GetArticles() returned nil; expected a slice of articles")
	}

	if len(articles) != 1 {
		t.Errorf("GetArticles() should return 1 article, got %d", len(articles))
	}

	if articles[0].Titre != "Test Article" {
		t.Errorf("Expected article title to be 'Test Article', got '%s'", articles[0].Titre)
	}
}

func TestGetArticleByID(t *testing.T) {
	tests := []struct {
		name          string
		articleID     int
		expectError   bool
		expectedTitle string
	}{
		{
			name:          "Retrieve existing article (ID 1)",
			articleID:     1,
			expectError:   false,
			expectedTitle: "Test Article",
		},
		{
			name:          "Retrieve non-existent article (ID 999)",
			articleID:     999,
			expectError:   true,
			expectedTitle: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article, err := GetArticleByID(tt.articleID)

			if (err != nil) != tt.expectError {
				t.Errorf("GetArticleByID() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError {
				if article.ID != tt.articleID {
					t.Errorf("GetArticleByID() got article with ID %d, want %d", article.ID, tt.articleID)
				}
				if article.Titre != tt.expectedTitle {
					t.Errorf("GetArticleByID() got article title '%s', want '%s'", article.Titre, tt.expectedTitle)
				}
			}
		})
	}
}
