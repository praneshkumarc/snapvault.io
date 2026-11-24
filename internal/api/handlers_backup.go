package api

import (
	"encoding/json"
	"net/http"
)

type backupInitResponse struct {
	BackupID  string `json:"backup_id"`
	UploadURL string `json:"upload_url"`
}

type backupCompleteResponse struct {
	BackupID string `json:"backup_id"`
	Status   string `json:"status"`
}

type backupCompleteRequest struct {
	BackupID string `json:"backup_id"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func handleBackupInit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := backupInitResponse{
		BackupID:  "12345",
		UploadURL: "https://example.com/upload",
	}
	json.NewEncoder(w).Encode(response)
}

func handleBackupComplete(w http.ResponseWriter, r *http.Request) {
	var req backupCompleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Error: "invalid request payload",
		})
		return
	}

	if req.BackupID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Error: "missing backup_id in request payload",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	response := backupCompleteResponse{
		BackupID: req.BackupID,
		Status:   "verification_scheduled",
	}
	json.NewEncoder(w).Encode(response)
}
