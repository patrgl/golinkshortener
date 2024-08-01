<h1>A Link/URL shortener made in Go</h1>

This is a simple Link/URL shortener made in Golang using only  Std and GORM libraries.
The random string after the "/" is guaranteed to always be unique by quering an SQLite DB.
Used HTMX in the HTML so when a user submits a link to be shortened, the new shortened link is shown without needing to refresh.
