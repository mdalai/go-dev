package goapipkg

func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}
