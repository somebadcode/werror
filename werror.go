package werror

type Error struct {
	err     error
	nextErr error
}

func New(err error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		err:     err,
		nextErr: nil,
	}
}

func Wrap(err error, wrap error) *Error {
	if err == nil && wrap == nil {
		return nil
	} else if err == nil && wrap != nil {
		return &Error{
			err:     wrap,
			nextErr: nil,
		}
	}
	return &Error{
		err:     err,
		nextErr: wrap,
	}
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Err() error {
	return e.err
}

func (e Error) Unwrap() error {
	return e.nextErr
}

func (e Error) Is(err error) bool {
	return e.err == err
}

func (e Error) Wrap(err error) *Error {
	return &Error{
		err:     err,
		nextErr: e,
	}
}
