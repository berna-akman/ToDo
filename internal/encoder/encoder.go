package encoder

import (
	"encoding/json"
	"net/http"
)

func Encoder(w http.ResponseWriter, model interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model)
	return err
}
