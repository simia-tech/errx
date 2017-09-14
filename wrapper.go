package errx

type wrapper interface {
	Message() string
	Underlying() error
}
