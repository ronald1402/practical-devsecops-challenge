package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"practical-devsecops-technical/dto"
	"practical-devsecops-technical/response"
	"practical-devsecops-technical/service"

	"github.com/sirupsen/logrus"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/index" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Something error, keep calm we will fix it soon", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something error", http.StatusInternalServerError)
		return
	}

}

func JDoodleHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Process JDoodle Request")
	var dto dto.JDoodleDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		logrus.Error(err.Error())
		response.Response(w, "unexpected error", http.StatusInternalServerError)
		return
	}
	logrus.Info("Execute service")
	responseData, statusCode, err := service.ExecuteJDoodle(dto)
	if err != nil {
		logrus.Error(err.Error())
		response.Response(w, "unexpected error", http.StatusInternalServerError)
		return
	}
	logrus.Info("Success")
	response.ResponseByte(w, responseData, statusCode)
}
