package main

import (
	"time"

	"gorm.io/gorm"
)

type Status string
type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

const (
	Pending    Status = "pending"
	InProgress Status = "in progress"
	Finished   Status = "finished"
	Cancelled  Status = "cancelled"
)

type Task struct {
	gorm.Model
	Task     string    `json:"task" gorm:"not null"`
	Priority Priority  `json:"priority"`
	DueDate  time.Time `json:"due_date"`
	DueTime  string    `json:"due_time"`
	Status   Status    `json:"status" gorm:"default:'pending'"`
}
