package helper

func FailedResponse(message string) map[string]any {
	return map[string]any{
		"message": message,
		"status":  "failed",
	}
}

func SuccessResponse(message string) map[string]any {
	return map[string]any{
		"message": message,
		"status":  "success",
	}
}

func SuccessWithDataResponse(message string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}
