package config

import "os"

type Env struct {
	LogPath    string
	DbDsn      string
	WebPort    string
	UploadPath string
	UploadDir  string
}

const newsApiLogPath = "NEWS_API_LOG_PATH"
const newsApiDbDsn = "NEWS_API_DB_DSN"
const newsApiWebPort = "NEWS_API_WEB_PORT"
const uploadPath = "NEWS_API_UPLOAD_PATH"
const uploadDir = "uploads"

// NewEnv Повертає об'єкт конфігурації, заповнений зі змінних оточення.
func NewEnv() Env {
	return Env{
		LogPath:    os.Getenv(newsApiLogPath),
		DbDsn:      os.Getenv(newsApiDbDsn),
		WebPort:    os.Getenv(newsApiWebPort),
		UploadPath: os.Getenv(uploadPath),
		UploadDir:  uploadDir,
	}
}
