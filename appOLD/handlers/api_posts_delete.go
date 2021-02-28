package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// APIPostsDelete deletes a post
func (hs Handlers) APIPostsDelete(c echo.Context) error {
	type Response struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		IsDraft   bool      `json:"is_draft"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	//id := c.Param("id")
	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	return err
	//}
	//post, err := postModel.Delete(hs.DB, id)
	//if err != nil {
	//	return err
	//}
	//res := Response{
	//	ID:        post.ID,
	//	Title:     post.Title,
	//	Content:   post.Content,
	//	IsDraft:   post.IsDraft,
	//	CreatedAt: post.CreatedAt,
	//	UpdatedAt: post.UpdatedAt,
	//}
	//return c.JSON(http.StatusOK, res)
	return nil
}
