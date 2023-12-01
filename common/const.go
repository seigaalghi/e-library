package common

import "os"

func GetString(key string, def string) string {
	res := os.Getenv(key)
	if res == "" {
		res = def
	}

	return res
}

var (
	JwtKey             = GetString("JWT_KEY", "somekey")
	OpenLibraryBaseUrl = GetString("OPEN_LIBRARY_BASE_URL", "https://openlibrary.org")
	DBFile             = GetString("DB_FILE", "data.db")
)
