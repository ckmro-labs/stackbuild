package errors

var (
	// ErrInvalidToken err invalid token
	ErrInvalidToken = New("无效的登录令牌")
	// ErrUnauthorized not authorized.
	ErrUnauthorized = New("未认证的身份")
	// ErrForbidden is forbidden.
	ErrForbidden = New("无权访问")
	// ErrNotFound is not found.
	ErrNotFound = New("资源未发现")
)

// Error json-encoded API error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
