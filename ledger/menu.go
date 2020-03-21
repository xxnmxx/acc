package ledger

import (
	"fmt"
	"log"
)

type menu struct {
	Items    []item
	Selected []int
}

type item interface {
	Name() string
	List() []string
}

func CreateMenu() *menu {
	m := new(menu)
	m.Items = make([]item, 0)
	m.Selected = make([]int, 0)
	return m
}

func (m *menu) InputSelected() {
	for i, item := range m.Items {
		var sel int
		temp := item.List()
		for j, v := range temp {
			fmt.Printf("%v. %v ", j, v)
		}
		fmt.Printf("\nSelect %v:", m.Items[i].Name)
		if _, err := fmt.Scan(&sel); err != nil {
			log.Fatal(err)
		}
		m.Selected = append(m.Selected, sel)
	}
}
