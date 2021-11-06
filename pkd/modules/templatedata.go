package modules

//TemplateDat hp;ds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntData   map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
