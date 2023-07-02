package api

type MovieFilter struct {
	Title    string `json:"title,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Director string `json:"director,omitempty"`
}

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Cast        string `json:"cast"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director,omitempty"`
}
