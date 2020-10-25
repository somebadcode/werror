package werror

type Error struct {
	err     error
	nextErr error
}

func New(err error) Error {
	return Error{
		err:     err,
		nextErr: nil,
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

func (e Error) Wrap(err error) Error {
	return Error{
		err:     err,
		nextErr: e,
	}
}