package controllerhelper

func ResponseFailed(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func ResponseOkNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func ResponseOkWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}
