package errx

type causer interface {
	Cause() error
}

// Cause returns the cause of the provided error.
func Cause(err error) error {
	var diag error
	if err, ok := err.(causer); ok {
		diag = err.Cause()
	}
	if diag != nil {
		return diag
	}
	return err
}
