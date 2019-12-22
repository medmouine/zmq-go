package subscriber

type Subscriber struct {
	ClientId []byte
	UserId   string
}

type Request struct {
	Id string
	RequestType string
	Message string
}

const (
	Accepted int = 202
)

const (
	Connection string = "CONNECT"
)

type Message struct {
	Code int
	Message string
}
