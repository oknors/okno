package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

// APIPagesUpdate updates a page
func (hs Handlers) APIPagesUpdate(c echo.Context) error {
	type Request struct {
		PostID         string `json:"post_id" validate:"required,min=1"`
		PositionChange int    `json:"position_change"`
		Slug           string `json:"slug" validate:"required,min=1"`
		InNavigation   bool   `json:"in_navigation"`
	}
	type Response struct {
		ID           string    `json:"id"`
		Index        int       `json:"index"`
		Slug         string    `json:"slug"`
		InNavigation bool      `json:"in_navigation"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	//req := Request{}
	//err := c.Bind(&req)
	//if err != nil {
	//	return err
	//}
	//id := c.Param("id")
	////id, err := strconv.Atoi(c.Param("id"))
	////if err != nil {
	////	return err
	////}
	//err = validator.New().Struct(req)
	//if err != nil {
	//	return err
	//}
	//page, err := pageModel.GetByID(hs.DB, id)
	//if err != nil {
	//	return err
	//}
	//err = page.Update(
	//	hs.DB,
	//	req.PostID,
	//	page.Index+req.PositionChange,
	//	req.Slug,
	//	req.InNavigation,
	//)
	//if err != nil {
	//	return err
	//}
	//res := Response{
	//	ID:           page.ID,
	//	Index:        page.Index,
	//	Slug:         page.Slug,
	//	InNavigation: page.InNavigation,
	//	CreatedAt:    page.CreatedAt,
	//	UpdatedAt:    page.UpdatedAt,
	//}
	//return c.JSON(http.StatusOK, &res)
	return nil
}
