package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ErrEncod = "erreur d'encodage json"

func encodeJSON(w http.ResponseWriter, a any) error {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(a)
	if err != nil {
		return fmt.Errorf("%v : %v", ErrEncod, err)
	}

	return nil
}
