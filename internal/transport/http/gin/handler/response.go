package handler

type APIResponse struct {
	Success 	bool        `json:"success"`
	Message 	string      `json:"message"`
	Error		*APIError	`json:"error,omitempty"`
	Data    	any 		`json:"data,omitempty"`
}

type APIError struct {
	Code		string	`json:"code"`
	Message		string	`json:"message"`
	Details		any		`json:"details,omitempty"`
}

