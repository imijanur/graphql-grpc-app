package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/imijanur/graphql-grpc-server/models" // Update with the actual path to the generated models

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/volatiletech/null/v8"
)

var (
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
)

func loadEnv() {
	// load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load the .env file: %v", err)
	}
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
}

func escapeString(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

func main() {
	// load the .env file
	loadEnv()

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true") // Update with your actual database connection details
	// Update with your actual database connection details
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Start the benchmark
	start := time.Now()

	// Generate 20,000 random users
	var users []string
	var contacts []string
	var addresses []string
	emailSet := make(map[string]struct{})

	for i := 0; i < 20000; i++ {
		var email string
		for {
			email = gofakeit.Email()
			if _, exists := emailSet[email]; !exists {
				emailSet[email] = struct{}{}
				break
			}
		}

		user := &models.User{
			Email:  email,
			Status: gofakeit.RandomString([]string{"active", "inactive"}),
		}

		users = append(users, fmt.Sprintf("('%s', '%s')",
			user.Email, user.Status))

		userID := i + 5 // Assuming auto-increment starts from 1

		// Generate random contacts for the user
		contact := &models.UserContact{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Phone:     gofakeit.Phone(),
			UserID:    userID,
		}

		contacts = append(contacts, fmt.Sprintf("('%s', '%s', '%s', %d)",
			contact.FirstName, contact.LastName, contact.Phone, contact.UserID))

		// Generate random addresses for the user
		for k := 0; k < rand.Intn(3)+1; k++ {
			address := &models.UserAddress{
				Name:           escapeString(gofakeit.Company()),
				Prefix:         null.StringFrom(escapeString(gofakeit.RandomString([]string{"Mr.", "Ms.", "Mrs.", "Dr."}))),
				StreetAddress1: escapeString(gofakeit.Street()),
				StreetAddress2: null.StringFrom(escapeString(gofakeit.Street())),
				City:           escapeString(gofakeit.City()),
				State:          escapeString(gofakeit.State()),
				ZipCode:        escapeString(gofakeit.Zip()),
				UserID:         userID,
			}

			value := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
				address.Name, address.Prefix.String, address.StreetAddress1, address.StreetAddress2.String, address.City, address.State, address.ZipCode, address.UserID)

			// fmt.Println(value)

			addresses = append(addresses, value)
		}
	}

	// Insert users in bulk
	userInsertQuery := fmt.Sprintf("INSERT INTO users (email, status) VALUES %s", strings.Join(users, ","))
	_, err = db.Exec(userInsertQuery)
	if err != nil {
		log.Fatalf("failed to insert users: %v", err)
	}

	// Insert contacts in bulk
	contactInsertQuery := fmt.Sprintf("INSERT INTO user_contact (first_name, last_name, phone, user_id) VALUES %s", strings.Join(contacts, ","))
	_, err = db.Exec(contactInsertQuery)
	if err != nil {
		log.Fatalf("failed to insert contacts: %v", err)
	}

	// Insert addresses in bulk
	addressInsertQuery := fmt.Sprintf("INSERT INTO user_address (name, prefix, street_address_1, street_address_2, city, state, zip_code, user_id) VALUES %s", strings.Join(addresses, ","))
	_, err = db.Exec(addressInsertQuery)
	if err != nil {
		log.Fatalf("failed to insert addresses: %v", err)
	}

	// End the benchmark
	elapsed := time.Since(start)
	log.Printf("Benchmark completed in %s", elapsed)
}
