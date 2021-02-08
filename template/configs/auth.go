package configs

type Auth struct {
	JWT
	Password
}

type JWT struct {
	Issuer      string `default:"jubo-space"`
	ExpireHours int    `default:"4"`
	Secret      string `default:"jubo-space"`
}

type Password struct {
	Strength int    `default:"10"`
	Salt     string `default:"jubo-space"`
}
