package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap        map[string]string
	IntMap           map[string]int
	FloatMap         map[string]float32
	Data             map[string]interface{}
	Regions          []Region
	Cities           []City
	Places           []Place
	CityID           int
	CitiesFromRegion map[int][]City
	PlacesFromCity   map[int][]Place
	CSRFToken        string
	Flash            string
	Warning          string
	Error            string
}
