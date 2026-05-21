package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"interview-clone/models"

	"github.com/gin-gonic/gin"
)

// GetCategories returns all categories
func GetCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, name, slug, description, icon, order_index, parent_id 
			FROM categories 
			ORDER BY order_index, name
		`)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		categories := []models.Category{}
		for rows.Next() {
			var cat models.Category
			err := rows.Scan(&cat.ID, &cat.Name, &cat.Slug, &cat.Description, &cat.Icon, &cat.OrderIndex, &cat.ParentID)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			// Get question count for each category
			var count int
			db.QueryRow("SELECT COUNT(*) FROM questions WHERE category_id = $1 AND is_active = true", cat.ID).Scan(&count)
			
			c.JSON(200, gin.H{
				"categories": categories,
			})
			return
		}

		c.JSON(200, gin.H{
			"categories": categories,
		})
	}
}

// GetQuestions returns questions with filters
func GetQuestions(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Query("category")
		difficulty := c.Query("difficulty")
		limit := c.DefaultQuery("limit", "20")
		offset := c.DefaultQuery("offset", "0")

		limitInt, _ := strconv.Atoi(limit)
		offsetInt, _ := strconv.Atoi(offset)

		query := `
			SELECT q.id, q.category_id, q.title, q.title_en, q.content, q.content_en, 
			       q.difficulty, q.order_index, q.is_active, q.created_at, q.updated_at,
			       c.name as category_name, c.slug as category_slug
			FROM questions q
			LEFT JOIN categories c ON q.category_id = c.id
			WHERE q.is_active = true
		`

		args := []interface{}{}
		argCount := 1

		if category != "" {
			query += " AND c.slug = $" + strconv.Itoa(argCount)
			args = append(args, category)
			argCount++
		}

		if difficulty != "" {
			query += " AND q.difficulty = $" + strconv.Itoa(argCount)
			args = append(args, difficulty)
			argCount++
		}

		query += " ORDER BY q.order_index, q.id LIMIT $" + strconv.Itoa(argCount) + " OFFSET $" + strconv.Itoa(argCount+1)
		args = append(args, limitInt, offsetInt)

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		questions := []map[string]interface{}{}
		for rows.Next() {
			var q models.Question
			var categoryName, categorySlug string
			err := rows.Scan(
				&q.ID, &q.CategoryID, &q.Title, &q.TitleEN, &q.Content, &q.ContentEN,
				&q.Difficulty, &q.OrderIndex, &q.IsActive, &q.CreatedAt, &q.UpdatedAt,
				&categoryName, &categorySlug,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			q.Category = &models.Category{
				Name: categoryName,
				Slug: categorySlug,
			}

			questionMap := map[string]interface{}{
				"id":          q.ID,
				"category_id": q.CategoryID,
				"title":       q.Title,
				"title_en":    q.TitleEN,
				"content":     q.Content,
				"content_en":  q.ContentEN,
				"difficulty":  q.Difficulty,
				"category":    q.Category,
			}
			questions = append(questions, questionMap)
		}

		// Get total count
		var total int
		countQuery := "SELECT COUNT(*) FROM questions q LEFT JOIN categories c ON q.category_id = c.id WHERE q.is_active = true"
		if category != "" {
			countQuery += " AND c.slug = '" + category + "'"
		}
		if difficulty != "" {
			countQuery += " AND q.difficulty = '" + difficulty + "'"
		}
		db.QueryRow(countQuery).Scan(&total)

		c.JSON(200, gin.H{
			"questions": questions,
			"total":     total,
			"limit":     limitInt,
			"offset":    offsetInt,
		})
	}
}

// GetQuestion returns a single question with answers
func GetQuestion(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		var q models.Question
		var categoryName, categorySlug string
		err := db.QueryRow(`
			SELECT q.id, q.category_id, q.title, q.title_en, q.content, q.content_en, 
			       q.difficulty, q.order_index, q.is_active, q.created_at, q.updated_at,
			       c.name, c.slug
			FROM questions q
			LEFT JOIN categories c ON q.category_id = c.id
			WHERE q.id = $1
		`, id).Scan(
			&q.ID, &q.CategoryID, &q.Title, &q.TitleEN, &q.Content, &q.ContentEN,
			&q.Difficulty, &q.OrderIndex, &q.IsActive, &q.CreatedAt, &q.UpdatedAt,
			&categoryName, &categorySlug,
		)

		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "Question not found"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		q.Category = &models.Category{
			ID:   q.CategoryID,
			Name: categoryName,
			Slug: categorySlug,
		}

		// Get answers
		answerRows, err := db.Query(`
			SELECT id, question_id, content, content_en, explanation, explanation_en, 
			       is_correct, order_index, created_at
			FROM answers
			WHERE question_id = $1
			ORDER BY order_index, id
		`, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer answerRows.Close()

		answers := []models.Answer{}
		for answerRows.Next() {
			var a models.Answer
			err := answerRows.Scan(
				&a.ID, &a.QuestionID, &a.Content, &a.ContentEN, &a.Explanation, 
				&a.ExplanationEN, &a.IsCorrect, &a.OrderIndex, &a.CreatedAt,
			)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			answers = append(answers, a)
		}

		c.JSON(200, gin.H{
			"question": q,
			"answers":  answers,
		})
	}
}

// Register handles user registration
func Register(db *sql.DB) gin.HandlerFunc {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Hash password (implement proper hashing)
		hashedPassword := req.Password // TODO: Implement proper password hashing

		var userID int
		err := db.QueryRow(
			"INSERT INTO users (email, password_hash, full_name) VALUES ($1, $2, $3) RETURNING id",
			req.Email, hashedPassword, req.FullName,
		).Scan(&userID)

		if err != nil {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}

		c.JSON(201, gin.H{
			"message": "User registered successfully",
			"user_id": userID,
		})
	}
}

// Login handles user login
func Login(db *sql.DB) gin.HandlerFunc {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		var userID int
		var storedHash string
		err := db.QueryRow(
			"SELECT id, password_hash FROM users WHERE email = $1",
			req.Email,
		).Scan(&userID, &storedHash)

		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Verify password (implement proper verification)
		if storedHash != req.Password { // TODO: Implement proper password verification
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token (implement JWT)
		token := "sample-jwt-token" // TODO: Implement JWT token generation

		c.JSON(200, gin.H{
			"token": token,
			"user": gin.H{
				"id":    userID,
				"email": req.Email,
			},
		})
	}
}

// SaveProgress saves user progress
func SaveProgress(db *sql.DB) gin.HandlerFunc {
	type ProgressRequest struct {
		QuestionID int  `json:"question_id"`
		IsCorrect  bool `json:"is_correct"`
	}

	return func(c *gin.Context) {
		userID := c.GetInt("user_id") // From middleware
		var req ProgressRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		_, err := db.Exec(`
			INSERT INTO user_progress (user_id, question_id, is_correct, attempts, last_attempted)
			VALUES ($1, $2, $3, 1, NOW())
			ON CONFLICT (user_id, question_id) 
			DO UPDATE SET 
				is_correct = EXCLUDED.is_correct,
				attempts = user_progress.attempts + 1,
				last_attempted = NOW()
		`, userID, req.QuestionID, req.IsCorrect)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Progress saved"})
	}
}

// GetProgress gets user progress
func GetProgress(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		rows, err := db.Query(`
			SELECT question_id, is_correct, attempts, last_attempted
			FROM user_progress
			WHERE user_id = $1
		`, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		progress := []map[string]interface{}{}
		for rows.Next() {
			var p map[string]interface{}
			rows.Scan(&p)
			progress = append(progress, p)
		}

		c.JSON(200, gin.H{"progress": progress})
	}
}

// UpdateProfile updates user profile
func UpdateProfile(db *sql.DB) gin.HandlerFunc {
	type ProfileRequest struct {
		FullName string `json:"full_name"`
	}

	return func(c *gin.Context) {
		userID := c.GetInt("user_id")
		var req ProfileRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		_, err := db.Exec(
			"UPDATE users SET full_name = $1, updated_at = NOW() WHERE id = $2",
			req.FullName, userID,
		)

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Profile updated"})
	}
}

// JSONResponse helper
func JSONResponse(c *gin.Context, status int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	json.NewEncoder(c.Writer).Encode(data)
}
