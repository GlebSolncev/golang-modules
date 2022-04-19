package models

type (
	Model struct {
		M map[string]interface{}
	}

	Todo struct {
		Id     int    `json:"id"`
		Slug   string `json:"slug"`
		Name   string `json:"name"`
		Status string `json:"status"`
		Model
	}
	Todos struct {
		Todos []*Todo `json:"todo"`
	}
)
