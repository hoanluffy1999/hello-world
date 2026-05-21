import { useState, useEffect } from 'react'
import { useParams, Link } from 'react-router-dom'
import Header from '../components/Header'
import { api } from '../services/api'

function QuestionPage() {
  const { id } = useParams()
  const [question, setQuestion] = useState(null)
  const [answers, setAnswers] = useState([])
  const [loading, setLoading] = useState(true)
  const [selectedAnswer, setSelectedAnswer] = useState(null)
  const [showResult, setShowResult] = useState(false)

  useEffect(() => {
    loadQuestion()
  }, [id])

  async function loadQuestion() {
    setLoading(true)
    try {
      const data = await api.getQuestion(id)
      setQuestion(data.question)
      setAnswers(data.answers || [])
    } catch (error) {
      console.error('Failed to load question:', error)
    } finally {
      setLoading(false)
    }
  }

  function handleSelectAnswer(answerId) {
    if (showResult) return
    setSelectedAnswer(answerId)
  }

  function checkAnswer() {
    if (!selectedAnswer) return
    setShowResult(true)
    // TODO: Save progress to backend
  }

  function resetQuestion() {
    setSelectedAnswer(null)
    setShowResult(false)
  }

  if (loading) {
    return (
      <div className="app">
        <Header />
        <main className="question-main">
          <div className="container">
            <div className="loading">Đang tải câu hỏi...</div>
          </div>
        </main>
      </div>
    )
  }

  if (!question) {
    return (
      <div className="app">
        <Header />
        <main className="question-main">
          <div className="container">
            <div className="error">Không tìm thấy câu hỏi</div>
          </div>
        </main>
      </div>
    )
  }

  return (
    <div className="app">
      <Header />
      
      <main className="question-main">
        <div className="container">
          <div className="question-container">
            <Link to="/c/all" className="back-link">← Quay lại</Link>
            
            <div className="question-card card">
              <div className="question-header">
                <span className={`difficulty-badge difficulty-${question.difficulty}`}>
                  {question.difficulty}
                </span>
                {question.category && (
                  <span className="category-tag">{question.category.name}</span>
                )}
              </div>
              
              <h1 className="question-title">{question.title}</h1>
              
              {question.content && (
                <div className="question-content">
                  {question.content}
                </div>
              )}
            </div>

            <div className="answers-section">
              <h2>Chọn đáp án đúng:</h2>
              
              <div className="answers-list">
                {answers.map((answer) => (
                  <button
                    key={answer.id}
                    className={`answer-card card ${
                      selectedAnswer === answer.id ? 'selected' : ''
                    } ${
                      showResult && answer.is_correct ? 'correct' : ''
                    } ${
                      showResult && selectedAnswer === answer.id && !answer.is_correct ? 'incorrect' : ''
                    }`}
                    onClick={() => handleSelectAnswer(answer.id)}
                    disabled={showResult}
                  >
                    <div className="answer-content">{answer.content}</div>
                  </button>
                ))}
              </div>

              {!showResult ? (
                <button
                  className="btn btn-primary btn-lg"
                  onClick={checkAnswer}
                  disabled={!selectedAnswer}
                >
                  Kiểm tra đáp án
                </button>
              ) : (
                <div className="result-section">
                  {answers.find(a => a.id === selectedAnswer)?.is_correct ? (
                    <div className="result-correct">✅ Chính xác!</div>
                  ) : (
                    <div className="result-incorrect">❌ Sai rồi!</div>
                  )}
                  
                  {answers.find(a => a.is_correct)?.explanation && (
                    <div className="explanation">
                      <strong>Giải thích:</strong>
                      <p>{answers.find(a => a.is_correct).explanation}</p>
                    </div>
                  )}
                  
                  <button className="btn btn-secondary btn-lg" onClick={resetQuestion}>
                    Làm lại
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default QuestionPage
