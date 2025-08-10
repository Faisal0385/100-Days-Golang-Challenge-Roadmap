package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ConnectDatabase()

	app := fiber.New()

	// Enable CORS for all origins (you can restrict later)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/tasks", GetTasks)
	app.Post("/tasks", CreateTask)
	app.Put("/tasks/:id", UpdateTask)
	app.Delete("/tasks/:id", DeleteTask)

	app.Listen(":8080")
}

// Handler to get all tasks
func GetTasks(c *fiber.Ctx) error {
	var tasks []Task
	DB.Find(&tasks)
	return c.JSON(tasks)
}

// Handler to create a new task
func CreateTask(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if !isValidPriority(task.Priority) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid priority value"})
	}

	if task.Status == "" {
		task.Status = Pending
	}

	if err := DB.Create(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create task"})
	}
	return c.Status(201).JSON(task)
}

// Handler to update a task
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task Task

	if err := DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	updatedTask := new(Task)
	if err := c.BodyParser(updatedTask); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if !isValidPriority(updatedTask.Priority) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid priority value"})
	}

	if updatedTask.Status == "" {
		updatedTask.Status = Pending
	} else if !isValidStatus(updatedTask.Status) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid status value"})
	}

	task.Task = updatedTask.Task
	task.Priority = updatedTask.Priority
	task.DueDate = updatedTask.DueDate
	task.DueTime = updatedTask.DueTime
	task.Status = updatedTask.Status

	if err := DB.Save(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update task"})
	}

	return c.JSON(task)
}

// Handler to delete a task
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task Task

	if err := DB.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Task not found"})
	}

	DB.Delete(&task)
	return c.SendStatus(204)
}

func isValidPriority(p Priority) bool {
	switch p {
	case Low, Medium, High:
		return true
	}
	return false
}

func isValidStatus(s Status) bool {
	switch s {
	case Pending, InProgress, Finished, Cancelled:
		return true
	}
	return false
}
