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
	response := backupInitResponse{
		BackupID:  "12345",
		UploadURL: "https://example.com/upload",
	}
	writeJSON(w, http.StatusCreated, response)
}

func handleBackupComplete(w http.ResponseWriter, r *http.Request) {
	var req backupCompleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{
			Error: "invalid request payload",
		})
		return
	}

	if req.BackupID == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{
			Error: "backup_id is required",
		})
		return
	}
	response := backupCompleteResponse{
		BackupID: req.BackupID,
		Status:   "verification_scheduled",
	}
	writeJSON(w, http.StatusAccepted, response)
}
