package main

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"school-api/models"
)

func main() {
	// Initialize the database and models as shown above

	r := chi.NewRouter()

	// Teachers
	r.Post("/teachers", createTeacher)
	r.Get("/teachers", getTeachers)
	r.Get("/teachers/{teacherID}/students", getStudentsByTeacher)

	// Classrooms
	r.Post("/classrooms", createClassroom)
	r.Get("/classrooms", getClassrooms)

	// Students
	r.Post("/students", createStudent)
	r.Get("/students/{studentID}/teacher", getTeacherByStudent)

	http.ListenAndServe(":8080", r)
}

func createTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&teacher)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func getTeachers(w http.ResponseWriter, r *http.Request) {
	var teachers []models.Teacher
	db.Find(&teachers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

func getStudentsByTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "teacherID")
	var students []models.Student
	db.Model(&models.Teacher{}).Where("ID = ?", teacherID).Association("Students").Find(&students)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func createClassroom(w http.ResponseWriter, r *http.Request) {
	var classroom models.Classroom
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&classroom); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&classroom)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classroom)
}

func getClassrooms(w http.ResponseWriter, r *http.Request) {
	var classrooms []models.Classroom
	db.Find(&classrooms)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classrooms)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&student)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func getTeacherByStudent(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "studentID")
	var teacher models.Teacher
	db.Model(&models.Student{}).Where("ID = ?", studentID).Association("Classroom.Teacher").Find(&teacher)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}
