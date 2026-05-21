import { Routes, Route } from 'react-router-dom'
import HomePage from './pages/HomePage'
import CategoryPage from './pages/CategoryPage'
import QuestionPage from './pages/QuestionPage'
import PricingPage from './pages/PricingPage'

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/c/:category" element={<CategoryPage />} />
      <Route path="/q/:id" element={<QuestionPage />} />
      <Route path="/pricing" element={<PricingPage />} />
    </Routes>
  )
}

export default App
