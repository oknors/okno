package wing

import (
	"encoding/json"
	"gioui.org/widget"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/models/wing/db"
	"github.com/w-ingsolutions/c/model"
	"net/http"
	"strconv"
)

type WingCal struct {
	Naziv          string
	Strana         string
	Edit           bool
	Materijal      map[int]*model.WingMaterijal
	Radovi         model.WingVrstaRadova
	IzbornikRadova *model.WingVrstaRadova
	Transfered     model.WingCalGrupaRadova
	Db             *db.DuoUIdb
	//Client           *model.Client
	PrikazaniElement *model.WingVrstaRadova
	Suma             *model.WingIzabraniElementi
}

func NewWingCal() *WingCal {
	wing := &WingCal{
		Naziv:            "W-ing Solutions - Kalkulator",
		Db:               db.DuoUIdbInit("DATABASE/wing"),
		PrikazaniElement: &model.WingVrstaRadova{},
		//Suma: &model.WingIzabraniElementi{
		//	UkupanNeophodanMaterijal: map[int]model.WingNeophodanMaterijal{},
		//},
	}
	//wing.NewMaterijal()
	wing.Radovi = model.WingVrstaRadova{
		Id:             0,
		Naziv:          "Radovi",
		Slug:           "radovi",
		Omogucen:       false,
		Baza:           false,
		Element:        false,
		PodvrsteRadova: wing.Db.DbReadAll("radovi"),
	}

	return wing
}

func (w *WingCal) GenerisanjeEdita() (edit *model.EditabilnaPoljaVrsteRadova) {
	//w.EditabilnaPoljaVrsteRadova = make(map[int]*model.EditabilnaPoljaVrsteRadova)
	//for rad, _ := range radovi {
	//	w.EditabilnaPoljaVrsteRadova[rad] =
	return &model.EditabilnaPoljaVrsteRadova{
		Id:    new(widget.Editor),
		Naziv: new(widget.Editor),
		Opis: &widget.Editor{
			SingleLine: false,
		},
		Obracun:  new(widget.Editor),
		Jedinica: new(widget.Editor),
		Cena:     new(widget.Editor),
		Slug:     new(widget.Editor),
		Omogucen: new(widget.Bool),
	}
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func (wc *WingCal) VrsteRadova(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	radovi := make(map[int]model.ElementMenu)
	for vr, rd := range wc.Radovi.PodvrsteRadova {
		radovi[vr] = model.ElementMenu{
			Id:    rd.Id,
			Title: rd.Naziv,
			Slug:  rd.Slug,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) PodvrsteRadova(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	radovi := make(map[int]model.ElementMenu)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	}
	for vr, rd := range wc.Radovi.PodvrsteRadova[id+1].PodvrsteRadova {
		radovi[vr] = model.ElementMenu{
			Id:    rd.Id,
			Title: rd.Naziv,
			Slug:  rd.Slug,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) Elementi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	radovi := make(map[int]model.ElementMenu)
	elementi := wc.Db.DbRead(params["id"], params["el"])
	for vr, rd := range elementi.PodvrsteRadova {
		var m bool
		if rd.NeophodanMaterijal != nil {
			m = true
		}
		radovi[vr] = model.ElementMenu{
			Id:        rd.Id,
			Title:     rd.Naziv,
			Slug:      rd.Slug,
			Materijal: m,
		}
	}
	json.NewEncoder(w).Encode(radovi)
}

func (wc *WingCal) Element(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//id, err := strconv.Atoi(params["id"])
	//if err != nil{}
	//el, err := strconv.Atoi(params["el"])
	//if err != nil{}
	e, err := strconv.Atoi(params["e"])
	if err != nil {
	}

	elementi := wc.Db.DbRead(params["id"], params["el"])

	json.NewEncoder(w).Encode(elementi.PodvrsteRadova[e-1])
}

//func (wc *WingCal)getPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	for _, item := range posts {
//		if item.ID == params["id"] {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(&Post{})
//}
