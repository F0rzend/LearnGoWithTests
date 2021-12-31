package main

import "fmt"

const (
	EnglishTemplate string = "Hello, %s!"
	RussianTemplate string = "Привет, %s!"

	WorldEnglish = "World"
	WorldRussian = "Мир"
)

func Hello(name, language string) string {
	if name == "" {
		switch language {
		case "Russian":
			name = WorldRussian
		default:
			name = WorldEnglish
		}
	}

	var result string

	switch language {
	case "Russian":
		result = fmt.Sprintf(RussianTemplate, name)
	default:
		result = fmt.Sprintf(EnglishTemplate, name)
	}

	return result
}

func main() {
}
