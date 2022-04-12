package models

import "fmt"

type Model struct {
	*Page
}

func (m Model) Store() {
	fmt.Println(m.Page)
}
