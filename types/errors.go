package types

const (
	ConvertRequestError = "Request参数转化失败 "
)

type EmptyUserError struct{}

func (e EmptyUserError) Error() string {
	return "user empty"
}

type EmptyNameError struct{}

func (e EmptyNameError) Error() string {
	return "name empty"
}

type EmptyMockContentError struct{}

func (e EmptyMockContentError) Error() string {
	return "mock content error"
}

type EmptyPathError struct{}

func (e EmptyPathError) Error() string {
	return "path empty"
}

type NotFoundMockError struct{}

func (e NotFoundMockError) Error() string {
	return "not found this mock api"
}
