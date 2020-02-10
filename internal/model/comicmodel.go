package comicModel

//Comic structure
type Comic struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Stock int    `json:"Stock"`
}
