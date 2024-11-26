package model

type SampleData struct {
	Data string
}

type Response struct {
	Status  int
	Data    []interface{}
	Message string
}
