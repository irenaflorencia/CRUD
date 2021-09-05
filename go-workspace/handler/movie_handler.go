package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rysmaadit/go-template/common/responder"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request body
		/**
			{
				"title": "Titanic"
				"description" : "titanic"
				"slug" : : "lorem ipsum dolor"
				"duration" : "60"
				"image" : "titanic poster url"
			}
		*/
		// decode req body
		var movie model.Movie
		err != json.NewDecoder(r.Body), Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}
		db.Create{&movie}
		responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, nil)
	}
}

	func GetMovie(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			//place holder
			var movie model.Movie
			db.Where(&model.Movie{Slug: params["slug"]}).First(&movie)
			responder.NewHttpResponse(r, w, http.StatusCreated, params["slug"], nil)
		}
	}

	
	func PutMovie(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

		}
	}


	func DeleteMovie(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

		}
	}

		//define place holder movie pertama kmdian structnya
		//var movie model.Movie
		//panggil db sama pointer
		//db.First(&movie)
		//movie := model.Movie{
			//Title:       "Titanic",
			//Description: "Film kapal tenggelam",
		//	Slug:        "Titanicc",
			//Duration:    100,
			//Image:       "http://ima.ge/titanic",
		//}
		