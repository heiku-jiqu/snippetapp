package main

import "github.com/heiku-jiqu/snippetapp/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
