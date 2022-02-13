package models

// Actor struct

type Actor struct {
	Id                string `json:"id"`
	Actor_First_Name  string `json:"actor_first_name"`
	Actor_Last_Name   string `json:"actor_last_name"`
	Film_id           string `json:"film_id"`
	Film_title        string `json:"film_title"`
	Film_description  string `json:"film_description"`
	Film_language_id  string `json:"film_language_id"`
	Film_length       string `json:"film_length"`
	Film_release_year string `json:"film_release_year"`
	Film_rating       string `json:"film_rating"`
	Category_name     string `json:"category_name"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Actor `json:"data"`
	Message string  `json:"message"`
}

// Artist struct

type Artist struct {
	Id                 string `json:"id"`
	Artist_Name        string `json:"artist_name"`
	Album_title        string `json:"album_title"`
	Track_name         string `json:"track_name"`
	Track_composer     string `json:"track_composer"`
	Track_milliseconds string `json:"track_milliseconds"`
	Genre_name         string `json:"genre_name"`
	Media_type_name    string `json:"media_type_name"`
	Playlist_name      string `json:"playlist_name"`
}

type JsonResponse2 struct {
	Type    string   `json:"type"`
	Data    []Artist `json:"data"`
	Message string   `json:"message"`
}

// customer struct

type Customer struct {
	Id                  string `json:"id"`
	Customer_First_Name string `json:"customer_first_name"`
	Customer_Last_Name  string `json:"customer_last_name"`
	Customer_company    string `json:"customer_company"`
	Customer_email      string `json:"customer_email"`
	Customer_phone      string `json:"customer_phone"`
	Invoice_date        string `json:"invoice_date"`
	Billing_address     string `json:"billing_address"`
	Billing_city        string `json:"billing_city"`
	Billing_country     string `json:"billing_country"`
	Billing_postal_code string `json:"billing_postal_code"`
	Billing_state       string `json:"billing_state"`
	Unit_price          string `json:"unit_price"`
	Quantity            string `json:"quantity"`
}

type JsonResponse3 struct {
	Type    string     `json:"type"`
	Data    []Customer `json:"data"`
	Message string     `json:"message"`
}

// Employee struct

type Employee struct {
	Id          string `json:"id"`
	Last_Name   string `json:"last_name"`
	First_Name  string `json:"first_name"`
	Title       string `json:"title"`
	Reports_to  int    `json:"reports_to"`
	Birth_date  string `json:"birth_date"`
	Hire_date   string `json:"hire_date"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Postal_code string `json:"postal_code"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Email       string `json:"email"`
}

type JsonResponse4 struct {
	Type    string     `json:"type"`
	Data    []Employee `json:"data"`
	Message string     `json:"message"`
}
