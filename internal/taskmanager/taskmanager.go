package taskmanager

import (
    "database/sql"
    "log"
    "time"

    _ "github.com/mattn/go-sqlite3"
)

type Task struct {
    ID          int
    Description string
    CreatedAt   time.Time
    IsComplete  bool
}


var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("sqlite3", "./tasks.db")
    if err != nil {
        log.Fatal(err)
    }

    createTable()
}

func createTable() {
    createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        created_at DATETIME NOT NULL,
        is_complete BOOLEAN NOT NULL DEFAULT 0
    );`
    _, err := db.Exec(createTableSQL)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
}

func AddTask(description string) {
    statement, err := db.Prepare("INSERT INTO tasks (description, created_at, is_complete) VALUES (?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    _, err = statement.Exec(description, time.Now(), false)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Added new task:", description)
}

func ListTasks(showAll bool) []Task {
    var query string
    if showAll {
        query = "SELECT id, description, created_at, is_complete FROM tasks"
    } else {
        query = "SELECT id, description, created_at, is_complete FROM tasks WHERE is_complete = 0"
    }

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var tasks []Task
    for rows.Next() {
        var t Task
        err = rows.Scan(&t.ID, &t.Description, &t.CreatedAt, &t.IsComplete)
        if err != nil {
            log.Fatal(err)
        }
        tasks = append(tasks, t)
    }
    return tasks
}

func CompleteTask(id int) {
    _, err := db.Exec("UPDATE tasks SET is_complete = 1 WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Completed task ID:", id)
}

func DeleteTask(id int) {
    _, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Deleted task ID:", id)
}
