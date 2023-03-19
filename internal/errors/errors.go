package errors

import "net/http"

var ErrorBoardNotFound = NewNotFound("Board bulunamadı.")

func NewNotFound(detail string) *ApplicationError {
	return New(http.StatusNotFound, "Kayıt bulunamadı.", detail)
}

func New(status int, title, detail string) *ApplicationError {
	return &ApplicationError{
		Title:   title,
		Status:  status,
		Message: detail,
	}
}

type ApplicationError struct {
	Title         string                   `json:"title,omitempty"`
	Status        int                      `json:"status,omitempty"`
	Detail        string                   `json:"detail,omitempty"`
	Host          string                   `json:"host,omitempty"`
	RequestUri    string                   `json:"requestUri,omitempty"`
	RequestMethod string                   `json:"requestMethod,omitempty"`
	Instant       string                   `json:"instant,omitempty"`
	ErrorDetails  []ApplicationErrorDetail `json:"errorDetails,omitempty"`
	Cause         string                   `json:"cause,omitempty"`
	CorrelationID string                   `json:"correlationId,omitempty"`
	Message       string                   `json:"message,omitempty"`
}

type ApplicationErrorDetail struct {
	ErrorCode    string       `json:"errorCode"`
	ErrorMessage string       `json:"errorMessage"`
	FormErrors   []*FormError `json:"formErrors,omitempty"`
}

type FormError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ApplicationError) Error() string {
	if len(e.Cause) > 0 {
		return e.Cause
	}
	if len(e.Detail) > 0 {
		return e.Detail
	}
	return e.Title
}
