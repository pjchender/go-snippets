package configs

type Database struct {
	DSN     string `default:"sslmode=disable host=localhost port=5432 dbname=jubo_space"`
	TestDSN string `default:"sslmode=disable host=localhost port=5432 dbname=jubo_space_test"`
}
