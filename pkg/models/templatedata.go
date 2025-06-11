package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // since we don't know what all types of data we use "interface{}"
	CSRFToken string
	Flash     string // when just posting a message to the screen
	Warning   string
	Error     string
}
