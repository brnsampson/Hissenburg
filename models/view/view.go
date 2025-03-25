package view

import "html/template"

type PageView struct {
	User string
	Contents template.HTML
}
