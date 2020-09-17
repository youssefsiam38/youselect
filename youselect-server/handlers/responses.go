package handlers

type MessageRes struct{
	Message string `json:"message"`
}
type ErrorRes struct{
	Error string `json:"error"`
}