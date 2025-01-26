package response

import "news/internal/models"

// Response Структура відповіді.
type Response struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  any    `json:"data"`
}

// NewResponse Повертає структуру відповіді.
func NewResponse(code int, error string, data any) *Response {
	return &Response{
		Code:  code,
		Error: error,
		Data:  data,
	}
}

// DocGetNewsArticlesResponse200 Приклад відповіді списка новин для генерації документації.
type DocGetNewsArticlesResponse200 struct {
	Code   int                     `json:"code" example:"200"`
	Errors []string                `json:"errors"`
	Data   []models.NewsArticleDTO `json:"data"`
}

// DocGetNewsArticlesResponse500 Приклад відповіді списка новин для генерації документації.
type DocGetNewsArticlesResponse500 struct {
	Code   int      `json:"code" example:"500"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsArticleResponse200 Приклад відповіді інформації про новину для генерації документації.
type DocGetNewsArticleResponse200 struct {
	Code   int                   `json:"code" example:"200"`
	Errors []string              `json:"errors"`
	Data   models.NewsArticleDTO `json:"data"`
}

// DocGetNewsArticleResponse400 Приклад відповіді інформації про новину для генерації документації.
type DocGetNewsArticleResponse400 struct {
	Code   int      `json:"code" example:"400"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsArticleResponse404 Приклад відповіді інформації про новину для генерації документації.
type DocGetNewsArticleResponse404 struct {
	Code   int      `json:"code" example:"404"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsArticleResponse500 Приклад відповіді інформації про новину для генерації документації.
type DocGetNewsArticleResponse500 struct {
	Code   int      `json:"code" example:"500"`
	Errors []string `json:"errors" example:"щось пішло не так"`
	Data   []string `json:"data"`
}

// DocGetNewsGroupsResponse200 Приклад відповіді списку груп новин для генерації документації.
type DocGetNewsGroupsResponse200 struct {
	Code   int                   `json:"code" example:"200"`
	Errors []string              `json:"errors"`
	Data   []models.NewsGroupDTO `json:"data"`
}

// DocGetNewsGroupsResponse500 Приклад відповіді списку груп новин для генерації документації.
type DocGetNewsGroupsResponse500 struct {
	Code   int      `json:"code" example:"500"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsGroupResponse200 Приклад відповіді інформації про групу новин для генерації документації.
type DocGetNewsGroupResponse200 struct {
	Code   int                 `json:"code" example:"200"`
	Errors []string            `json:"errors"`
	Data   models.NewsGroupDTO `json:"data"`
}

// DocGetNewsGroupResponse400 Приклад відповіді інформації про групу новин для генерації документації.
type DocGetNewsGroupResponse400 struct {
	Code   int      `json:"code" example:"400"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsGroupResponse404 Приклад відповіді інформації про групу новин для генерації документації.
type DocGetNewsGroupResponse404 struct {
	Code   int      `json:"code" example:"404"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetNewsGroupResponse500 Приклад відповіді інформації про групу новин для генерації документації.
type DocGetNewsGroupResponse500 struct {
	Code   int      `json:"code" example:"500"`
	Errors []string `json:"errors" example:"щось пішло не так"`
	Data   []string `json:"data"`
}

// DocGetFileUploads200 Приклад відповіді списку Файлів для генерації документації.
type DocGetFileUploads200 struct {
	Code   int                    `json:"code" example:"200"`
	Errors []string               `json:"errors"`
	Data   []models.FileUploadDto `json:"data"`
}

// DocGetFileUpload200 Приклад відповіді інформації про файл для генерації документації.
type DocGetFileUpload200 struct {
	Code   int                  `json:"code" example:"200"`
	Errors []string             `json:"errors"`
	Data   models.FileUploadDto `json:"data"`
}

// DocGetFileUpload400 Приклад відповіді інформації про файл для генерації документації.
type DocGetFileUpload400 struct {
	Code   int      `json:"code" example:"400"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetFileUpload404 Приклад відповіді інформації про файл для генерації документації.
type DocGetFileUpload404 struct {
	Code   int      `json:"code" example:"404"`
	Errors []string `json:"errors"`
	Data   []string `json:"data"`
}

// DocGetFileUpload500 Приклад відповіді інформації про файл для генерації документації.
type DocGetFileUpload500 struct {
	Code   int      `json:"code" example:"500"`
	Errors []string `json:"errors" example:"щось пішло не так"`
	Data   []string `json:"data"`
}
