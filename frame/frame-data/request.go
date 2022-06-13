package frame_data

// KeyRequestData 请求数据key
const (
	KeyRequestData = "KEY_REQUEST_DATA"
	KeyRequestKey  = "KEY_REQUEST_KEY"
)

// RequestData 请求信息
type RequestData struct {
	RequestKey string
	ClientIp   string
	Time       string
	Method     string
	Path       string
	Protocal   string
	UserAgent  string
	Uri        string
}

// ResponseData 响应数据
type ResponseData struct {
	RequestKey string
	Time       string
	Method     string
	StatusCode int
	Latency    string
	ApiData    string
}
