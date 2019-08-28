package errorscode

// ErrorStruct - Error Default Struct
type ErrorStruct struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetError - Returns a Default Error Struct with code and message
func GetError(code string) (result ErrorStruct) {
	result.Code = code
	result.Message = Errors[code]
	return
}

// Errors - Store errors code and message
var Errors = map[string]string{
	// 10xx - General Server or Network issues
	"1000": "An unknown error occured while processing the request.",
	"1001": "Internal error; unable to process your request. Please try again.",

	// 11xx - Request issues
	"1100": "An unknown parameter was sent.",
	"1101": "Invalid country.",
	"1102": "Invalid 2FA code.",
	"1103": "Invalid e-mail auth code.",
	"1104": "CPF already exists.",
	"1105": "Username already exists.",
}
