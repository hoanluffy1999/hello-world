import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { api } from '../services/api'
import Header from '../components/Header'

const categories = [
  { name: 'HTML', slug: 'html', icon: '📄', count: 50 },
  { name: 'CSS', slug: 'css', icon: '🎨', count: 89 },
  { name: 'JavaScript', slug: 'javascript', icon: '⚡', count: 107 },
  { name: 'TypeScript', slug: 'typescript', icon: '📘', count: 59 },
  { name: 'React', slug: 'react', icon: '⚛️', count: 140 },
  { name: 'Next.js', slug: 'nextjs', icon: '▲', count: 50 },
  { name: 'Vue.js', slug: 'vuejs', icon: '💚', count: 56 },
  { name: 'Angular', slug: 'angular', icon: '🅰️', count: 60 },
  { name: 'Node.js', slug: 'nodejs', icon: '📦', count: 64 },
  { name: 'NestJS', slug: 'nestjs', icon: '🔴', count: 51 },
  { name: 'Python', slug: 'python', icon: '🐍', count: 55 },
  { name: 'Golang', slug: 'golang', icon: '🔵', count: 58 },
  { name: 'Java', slug: 'java', icon: '☕', count: 88 },
  { name: 'PHP', slug: 'php', icon: '🐘', count: 34 },
  { name: 'Laravel', slug: 'laravel', icon: '❤️', count: 37 },
  { name: 'System Design', slug: 'system-design', icon: '🏗️', count: 44 },
  { name: 'Database', slug: 'database', icon: '🗄️', count: 51 },
  { name: 'DevOps', slug: 'devops', icon: '🔄', count: 44 },
]

function HomePage() {
  const [stats, setStats] = useState({ totalQ: 2170, totalTopics: 42 })

  return (
    <div className="app">
      <Header />
      
      <main className="home-main">
        {/* Hero Section */}
        <section className="home-hero">
          <div className="container">
            <div className="home-hero-copy">
              <span className="home-kicker">
                🎯 Luyện phỏng vấn IT · dành cho dev Việt
              </span>
              <h1 className="home-title">
                Ôn đúng trọng tâm, nhớ lâu hơn
              </h1>
              <p className="home-hero-sub">
                2.000+ câu hỏi theo cấp độ, kèm flashcards và tiến độ tự lưu cho từng buổi luyện.
              </p>
              <div className="home-hero-stats">
                <span><strong>{stats.totalQ.toLocaleString()}</strong> Câu hỏi</span>
                <span><strong>{stats.totalTopics}</strong> Chủ đề</span>
                <span><strong>1,005+</strong> Dev tin dùng</span>
              </div>
              <div className="home-hero-actions">
                <Link to="/c/all" className="btn btn-primary">
                  📖 Xem câu hỏi
                </Link>
                <Link to="/pricing" className="btn btn-secondary">
                  💎 Xem gói
                </Link>
              </div>
            </div>
          </div>
        </section>

        {/* Categories Section */}
        <section className="home-directory">
          <div className="container">
            <div className="home-directory-copy">
              <span className="home-section-eyebrow">Chủ đề luyện</span>
              <h2>Chọn stack, luyện đúng trọng tâm</h2>
              <p>Lọc theo cấp độ, ôn bằng flashcards. Tiến độ được lưu tự động.</p>
            </div>
            
            <div className="tech-list">
              {categories.map((cat) => (
                <Link 
                  key={cat.slug} 
                  to={`/c/${cat.slug}`}
                  className="tech-pill"
                >
                  <span className="tech-icon">{cat.icon}</span>
                  <span className="tech-name">{cat.name}</span>
                  <span className="tech-count">({cat.count})</span>
                </Link>
              ))}
            </div>
          </div>
        </section>
      </main>

      <footer className="app-footer">
        <div className="container">
          <p>© 2026 Luyện Phỏng Vấn IT. Built with ❤️ for Vietnamese developers.</p>
        </div>
      </footer>
    </div>
  )
}

export default HomePage
