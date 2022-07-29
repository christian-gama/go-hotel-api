package errors

import "fmt"

// numbers represents a generic number type.
type numbers interface {
	int | float64 | float32 | uint | uint8 | uint16 | uint32 | uint64
}

// NonZero formats an error message for a non-zero field.
func NonZero(field string) error {
	return fmt.Errorf("%s must be greater than zero", field)
}

// NonNil formats an error message for a non-nil field.
func NonNil(field string) error {
	return fmt.Errorf("%s cannot be nil", field)
}

// NonEmpty formats an error message for a non-empty field.
func NonEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}

// NonNegative formats an error message for a non-negative field.
func NonNegative(field string) error {
	return fmt.Errorf("%s cannot be negative", field)
}

// MaxLength formats an error message based on the maximum length of a field.
func MaxLength(field string, max int) error {
	return fmt.Errorf("%s length cannot be greater than %d", field, max)
}

// MinLength formats an error message based on the min length of a field.
func MinLength(field string, min int) error {
	return fmt.Errorf("%s length cannot be less than %d", field, min)
}

// Max formats an error message based on the maximum value of a field.
func Max[T numbers](field string, max T) error {
	return fmt.Errorf("%s cannot be greater than %v", field, max)
}

// Min formats an error message based on the minimum value of a field.
func Min[T numbers](field string, min T) error {
	return fmt.Errorf("%s cannot be less than %v", field, min)
}

// Equal formats an error message for a field that must be equal to a given value.
func Equal(field string, equalTo any) error {
	return fmt.Errorf("%s must be equal to %v", field, equalTo)
}

// Different formats an error message for a field that must be different from a given value.
func Different(field string, differentFrom any) error {
	return fmt.Errorf("%s must be different from %v", field, differentFrom)
}

// NonDateBefore formats an error message to specify a field must happen before another field.
func NonDateBefore(field string, beforeField string) error {
	return fmt.Errorf("%s cannot be before %s", field, beforeField)
}

// NonDateAfter formats an error message to specify a field must happen after another field.
func NonDateAfter(field string, afterField string) error {
	return fmt.Errorf("%s cannot be after %s", field, afterField)
}

// MustBeMadeAfter formats an error message to specify an action must be made after a certain time.
func MustBeMadeAfter[T numbers](field string, mustWait T, measureTime string, waitFor string) error {
	return fmt.Errorf("%s must be made %v %s after %s", field, mustWait, measureTime, waitFor)
}
