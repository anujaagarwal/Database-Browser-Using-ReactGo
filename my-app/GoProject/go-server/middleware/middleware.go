package middleware

// imported packages
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-server/models"

	_ "github.com/lib/pq"
)

// postgres database server details
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "anuja"
	dbname   = "testdb"
)

// setting up connection with the db

func setupDB() *sql.DB {
	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open database
	db, err := sql.Open("postgres", psqlconn)

	CheckError(err)

	// fmt.Println("Connected!")

	return db
}

// Get all Actors

// response and request handlers
func GetActors(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := setupDB()

	printMessage("Getting actors...")

	query1 := `SELECT actors.id, actors.actor_first_name, actors.actor_last_name,       film_actor.film_id, films.film_title, films.film_description,
				films.film_language_id, films.film_length, films.film_release_year, films.film_rating, categories.category_name
				FROM actors
				INNER JOIN film_actor
				ON actors.id = film_actor.actor_id
				JOIN film_category
				ON film_category.film_id = film_actor.film_id
				JOIN films
				ON films.id = film_category.film_id
				JOIN categories
				ON categories.id = film_category.category_id`

	// Get all actors by joining film and category tables with it based on foreign key.
	rows, err := db.Query(query1)

	// check errors
	CheckError(err)

	// var response []JsonResponse
	var actors []models.Actor

	// Foreach actor
	for rows.Next() {
		var id string
		var actor_first_name string
		var actor_last_name string
		var film_id string
		var film_title string
		var film_description string
		var film_language_id string
		var film_length string
		var film_release_year string
		var film_rating string
		var category_name string

		err = rows.Scan(&id, &actor_first_name, &actor_last_name, &film_id, &film_title, &film_description, &film_language_id, &film_length, &film_rating, &film_release_year, &category_name)

		// check errors
		CheckError(err)

		actors = append(actors, models.Actor{Id: id, Actor_First_Name: actor_first_name, Actor_Last_Name: actor_last_name, Film_id: film_id, Film_title: film_title, Film_description: film_description, Film_language_id: film_language_id, Film_length: film_length, Film_release_year: film_release_year, Film_rating: film_rating, Category_name: category_name})
	}

	var response = models.JsonResponse{Type: "success", Data: actors}

	json.NewEncoder(w).Encode(response)
}

// Get all Artists
func GetArtists(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := setupDB()

	printMessage("Getting artists...")

	query2 := `SELECT
				DISTINCT artists.id,
				artists.artist_name, albums.album_title, tracks.track_name, 
				COALESCE(tracks.track_composer, 'No Information') track_composer,
				tracks.track_milliseconds,
				genres.genre_name, media_types.media_type_name, playlists.playlist_name
				FROM artists
				JOIN albums
				ON albums.artist_id = artists.id
				JOIN tracks
				ON tracks.album_id = albums.id
				JOIN genres
				ON genres.id = tracks.genre_id
				JOIN media_types
				ON media_types.id = tracks.genre_id
				JOIN playlist_track
				ON playlist_track.track_id = tracks.id
				JOIN playlists
				ON playlists.id = playlist_track.playlist_id`

	// Get all artists from artists table and joining tables like album, genre, playlist and track with it.
	rows, err := db.Query(query2)

	// check errors
	CheckError(err)

	// var response []JsonResponse
	var artists []models.Artist

	// Foreach artist
	for rows.Next() {
		var id string
		var artist_name string
		var album_title string
		var track_name string
		var track_composer string
		var track_milliseconds string
		var genre_name string
		var media_type_name string
		var playlist_name string

		err = rows.Scan(&id, &artist_name, &album_title, &track_name, &track_composer, &track_milliseconds, &genre_name, &media_type_name, &playlist_name)

		// check errors
		CheckError(err)

		artists = append(artists, models.Artist{Id: id, Artist_Name: artist_name, Album_title: album_title, Track_name: track_name, Track_composer: track_composer, Track_milliseconds: track_milliseconds, Genre_name: genre_name, Media_type_name: media_type_name, Playlist_name: playlist_name})
	}

	var response = models.JsonResponse2{Type: "success", Data: artists}

	json.NewEncoder(w).Encode(response)
}

