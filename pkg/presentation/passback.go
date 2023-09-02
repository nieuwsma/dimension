package presentation

type APIPayload struct {
	StatusCode int
	Obj        interface{}
	Error      Problem7807
	IsError    bool
}

func BuildErrorPassback(StatusCode int, Err error) (passback APIPayload) {
	passback.StatusCode = StatusCode
	passback.Obj = nil
	passback.Error = GetFormattedErrorMessage(Err, StatusCode)
	passback.IsError = true
	return passback
}

func BuildSuccessPassback(StatusCode int, Obj interface{}) (passback APIPayload) {
	passback.StatusCode = StatusCode
	passback.Obj = Obj
	passback.IsError = false
	return passback
}
