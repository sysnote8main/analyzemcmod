package common

type ModInfo struct {
	Modid        string    `json:"modid"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Version      string    `json:"version"`
	Mcversion    string    `json:"mcversion"`
	Url          *string   `json:"url"`
	AuthorList   *[]string `json:"authorList"`
	LogoFile     *string   `json:"logoFile"`
	Dependencies *[]string `json:"dependencies"`
}
