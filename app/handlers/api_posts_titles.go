package handlers

import (
	"github.com/labstack/echo/v4"
)

// APIPostsTitles gets titles of posts for select elements
func (hs Handlers) APIPostsTitles(c echo.Context) error {
	type ResponseItem struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	//posts, err := postModel.GetAll(hs.DB, -1)
	//if err != nil {
	//	return err
	//}
	//res := []ResponseItem{}
	//for _, post := range *posts {
	//	resItem := ResponseItem{
	//		ID:    post.ID,
	//		Title: post.Title,
	//	}
	//	res = append(res, resItem)
	//}
	//return c.JSON(http.StatusOK, &res)
	return nil
}
