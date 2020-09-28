package functions

import (
	"encoding/json"
	"net/http"
)

// Payload is the expected incoming payload for our function.
type Payload struct {
	Name string
}

// IncomingHTTP takes the incoming payload from the function host
// decodes it and sends back the value from the name parameter
// in the body.
func IncomingHTTP(w http.ResponseWriter, r *http.Request) {
	var req HTTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload Payload
	if err := json.Unmarshal([]byte(req.Data.Req.Body), &payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := HTTPResponse{}
	response.Outputs.Res.Body = "Hello " + payload.Name

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
