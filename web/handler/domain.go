package handler

type message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StatusOK(msg string, data interface{}) message {
	return message{
		Code:    0,
		Message: msg,
		Data:    data,
	}
}

func StatusBad(msg string, data interface{}) message {
	return message{
		Code:    -1,
		Message: msg,
		Data:    data,
	}
}
