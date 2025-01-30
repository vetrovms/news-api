package config

import "os"

// Env Об'єкт конфігурації зі змінних оточення та інші налаштування.
type Env struct {
	LogPath          string // Шлях до файлу логування.
	DbDsn            string // Доступ до бази даних.
	WebPort          string // Порт веб застосунку.
	UploadPath       string // Шлях до директорії завантаження файлів.
	UploadDir        string // Директорія завантаження файлів.
	JwtSecretKey     string // Таємний ключ для генерації jwt токена.
	RetrospectiveUrl string // URL на сервері авторизації для ретроспективної перевірки jwt токена.
}

// newsApiLogPath Назва змінної оточення що містить шлях до файлу логування.
const newsApiLogPath = "NEWS_API_LOG_PATH"

// newsApiDbDsn Назва змінної оточення що містить доступ до бази даних.
const newsApiDbDsn = "NEWS_API_DB_DSN"

// newsApiWebPort Назва змінної оточення що містить порт веб застосунку.
const newsApiWebPort = "NEWS_API_WEB_PORT"

// uploadPath Назва змінної оточення що містить шлях до директорії завантаження файлів.
const uploadPath = "NEWS_API_UPLOAD_PATH"

// uploadDir Директорія завантаження файлів.
const uploadDir = "uploads"

// jwtSecretKey Назва змінної оточення що містить таємний ключ для генерації jwt токена.
const jwtSecretKey = "NEWS_API_JWT_SECRET_KEY"

// retrospectiveUrl Назва змінної оточення що містить URL на сервері авторизації для ретроспективної перевірки jwt токена.
const retrospectiveUrl = "NEWS_API_RETROSPECTIVE_URL"

// NewEnv Повертає об'єкт конфігурації, заповнений зі змінних оточення.
func NewEnv() Env {
	return Env{
		LogPath:          os.Getenv(newsApiLogPath),
		DbDsn:            os.Getenv(newsApiDbDsn),
		WebPort:          os.Getenv(newsApiWebPort),
		UploadPath:       os.Getenv(uploadPath),
		UploadDir:        uploadDir,
		JwtSecretKey:     os.Getenv(jwtSecretKey),
		RetrospectiveUrl: os.Getenv(retrospectiveUrl),
	}
}
