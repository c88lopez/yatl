package main

// Todo struct
type Todo struct {
	Task string `json:"task"`
	// Completed bool   `json:"completed"`
	// Due       time.Time `json:"due"`
}

// Todos type
type Todos []Todo
