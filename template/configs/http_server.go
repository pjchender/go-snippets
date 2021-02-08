package configs

type HTTPServer struct {
	ListenAddr string `default:""`
	Port       string `default:"3000"`
	SSL        struct {
		Enabled         *bool  `default:"false"`
		RedirectToHTTPS *bool  `default:"true"`
		ListenAddr      string `default:""`
		Port            int    `default:"443"`
		CertFile        string `default:"./certificate/fullchain.pem"`
		CertKey         string `default:"./certificate/privkey.pem"`
		LetsEncrypt     struct {
			Enabled   *bool  `default:"false"`
			AcceptTOS *bool  `default:"false"`
			Cache     string `default:"data/certs"`
			Hosts     []string
		}
	}
	ResponseHeaders map[string]string
}
