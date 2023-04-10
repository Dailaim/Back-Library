package handlers


// CDNUsersPhotos godoc
// @Summary CDN Books Images
// @Description CDN Books Images
// @Param book_image_id path string true "Book Image ID"
// @Tags images
// @produce png
// @Success 200 
// @Failure 400 
// @Router /uploads/ImagesBooks/{book_image_id} [get]
func BooksImages() string {
	return "./uploads/ImagesBooks"
}