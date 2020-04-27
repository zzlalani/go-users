package classes

func PrepareValidationErrorsMessages(messages []string) map[string][]string {
	return map[string][]string {
		"validation": messages,
	}
}