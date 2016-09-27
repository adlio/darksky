package darksky

type Alert struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URI         string `json:"uri"`
	Expires     int    `json:"expires"`
}
