package handlers

import (
    "net/http"
    "go-rest-api/src/utils"
)

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
    key, _ := utils.GetCurrentKey()
    w.Write([]byte(key))
}

func ResetKeyHandler(w http.ResponseWriter, r *http.Request) {
    utils.RotateKey()
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Key has been reset"))
}