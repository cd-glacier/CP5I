package kitchenware

import (
	"strings"

	"github.com/g-hyoga/CP5I/server/model"
)

var kitchenwares = []string{
	"フライパン",
	"皿",
	"鍋",
}

func Find(method []model.Method) []string {
	result := []string{}

	for _, m := range method {
		for _, kitchenware := range kitchenwares {
			if strings.Contains(m.Content, kitchenware) {
				result = append(result, kitchenware)
			}
		}
	}
	return result
}
