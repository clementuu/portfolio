package cv

import (
	"back/store"
	"testing"
)

func TestMain(m *testing.M) {
	ramStore := store.SetupTestData()

	Setup(ramStore)

	m.Run()
}
