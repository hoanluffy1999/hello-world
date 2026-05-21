import { useState, useEffect } from 'react'
import { useParams, Link } from 'react-router-dom'
import Header from '../components/Header'
import { api } from '../services/api'

const difficultyColors = {
  easy: '#10b981',
  medium: '#f59e0b',
  hard: '#ef4444',
}

function CategoryPage() {
  const { category } = useParams()
  const [questions, setQuestions] = useState([])
  const [loading, setLoading] = useState(true)
  const [total, setTotal] = useState(0)

  useEffect(() => {
    loadQuestions()
  }, [category])

  async function loadQuestions() {
    setLoading(true)
    try {
      const data = await api.getQuestions({ 
        category: category === 'all' ? undefined : category,
        limit: 50 
      })
      setQuestions(data.questions || [])
      setTotal(data.total || 0)
    } catch (error) {
      console.error('Failed to load questions:', error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app">
      <Header />
      
      <main className="category-main">
        <div className="container">
          <div className="category-header">
            <Link to="/" className="back-link">← Quay lại</Link>
            <h1>
              {category === 'all' ? 'Tất cả câu hỏi' : `Câu hỏi ${category.toUpperCase()}`}
            </h1>
            <p className="category-stats">{total} câu hỏi</p>
          </div>

          {loading ? (
            <div className="loading">Đang tải...</div>
          ) : (
            <div className="question-list">
              {questions.map((q) => (
                <Link 
                  key={q.id} 
                  to={`/q/${q.id}`}
                  className="question-card card"
                >
                  <div className="question-header">
                    <h3 className="question-title">{q.title}</h3>
                    <span 
                      className="difficulty-badge"
                      style={{ background: difficultyColors[q.difficulty] }}
                    >
                      {q.difficulty}
                    </span>
                  </div>
                  {q.category && (
                    <div className="question-meta">
                      <span className="category-tag">{q.category.name}</span>
                    </div>
                  )}
                </Link>
              ))}
            </div>
          )}
        </div>
      </main>
    </div>
  )
}

export default CategoryPage
