import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'

function Header() {
  const [darkMode, setDarkMode] = useState(false)
  const [lang, setLang] = useState('VI')

  useEffect(() => {
    const savedTheme = localStorage.getItem('theme') || 'dark'
    setDarkMode(savedTheme === 'dark')
    document.documentElement.setAttribute('data-theme', savedTheme)
  }, [])

  const toggleTheme = () => {
    const newTheme = darkMode ? 'light' : 'dark'
    setDarkMode(!darkMode)
    localStorage.setItem('theme', newTheme)
    document.documentElement.setAttribute('data-theme', newTheme)
  }

  const toggleLang = () => {
    setLang(lang === 'VI' ? 'EN' : 'VI')
  }

  return (
    <header className="app-bar" role="banner">
      <div className="container app-bar-container">
        <Link to="/" className="app-bar-brand">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
            <path d="M12 2L2 7l10 5 10-5-10-5z"/>
            <path d="M2 17l10 5 10-5"/>
            <path d="M2 12l10 5 10-5"/>
          </svg>
          <span>Luyện Phỏng Vấn IT</span>
        </Link>

        <nav className="app-bar-actions" aria-label="Điều hướng chính">
          <Link to="/pricing" className="app-bar-btn" title="Gói Pro">
            💎 Bảng giá
          </Link>
          
          <span className="app-bar-divider"></span>
          
          <button 
            className="app-bar-icon-btn" 
            onClick={toggleTheme}
            title={darkMode ? 'Chế độ sáng' : 'Chế độ tối'}
          >
            {darkMode ? '☀️' : '🌙'}
          </button>
          
          <button 
            className="app-bar-icon-btn" 
            onClick={toggleLang}
            title="English"
          >
            🌐 {lang}
          </button>
          
          <span className="app-bar-divider"></span>
          
          <Link to="/login" className="btn btn-primary btn-sm">
            Đăng nhập
          </Link>
        </nav>
      </div>
    </header>
  )
}

export default Header
