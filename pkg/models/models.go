package models

// Region is the region model
type Region struct {
	ID        int
	CityCount int
	Area      int
	Country   string
	Name      string
	ImgHref   string
}

// City is the city model
type City struct {
	ID         int
	RegionID   int
	Population int
	TimeZone   string
	Name       string
	ImgHref    string
}

// Place is the place model
type Place struct {
	ID          int
	CityID      int
	ImgHref     string
	Name        string
	Description string
}
