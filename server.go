package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"git.01.alem.school/Kusbek/social-network/api/handler"
	"git.01.alem.school/Kusbek/social-network/pkg/db/repository"
	"git.01.alem.school/Kusbek/social-network/pkg/db/sqlite"
	"git.01.alem.school/Kusbek/social-network/usecase/session"
	"git.01.alem.school/Kusbek/social-network/usecase/subscription"
	"git.01.alem.school/Kusbek/social-network/usecase/user"
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

	r := http.NewServeMux()
	userService := newUser(db)
	subscriptionService := newSubscription(db, userService)
	sessionService := newSession()

	handler.MakeFileHandlers(r)
	handler.MakeUserHandlers(r, sessionService, userService, subscriptionService)
	handler.MakeSubscriptionHandlers(r, sessionService, subscriptionService)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", apiPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func newSession() session.UseCase {
	repo := repository.NewSessionRepository()
	sessionService := session.NewService(repo)
	return sessionService
}

func newUser(db *sql.DB) user.UseCase {
	userRepo := repository.NewUserRepository(db)
	userService := user.NewService(userRepo)
	_, err := userService.FindUser(
		"kusbek",
	)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = userService.CreateUser(
				"kusbek",
				"kusbek1994@gmail.com",
				"Bekarys",
				"Kuspan",
				"Something is important, but not this project",
				"/img/avatars/2021-03-15 20.32.55.jpg",
				"1994-09-18",
				"123456")
			_, err = userService.CreateUser(
				"scarlett",
				"scarlett@gmail.com",
				"Scarlett",
				"Johanson",
				"I'm a black widow and used to have big tits",
				"/img/avatars/2cd43b_e0d52be694a24ce2ba7a0aeca9002e5e_mv2_d_1600_2532_s_2.png",
				"1978-05-12",
				"123456")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	return userService
}

func newSubscription(db *sql.DB, userService user.UseCase) subscription.UseCase {
	repo := repository.NewSubscriptionRepository(db)
	service := subscription.NewService(repo, userService)
	return service
}
