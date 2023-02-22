package entities

// import "go.mongodb.org/mongo-driver/bson"

type Product struct {
	// Id       int64 "json"
	// Name     string "json"
	// Price    float32 "json"
	// Quantity int "json"
	// Status   bool "json"
	// var id int64
	// var name string
	// var price float32
	// var quantity int
	// var status bool

	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Edad      int    `json:"edad"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
