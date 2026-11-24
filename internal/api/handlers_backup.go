package api

import (
	"encoding/json"
	"net/http"
)

type backupInitResponse struct {
	BackupID  string `json:"backup_id"`
	UploadURL string `json:"upload_url"`
}

func handleBackupInit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := backupInitResponse{
		BackupID:  "12345",
		UploadURL: "https://example.com/upload",
	}
	json.NewEncoder(w).Encode(response)

func handleBackupComplete(w http.ResponseWriter, r *http.Request) {
	
}

}
