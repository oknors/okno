package handlers

import (
	"github.com/labstack/echo/v4"
)

// Page renders a navigable page that's not just a post
func (hs Handlers) Page(c echo.Context) error {
	//page, err := pageModel.GetBySlug(hs.j, c.Param("slug"))
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return c.String(http.StatusNotFound, "Not found")
	//	}
	//	return err
	//}
	//post, err := postModel.GetByID(hs.DB, page.PostID)
	//if err != nil {
	//	return err
	//}
	//pages, err := pageModel.GetAll(hs.DB, true)
	//if err != nil {
	//	return err
	//}
	//return c.Render(
	//	http.StatusOK,
	//	"page.html",
	//	map[string]interface{}{"post": post, "pages": pages},
	//)
	return nil
}
