package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// APIPostsUpdate updates a post over JSON
func (hs Handlers) APIPostsUpdate(c echo.Context) error {
	type Request struct {
		Title   string `json:"title" validate:"required,min=1"`
		Content string `json:"content" validate:"min=10"`
		Slug    string `json:"slug" validate:"required,min=1"`
		IsDraft bool   `json:"is_draft"`
	}
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
	//req := new(Request)
	//err := c.Bind(req)
	//if err != nil {
	//	return err
	//}
	//err = validator.New().Struct(req)
	//if err != nil {
	//	return err
	//}
	//post, err := postModel.Update(
	//	hs.DB,
	//	id,
	//	req.Title,
	//	req.Content,
	//	req.Slug,
	//	req.IsDraft,
	//)
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
