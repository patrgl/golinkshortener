package main

import (
	"fmt"
	"golinkshortener/links"
	"golinkshortener/models"
	"net/http"
	"text/template"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func redirect_to_original(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		short := r.PathValue("short")

		//query short link given
		var link models.Link
		result := db.First(&link, "short = ?", short)
		if result.Error != nil {
			fmt.Fprintf(w, "Short link not found, please try again")
		}

		http.Redirect(w, r, link.Original, http.StatusFound)

	}
}

func submit_link(db *gorm.DB, base_url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		original_link := r.PostFormValue("original-link")
		short_link := links.ShortenLink(base_url, original_link, db)
		fmt.Fprintf(w, `<a href="%s">%s</a>`, short_link, short_link)
	}
}

func main() {
	const base_url string = "localhost:8080/"

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&models.Link{})

	mux := http.NewServeMux()

	//routes
	mux.HandleFunc("/", index)
	mux.HandleFunc("/{short}", redirect_to_original(db))
	mux.HandleFunc("POST /submit-link", submit_link(db, base_url))

	fmt.Println("Running on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
