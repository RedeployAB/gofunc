package functions

// HTTPRequest is the expected incoming payload from a HTTP Trigger
// from the function host.
type HTTPRequest struct {
	Data struct {
		Req struct {
			Body string
		} `json:"req"`
	}
	Metadata map[string]interface{}
}

// HTTPResponse is the expected outgoing payload to a HTTP Trigger
// to function host.
type HTTPResponse struct {
	Outputs struct {
		Res struct {
			Body string `json:"body"`
		} `json:"res"`
	}
	Logs        []string
	ReturnValue interface{}
}
