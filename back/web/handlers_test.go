package web

import (
	"back/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCompetenceByID(t *testing.T) {
	ts := []struct {
		name       string
		id         string
		wantStatus int
		wantBody   *service.Competence
	}{
		{name: "OK", id: "1", wantStatus: http.StatusOK, wantBody: &service.Competence{ID: 1, Name: "Go", Rating: 4}},
		{name: "Bad request", id: "abc", wantStatus: http.StatusBadRequest, wantBody: nil},
		{name: "Not Found", id: "999", wantStatus: http.StatusInternalServerError, wantBody: nil},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(`/competences/%s`, tc.id), nil)
			req.SetPathValue("id", tc.id)

			rr := httptest.NewRecorder()

			GetCompetenceByID(rr, req)

			if rr.Code != tc.wantStatus {
				t.Fatalf("expected status %d, got %d", tc.wantStatus, rr.Code)
			}

			if tc.wantBody != nil {
				var got service.Competence
				if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
					t.Fatalf("could not decode response body: %v", err)
				}
				if got.ID != tc.wantBody.ID || got.Name != tc.wantBody.Name || got.Rating != tc.wantBody.Rating {
					t.Errorf("got body %+v, want %+v", got, *tc.wantBody)
				}
			}
		})
	}
}

func TestGetAllCompetences(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/competences", nil)
	rr := httptest.NewRecorder()

	GetAllCompetences(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var got []service.Competence
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	// We have 2 competences in our mock data
	if len(got) != 2 {
		t.Errorf("expected 2 competences, got %d", len(got))
	}
}

func TestGetAllProjets(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/projets", nil)
	rr := httptest.NewRecorder()

	GetAllProjets(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var got []service.Projet
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("expected 2 projets, got %d", len(got))
	}
}

func TestGetProjetByID(t *testing.T) {
	ts := []struct {
		name       string
		id         string
		wantStatus int
		wantBody   *service.Projet
	}{
		{name: "OK", id: "1", wantStatus: http.StatusOK, wantBody: &service.Projet{ID: 1, Name: "Maquette 2cv", Sujet: "Monter une maquette de 2cv avec le padre"}},
		{name: "Bad request", id: "abc", wantStatus: http.StatusBadRequest, wantBody: nil},
		{name: "Not Found", id: "999", wantStatus: http.StatusInternalServerError, wantBody: nil},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(`/projets/%s`, tc.id), nil)
			req.SetPathValue("id", tc.id)

			rr := httptest.NewRecorder()

			GetProjetByID(rr, req)

			if rr.Code != tc.wantStatus {
				t.Fatalf("expected status %d, got %d", tc.wantStatus, rr.Code)
			}

			if tc.wantBody != nil {
				var got service.Projet
				if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
					t.Fatalf("could not decode response body: %v", err)
				}
				if got.ID != tc.wantBody.ID || got.Name != tc.wantBody.Name || got.Sujet != tc.wantBody.Sujet {
					t.Errorf("got body %+v, want %+v", got, *tc.wantBody)
				}
			}
		})
	}
}
