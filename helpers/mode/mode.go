package mode

type Mode string

var (
	Dev  Mode = "dev"
	Prod Mode = "prod"
	Test Mode = "test"
)

var mode = Dev

// Set mode with Set(mode.Dev)
func Set(newMode Mode) {
	mode = newMode
}

func Get() Mode {
	return mode
}
