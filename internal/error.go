package internal

func Error(err error) string {
	return BuildTemplate(htmlErrorTemplate, map[string]interface{}{
		"Error": err.Error(),
	})
}
