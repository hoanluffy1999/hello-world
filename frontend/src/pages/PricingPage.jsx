import { Link } from 'react-router-dom'
import Header from '../components/Header'

const pricingPlans = [
  {
    name: 'Free',
    price: '0đ',
    period: 'miễn phí mãi mãi',
    features: [
      '✅ Truy cập 500+ câu hỏi',
      '✅ Cơ bản HTML, CSS, JavaScript',
      '✅ Flashcards cơ bản',
      '❌ Không có tiến độ học',
      '❌ Không có giải thích chi tiết',
    ],
    cta: 'Bắt đầu miễn phí',
    popular: false,
  },
  {
    name: 'Pro',
    price: '99.000đ',
    period: '/tháng',
    features: [
      '✅ Truy cập 2000+ câu hỏi',
      '✅ Tất cả chủ đề (Frontend, Backend, Mobile...)',
      '✅ Flashcards nâng cao',
      '✅ Theo dõi tiến độ học',
      '✅ Giải thích chi tiết',
      '✅ Chế độ luyện thi',
      '✅ Hỗ trợ ưu tiên',
    ],
    cta: 'Đăng ký Pro',
    popular: true,
  },
  {
    name: 'Lifetime',
    price: '990.000đ',
    period: 'thanh toán một lần',
    features: [
      '✅ Tất cả tính năng Pro',
      '✅ Truy cập trọn đời',
      '✅ Cập nhật miễn phí',
      '✅ Ưu tiên hỗ trợ 24/7',
      '✅ Certificate hoàn thành',
    ],
    cta: 'Mua trọn đời',
    popular: false,
  },
]

function PricingPage() {
  return (
    <div className="app">
      <Header />
      
      <main className="pricing-main">
        <div className="container">
          <div className="pricing-header">
            <h1>Chọn gói phù hợp với bạn</h1>
            <p>Đầu tư cho sự nghiệp developer của bạn</p>
          </div>

          <div className="pricing-grid">
            {pricingPlans.map((plan) => (
              <div 
                key={plan.name} 
                className={`pricing-card card ${plan.popular ? 'popular' : ''}`}
              >
                {plan.popular && <span className="popular-badge">Phổ biến nhất</span>}
                
                <h3 className="plan-name">{plan.name}</h3>
                <div className="plan-price">
                  <span className="price">{plan.price}</span>
                  <span className="period">{plan.period}</span>
                </div>
                
                <ul className="plan-features">
                  {plan.features.map((feature, i) => (
                    <li key={i}>{feature}</li>
                  ))}
                </ul>
                
                <button className={`btn ${plan.popular ? 'btn-primary' : 'btn-secondary'} btn-lg`}>
                  {plan.cta}
                </button>
              </div>
            ))}
          </div>

          <div className="pricing-faq">
            <h2>Câu hỏi thường gặp</h2>
            <div className="faq-list">
              <div className="faq-item card">
                <h3>Tôi có thể hủy bất cứ lúc nào?</h3>
                <p>Có, bạn có thể hủy gói Pro bất cứ lúc nào. Tài khoản sẽ vẫn hoạt động đến hết chu kỳ thanh toán.</p>
              </div>
              <div className="faq-item card">
                <h3>Có đảm bảo hoàn tiền không?</h3>
                <p>Chúng tôi hoàn tiền 100% trong vòng 7 ngày nếu bạn không hài lòng.</p>
              </div>
              <div className="faq-item card">
                <h3>Nội dung có được cập nhật thường xuyên?</h3>
                <p>Có, chúng tôi cập nhật câu hỏi mới hàng tuần dựa trên trend phỏng vấn thực tế.</p>
              </div>
            </div>
          </div>
        </div>
      </main>

      <footer className="app-footer">
        <div className="container">
          <p>© 2026 Luyện Phỏng Vấn IT. Built with ❤️ for Vietnamese developers.</p>
        </div>
      </footer>
    </div>
  )
}

export default PricingPage
