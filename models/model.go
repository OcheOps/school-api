package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name       string
	Classrooms []Classroom `gorm:"many2many:teacher_classrooms;"`
}

type Classroom struct {
	gorm.Model
	Name      string
	TeacherID uint
	Teacher   Teacher
	Students  []Student
}

type Student struct {
	gorm.Model
	Name        string
	ClassroomID uint
	Classroom   Classroom
}
