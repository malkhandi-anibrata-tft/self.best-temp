package utility

func CreateErrorResponse(err error) map[string]string {
	return map[string]string{
		"Error": err.Error(),
	}
}
