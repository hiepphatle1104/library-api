package main

import (
	"library-api/auth"
	"library-api/database"
	"library-api/handler"
	"library-api/repository"
	"library-api/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// database
	db, err := database.NewConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.With(auth.VerifyPostData).Post("/sign-up", userHandler.NewAccount)
			r.With(auth.VerifyPostData).Post("/sign-in", userHandler.NewSession)
			r.Get("/sign-out", userHandler.ClearSession)
		})
	})

	http.ListenAndServe(":8899", r)
}
