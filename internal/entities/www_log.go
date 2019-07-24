package entities

// WWWLog BadgerHole log format
type WWWLog interface {
	ToJSON() []byte
	GetRequestID() string
}
