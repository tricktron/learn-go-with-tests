package go_specs_greet //nolint:revive,stylecheck

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
