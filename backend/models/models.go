package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FullName     string    `json:"full_name,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Category struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	Description string     `json:"description,omitempty"`
	Icon        string     `json:"icon,omitempty"`
	OrderIndex  int        `json:"order_index"`
	ParentID    *int       `json:"parent_id,omitempty"`
	Questions   []Question `json:"questions,omitempty"`
}

type Question struct {
	ID          int        `json:"id"`
	CategoryID  int        `json:"category_id"`
	Title       string     `json:"title"`
	TitleEN     string     `json:"title_en,omitempty"`
	Content     string     `json:"content,omitempty"`
	ContentEN   string     `json:"content_en,omitempty"`
	Difficulty  string     `json:"difficulty"`
	OrderIndex  int        `json:"order_index"`
	IsActive    bool       `json:"is_active"`
	Answers     []Answer   `json:"answers,omitempty"`
	Category    *Category  `json:"category,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Answer struct {
	ID             int       `json:"id"`
	QuestionID     int       `json:"question_id"`
	Content        string    `json:"content"`
	ContentEN      string    `json:"content_en,omitempty"`
	Explanation    string    `json:"explanation,omitempty"`
	ExplanationEN  string    `json:"explanation_en,omitempty"`
	IsCorrect      bool      `json:"is_correct"`
	OrderIndex     int       `json:"order_index"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserProgress struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	QuestionID    int       `json:"question_id"`
	IsCorrect     bool      `json:"is_correct"`
	Attempts      int       `json:"attempts"`
	LastAttempted time.Time `json:"last_attempted"`
	CreatedAt     time.Time `json:"created_at"`
}

type Bookmark struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	QuestionID int       `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
}
