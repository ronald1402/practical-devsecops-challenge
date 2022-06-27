package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"practical-devsecops-technical/config"
	"practical-devsecops-technical/handler"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func initConfig() {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err.Error())
	}

	dirpath := fmt.Sprintf(path + "/environtment")

	viper.SetConfigType("json")
	viper.AddConfigPath(dirpath)
	viper.SetConfigName("app.config")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err.Error())
	}

	err = viper.Unmarshal(&config.Config)
	if err != nil {
		log.Error(err.Error())
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	initConfig()
}

func main() {
	// mux := http.NewServeMux()
	r := chi.NewRouter()
	r = setCors(r)
	r.Use(middleware.Logger)

	r.Get("/index", handler.HomeHandler)
	r.Get("/ping", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Println("masuk cuk")
		data := map[string]interface{}{
			"data": "masuk",
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(data)

	})
	r.Post("/v1/execute", handler.JDoodleHandler)
	fileServer := http.FileServer(http.Dir("assets"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	log.Info("Running on port : " + config.Config.Port)
	err := http.ListenAndServe(":"+config.Config.Port, r)
	fmt.Println("server started at localhost:8080")
	log.Fatal(err)
}

func setCors(r *chi.Mux) *chi.Mux {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	return r
}
