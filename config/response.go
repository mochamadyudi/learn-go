package config

type Pagination struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	Limit   int `json:"limit"`
	Maxpage int `json:"max_page"`
}

type Response[T any] struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Error      any         `json:"error"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       T           `json:"data"`
}
type ResponsePaginate[T any] struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       T           `json:"data"`
}

func ResponseError(message string) Response[string] {
	return Response[string]{
		Code:    "99",
		Message: message,
	}
}

func ResponseValidator[T any](error T, message string) Response[T] {
	return Response[T]{
		Code:    "99",
		Message: message,
		Error:   error,
	}
}

func ResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Code:    "00",
		Message: "success",
		Data:    data,
	}
}

func ResponseCreateSuccess[T any](data T, message string) Response[T] {
	return Response[T]{
		Code:    "00",
		Message: message,
		Data:    data,
	}
}
func ResponseWithPagination[T any](data T, pagination Pagination) ResponsePaginate[T] {
	return ResponsePaginate[T]{
		Code:    "00",
		Message: "success",
		Pagination: &Pagination{
			Total:   pagination.Total,
			Page:    pagination.Page,
			Limit:   pagination.Limit,
			Maxpage: pagination.Maxpage,
		},
		Data: data,
	}
}
