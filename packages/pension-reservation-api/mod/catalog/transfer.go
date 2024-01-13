package catalog

type getCatalogResponseItem struct {
	Dayfee   int    `json:"dayfee"`
	Id       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
