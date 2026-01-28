package web

import (
	"back/service/competence"
	"back/service/cv"
	"back/service/projet"
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.SetupTestData()

	competence.Setup(ramStore)
	projet.Setup(ramStore)
	cv.Setup(ramStore)

	m.Run()
}
