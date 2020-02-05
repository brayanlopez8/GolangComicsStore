package tools

//ReadConfig read key config by name
func ReadConfig(key string) string {

	// load .env file
	err := godotenv.Load("configs/.envdev")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
