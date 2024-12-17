package backend

type Config struct {
	token          string
	amountCurrency int
}

func NewConfig() *Config {
	return &Config{
		token:          "c2d4b21b-6a83-4098-87be-24d4e4919fd0",
		amountCurrency: 20,
	}
}
