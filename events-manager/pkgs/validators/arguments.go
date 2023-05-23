package validators

import "fmt"

// This function can be used to check if a certain condition is met
// and return an error if it is not.
// It checks if the `expression` is `false`.
// If it is `false`, it returns an error with the message `msg`.
// If the `expression` is `true`, it returns `nil`.
func CheckArgument(expression bool, msg string) error {
	if !expression {
		return fmt.Errorf(msg)
	}
	return nil
}
