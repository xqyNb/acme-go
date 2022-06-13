package frame_data

// PanicData 异常
type PanicData struct {
	ReqData RequestData `json:"reqData"`
	Detail  string      `json:"detail"`
}
