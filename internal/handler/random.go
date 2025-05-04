package handler

import (
    "encoding/json"
    "net/http"
    "app/internal/service"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            json.NewEncoder(w).Encode(service.GetRandom())
        case http.MethodPost:
            var msg int
            if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
                http.Error(w, "Bad request", http.StatusBadRequest)
                return
            }
            service.SetLimit(msg)
            w.WriteHeader(http.StatusCreated)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
