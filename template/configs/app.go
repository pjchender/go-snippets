package configs

type App struct {
	Name            string `default:"jubo-space"`
	Mode            string `default:"development"`
	DefaultPageSize int    `default:"10"`
	MaxPageSize     int    `default:"100"`
}
