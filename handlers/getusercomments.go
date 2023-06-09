package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetUserComments(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Please authorize")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	claims := getData(c)

	var comments []models.Comment

	if result := h.DB.Where("author_id = ?", claims.Data.Id).Preload("Author").
		Preload("Item").Find(&comments); result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)

	fmt.Println("comments are sent")
}
