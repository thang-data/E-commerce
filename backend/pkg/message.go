package pkg

const (
	MessageTypeError   = "error"
	MessageTypeInfo    = "info"
	MessageTypeWarning = "warning"
)

type Message struct {
	Type    string `json:"type"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Title   string `json:"title"`
	Code    string `json:"code,omitempty"`
}

const (
	MsgErrInvalid                               = "Invalid structure"
	MsgErrPasswordMustContainAtLeast1Alphabet   = "The password must have the character"
	MsgErrPasswordMustContainAtLeast1Number     = "Password must have a number"
	MsgErrPasswordMustContainAtLeast8Characters = "Password must be more than 8 characters"
	MsgErrPasswordAndRepeatPasswordMustBeSame   = "The reset password must be the same as the previously set password"
	MsgErrAlreadyExists                         = "%sAlready exists"
	MsgErrWrong                                 = "%sIs wrong"
)

func NewError(message string, fields ...string) *Message {
	field := ""
	if len(fields) > 0 {
		field = fields[0]
	}

	msgErr := &Message{
		Type:    MessageTypeError,
		Field:   field,
		Message: message,
	}
	return msgErr
}
func (e *Message) Error() string {
	return e.Message
}
