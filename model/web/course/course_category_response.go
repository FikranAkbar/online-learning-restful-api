package course

type CategoryResponse struct {
	Id               uint   `json:"id"`
	CategoryName     string `json:"category_name"`
	CategoryPhotoUrl string `json:"category_photo_url"`
}
