package main

import "snippetbox.0xsingularity.com/internal/models"

//A templateData type that will act as a holding structure for any dynamic data
//that we want to pass to our HTML templates

type templateData struct {
	Snippet models.Snippet
}
