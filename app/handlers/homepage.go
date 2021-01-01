package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Homepage renders the index page
func (hs *Handlers) Homepage() func(w http.ResponseWriter, r *http.Request) {
	//posts, err := post.GetAll(hs.JDB, 0)
	//if err != nil {
	//}
	//pages, err := page.GetAll(hs.j, true)
	//if err != nil {
	//	return err
	//}

	//fmt.Println("posts", posts)
	//fmt.Println("pages", pages)

	//return c.Render(
	//	http.StatusOK,
	//	"homepage.html",
	//    //map[string]interface{}{"posts": posts, "pages": pages},
	//    map[string]interface{}{"posts": posts},
	//)
	//renderer := template.GetRenderer("./js/themes/%s/views/*.html", "default")

	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Asas")
		//vars := mux.Vars(r)
		//w.WriteHeader(http.StatusOK)
		//fmt.Println("asasasassa")
		//fmt.Fprintf(w, "Category: %v\n", vars["coin"])

		//js, err := json.Marshal(posts)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//w.Header().Set("Content-Type", "application/json")
		//w.Write(js)
		//err := renderer.Render(w, "homepage.html", map[string]interface{}{"posts": posts})
		//if err != nil {
		//
		//}

	}
}

func ApiData() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Asas")
		vars := mux.Vars(r)
		w.WriteHeader(http.StatusOK)
		fmt.Println("asasasassa")
		fmt.Fprintf(w, "Category: %v\n", vars["coin"])
	}
}
