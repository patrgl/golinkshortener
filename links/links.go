package links

import (
	"fmt"
	"golinkshortener/models"
	"math/rand"

	"gorm.io/gorm"
)

func ShortenLink(base_url string, original string, db *gorm.DB) string {
	for {
		//create short
		const chars string = "abcdefghijklmnopqrstuvwxyz0123456789"
		short := make([]byte, 6)
		for i := 0; i < 6; i++ {
			short[i] += chars[rand.Intn(len(chars))]
		}
		string_short := string(short)

		//check if randomly-generated already exists in db, if it doesnt, function ends and returns the shortened link
		query_link := models.Link{}
		result := db.First(&query_link, "short = ?", string_short)
		if result.Error != nil { // if new short link does not exist, add to db and return it
			submit_link := models.Link{Original: original, Short: string_short}
			s := db.Create(&submit_link)
			if s.Error != nil {
				fmt.Println("Error submitting new link")
			}
			return base_url + string_short
		}
		fmt.Println("Randomly generated short-link already exists, trying again...")

	}

}
