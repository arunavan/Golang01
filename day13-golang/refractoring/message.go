package message

import "fmt"

func FormatMessage(name string) string {
	if len(name) == 0 {
		return "Welcome"
	} else {
		return fmt.Sprintf("Hi, %s", name)
	}
}