// Get all customers

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := setupDB()

	printMessage("Getting customers...")

	query3 := `SELECT
				DISTINCT customers.id,
				customers.customer_first_name, customers.customer_last_name,
				COALESCE(customers.customer_company, 'NO COMPANY INFO') customer_company,
				COALESCE(customers.customer_phone, 'XXXXXX') customer_phone,
				customers.customer_email, invoices.invoice_date, invoices.billing_address,
				invoices.billing_city, 
				COALESCE(invoices.billing_state, 'NA') billing_state, 
				invoices.billing_country, 
				COALESCE(invoices.billing_postal_code, 'NO CODE') billing_postal_code,
				invoice_lines.unit_price, invoice_lines.quantity
				FROM customers
				JOIN invoices
				ON invoices.customer_id = customers.id
				JOIN invoice_lines
				ON invoice_lines.invoice_id = invoices.id`

	// Get all customers from customer table by joining invoice table with it.
	rows, err := db.Query(query3)

	// check errors
	CheckError(err)

	// var response []JsonResponse
	var customers []models.Customer

	// Foreach customer
	for rows.Next() {
		var id string
		var customer_first_name string
		var customer_last_name string
		var customer_company string
		var customer_phone string
		var customer_email string
		var invoice_date string
		var billing_address string
		var billing_city string
		var billing_country string
		var billing_postal_code string
		var billing_state string
		var unit_price string
		var quantity string

		err = rows.Scan(&id, &customer_first_name, &customer_last_name, &customer_company, &customer_phone, &customer_email, &invoice_date, &billing_address, &billing_city, &billing_country, &billing_postal_code, &billing_state, &unit_price, &quantity)

		// check errors
		CheckError(err)

		customers = append(customers, models.Customer{Id: id, Customer_First_Name: customer_first_name, Customer_Last_Name: customer_last_name, Customer_company: customer_company, Customer_phone: customer_phone, Customer_email: customer_email, Invoice_date: invoice_date,
			Billing_address: billing_address, Billing_city: billing_city, Billing_country: billing_country, Billing_postal_code: billing_postal_code, Billing_state: billing_state, Unit_price: unit_price, Quantity: quantity})
	}

	var response = models.JsonResponse3{Type: "success", Data: customers}

	json.NewEncoder(w).Encode(response)
}

// Get all employees

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := setupDB()

	printMessage("Getting employees...")

	query4 := `SELECT * FROM employees`

	// Get all employee from employees table.
	rows, err := db.Query(query4)

	// check errors
	CheckError(err)

	// var response []JsonResponse
	var employees []models.Employee

	// Foreach employee
	for rows.Next() {
		var id string
		var last_name string
		var first_name string
		var title string
		var reports_to int
		var birth_date string
		var hire_date string
		var address string
		var city string
		var state string
		var country string
		var postal_code string
		var phone string
		var fax string
		var email string

		err = rows.Scan(&id, &last_name, &first_name, &title, &reports_to, &birth_date, &hire_date, &address, &city, &state, &country, &postal_code, &phone, &fax, &email)

		// check errors
		CheckError(err)

		employees = append(employees, models.Employee{Id: id, Last_Name: last_name, First_Name: first_name, Title: title, Reports_to: reports_to, Birth_date: birth_date, Hire_date: hire_date, Address: address, City: city, State: state, Country: country, Postal_code: postal_code, Phone: phone, Fax: fax, Email: email})
	}

	var response = models.JsonResponse4{Type: "success", Data: employees}

	json.NewEncoder(w).Encode(response)
}

// Function for handling messages

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// function to enable cors

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
