package configs

type GRPC struct {
	Server
	Client
}

type Server struct {
	ListenAddr string `default:""`
	Port       string `default:"50051"`
	CertFile   string `default:"./certificate/fullchain.pem"`
	CertKey    string `default:"./certificate/privkey.pem"`
}

type Client struct {
	Port string `default:"50051"`
}
