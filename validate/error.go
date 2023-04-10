package validate

var _ error = (*ValidateErr)(nil)

type ValidateErr struct {
	Message string
}

func NewValidateError(message string) ValidateErr {
	return ValidateErr{
		Message: message,
	}
}

func (e ValidateErr) Error() string {
	return e.Message
}
