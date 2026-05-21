const API_BASE_URL = 'http://localhost:8080/api'

export const api = {
  async getCategories() {
    const res = await fetch(`${API_BASE_URL}/categories`)
    if (!res.ok) throw new Error('Failed to fetch categories')
    return res.json()
  },

  async getQuestions(params = {}) {
    const queryString = new URLSearchParams(params).toString()
    const res = await fetch(`${API_BASE_URL}/questions?${queryString}`)
    if (!res.ok) throw new Error('Failed to fetch questions')
    return res.json()
  },

  async getQuestion(id) {
    const res = await fetch(`${API_BASE_URL}/questions/${id}`)
    if (!res.ok) throw new Error('Failed to fetch question')
    return res.json()
  },

  async register(data) {
    const res = await fetch(`${API_BASE_URL}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })
    if (!res.ok) throw new Error('Registration failed')
    return res.json()
  },

  async login(data) {
    const res = await fetch(`${API_BASE_URL}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    })
    if (!res.ok) throw new Error('Login failed')
    return res.json()
  },

  async saveProgress(token, data) {
    const res = await fetch(`${API_BASE_URL}/progress`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(data),
    })
    if (!res.ok) throw new Error('Failed to save progress')
    return res.json()
  },

  async getProgress(token) {
    const res = await fetch(`${API_BASE_URL}/progress`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    if (!res.ok) throw new Error('Failed to fetch progress')
    return res.json()
  },
}
