package goapipkg

func Response(body interface{}) ImplResponse {
	return ImplResponse{Body: body}
}
