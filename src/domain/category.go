package domain

import (
	"strings"
)

/**
 * Category class
 * Groupings of listings of the same "category". E.g. Electronics, Fashion etc
 */

type Category struct {
	Name string
	Count int
}

func NewCategory(name string) *Category {
	return &Category{Name: strings.ToLower(name), Count: 0}
}