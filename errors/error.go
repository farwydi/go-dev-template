package errors

type Err struct {
    msg string
}

func New(msg string) error {
    return Err{msg}
}

func Wrap(err error, msg string) error {
    return Err{msg + ": " + err.Error()}
}

func (e Err) Error() string {
    return e.msg
}
