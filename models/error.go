package models

type CustomError struct {
	Message string
	Code    uint8
}

func (e *CustomError) GetName() string {
	Names := []string{"Unknown", "BuiltIn", "SQL", "InvalidArg", "Serialize"}
	return Names[e.Code]
}

func (e *CustomError) Error() string {
	return e.GetName() + ":" + e.Message
}

func UnknownError() *CustomError {
	return &CustomError{"Unknown", 0}
}

func BuiltInError(err error) *CustomError {
	return &CustomError{err.Error(), 1}
}

func SQLError(msg string) *CustomError {
	return &CustomError{msg, 2}
}

func InvalidArgError(msg string) *CustomError {
	return &CustomError{msg, 3}
}

func SerializeError() *CustomError {
	return &CustomError{"Serialize Error", 4}
}
