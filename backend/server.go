package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"git.01.alem.school/Kusbek/social-network/backend/api/handler"
	"git.01.alem.school/Kusbek/social-network/backend/pkg/db/repository"
	"git.01.alem.school/Kusbek/social-network/backend/pkg/db/sqlite"
	"git.01.alem.school/Kusbek/social-network/backend/usecase/user"
)

func main() {
	//Options
	sqlOpts := &sqlite.Options{
		Address: "local.db",
	}
	apiPort := 8080

	//Main
	db, err := sqlite.Init(sqlOpts)
	if err != nil {
		log.Fatalf("Failed to start db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()

	userRepo := repository.NewUserRepository(db)
	userService := user.NewService(userRepo)
	_, err = userService.FindUser(
		"kusbek",
	)
	if err != nil {
		if err == sql.ErrNoRows {
			userService.CreateUser(
				"kusbek",
				"kusbek1994@gmail.com",
				"Bekarys",
				"Kuspan",
				"Something is important, but not this project",
				"./images/avatars/somephoto.jpg",
				"1994-09-18",
				"123456")
		} else {
			log.Fatal(err)
		}
	}

	r := http.NewServeMux()

	handler.MakeUserHandlers(r, userService)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", apiPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
