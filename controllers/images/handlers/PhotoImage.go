package handlers

// CDNUsersPhotos godoc
// @Summary CDN Users Photos
// @Description CDN Users Photos
// @Param user_photo_id path string true "User Photo ID"
// @Tags images
// @produce png
// @Success 200 
// @Failure 400 
// @Router /images/photos/{user_photo_id} [get]
func UsersPhotos() string {
	return "./uploads/photos"
}