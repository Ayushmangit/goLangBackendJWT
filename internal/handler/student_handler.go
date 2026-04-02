package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/utils"
	"github.com/Ayushmangit/goLangBackendJWT/internal/service"
	"github.com/google/uuid"
)

type StudentHandler struct {
	service service.StudentService
}

func NewStudentHandler(s service.StudentService) *StudentHandler {
	return &StudentHandler{
		service: s,
	}
}

type createStudentRequest struct {
	FullName   string `json:"full_name"`
	RollNumber string `json:"roll_number"`
	ClassID    string `json:"class_id"`
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var req createStudentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	authUser, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	classID, err := uuid.Parse(req.ClassID)
	if err != nil {
		http.Error(w, "invalid class_id", http.StatusBadRequest)
		return
	}

	student, err := h.service.Create(
		r.Context(),
		authUser.ID, // 👈 secure
		req.FullName,
		req.RollNumber,
		classID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}
