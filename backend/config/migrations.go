package config

import (
	"database/sql"
	"log"
)

func runMigrations(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			full_name VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			slug VARCHAR(100) UNIQUE NOT NULL,
			description TEXT,
			icon VARCHAR(100),
			order_index INTEGER DEFAULT 0,
			parent_id INTEGER REFERENCES categories(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS questions (
			id SERIAL PRIMARY KEY,
			category_id INTEGER REFERENCES categories(id),
			title VARCHAR(500) NOT NULL,
			title_en VARCHAR(500),
			content TEXT,
			content_en TEXT,
			difficulty VARCHAR(20) DEFAULT 'medium',
			order_index INTEGER DEFAULT 0,
			is_active BOOLEAN DEFAULT true,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS answers (
			id SERIAL PRIMARY KEY,
			question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
			content TEXT NOT NULL,
			content_en TEXT,
			explanation TEXT,
			explanation_en TEXT,
			is_correct BOOLEAN DEFAULT false,
			order_index INTEGER DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		`CREATE TABLE IF NOT EXISTS user_progress (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			question_id INTEGER REFERENCES questions(id),
			is_correct BOOLEAN,
			attempts INTEGER DEFAULT 0,
			last_attempted TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(user_id, question_id)
		)`,

		`CREATE TABLE IF NOT EXISTS bookmarks (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(user_id, question_id)
		)`,

		`CREATE INDEX IF NOT EXISTS idx_questions_category ON questions(category_id)`,
		`CREATE INDEX IF NOT EXISTS idx_questions_difficulty ON questions(difficulty)`,
		`CREATE INDEX IF NOT EXISTS idx_answers_question ON answers(question_id)`,
		`CREATE INDEX IF NOT EXISTS idx_user_progress_user ON user_progress(user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_bookmarks_user ON bookmarks(user_id)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			log.Printf("Migration error: %v", err)
			return err
		}
	}

	// Seed initial categories
	seedCategories(db)

	return nil
}

func seedCategories(db *sql.DB) {
	categories := []struct {
		name        string
		slug        string
		description string
		icon        string
	}{
		{"HTML", "html", "Câu hỏi phỏng vấn HTML", "html.svg"},
		{"CSS", "css", "Câu hỏi phỏng vấn CSS", "css.svg"},
		{"JavaScript", "javascript", "Câu hỏi phỏng vấn JavaScript", "javascript.svg"},
		{"TypeScript", "typescript", "Câu hỏi phỏng vấn TypeScript", "typescript.svg"},
		{"React", "react", "Câu hỏi phỏng vấn React", "react.svg"},
		{"Next.js", "nextjs", "Câu hỏi phỏng vấn Next.js", "nextjs.svg"},
		{"Vue.js", "vuejs", "Câu hỏi phỏng vấn Vue.js", "vuejs.svg"},
		{"Angular", "angular", "Câu hỏi phỏng vấn Angular", "angular.svg"},
		{"Node.js", "nodejs", "Câu hỏi phỏng vấn Node.js", "nodejs.svg"},
		{"NestJS", "nestjs", "Câu hỏi phỏng vấn NestJS", "nestjs.svg"},
		{"Python", "python", "Câu hỏi phỏng vấn Python", "python.svg"},
		{"Golang", "golang", "Câu hỏi phỏng vấn Golang", "golang.svg"},
		{"Java", "java", "Câu hỏi phỏng vấn Java", "java.svg"},
		{"PHP", "php", "Câu hỏi phỏng vấn PHP", "php.svg"},
		{"Laravel", "laravel", "Câu hỏi phỏng vấn Laravel", "laravel.svg"},
		{"System Design", "system-design", "Câu hỏi phỏng vấn System Design", "system-design.svg"},
		{"Database", "database", "Câu hỏi phỏng vấn Database", "database.svg"},
		{"DevOps", "devops", "Câu hỏi phỏng vấn DevOps", "devops.svg"},
	}

	for _, cat := range categories {
		var exists int
		err := db.QueryRow("SELECT COUNT(*) FROM categories WHERE slug = $1", cat.slug).Scan(&exists)
		if err == nil && exists == 0 {
			db.Exec(
				"INSERT INTO categories (name, slug, description, icon) VALUES ($1, $2, $3, $4)",
				cat.name, cat.slug, cat.description, cat.icon,
			)
		}
	}
}
