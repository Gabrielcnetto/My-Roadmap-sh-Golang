package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/urfave/cli"
)

var pool *pgxpool.Pool
var ctx = context.Background()

type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InitConn() {
	var err error
	pool, err = pgxpool.New(ctx, "postgres://postgres:Netinho292929@@localhost:5432/pedido")

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping database:", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
}

func createTask(text string) error {
	sql := `
        INSERT INTO tasks (text, completed)
        VALUES ($1, $2)
        RETURNING id
    `

	var id int
	err := pool.QueryRow(ctx, sql, text, false).Scan(&id)
	if err != nil {
		return fmt.Errorf("error creating task: %w", err)
	}

	fmt.Printf("Created task with ID: %d\n", id)
	return nil
}

func GetAllTasks() ([]Task, error) {
	sql := `
        SELECT id, text, completed, created_at, updated_at
        FROM tasks
        ORDER BY created_at DESC
    `
	data, err := pool.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer data.Close()
	var tasks []Task
	for data.Next() {
		var task Task
		err := data.Scan(
			&task.ID, &task.Text, &task.Completed, &task.CreatedAt, &task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := data.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task rows: %w", err)
	}
	return tasks, nil
}
func completeTask(id int) error {
	sql := `
		UPDATE tasks SET completed = true, updated_at = NOW() WHERE id = $1
	`
	commandTag, err := pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("error completing task: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}

func deleteItem(id int) error {
	sql := `DELETE FROM tasks where id = $1`
	commandTag, err := pool.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("no delete items with id:", id)
	}
	return nil
}

func getPendingTasks() ([]Task, error) {
	sql := `
        SELECT id, text, completed, created_at, updated_at
        FROM tasks
        WHERE completed = false
        ORDER BY created_at DESC
    `

	rows, err := pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying pending tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Text,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task rows: %w", err)
	}

	return tasks, nil
}

func getCompletedTasks() ([]Task, error) {
	sql := `
        SELECT id, text, completed, created_at, updated_at
        FROM tasks
        WHERE completed = false
        ORDER BY created_at DESC
    `

	rows, err := pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying completed tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Text,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task rows: %w", err)
	}

	return tasks, nil
}

func main() {
	app := &cli.App{
		Name:  "gotodo",
		Usage: "A simple CLI program to manage your tasks",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	InitConn()
	createTask := createTask("Cortar cana")
	if createTask != nil {
		fmt.Println(createTask.Error())
	}
	tasks, err := GetAllTasks()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println("Tasks:", tasks)

	update := completeTask(1)
	if update != nil {
		fmt.Println(update.Error())
	}
	//delete := deleteItem(1)
	//if delete != nil {
	//	fmt.Println(delete.Error())
	//}
	pendingItems, err := getCompletedTasks()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n pending: %v\n", pendingItems)

}
