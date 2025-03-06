import './App.css'
import PacksCalculator from './components/PackCalculator'
import PacksCrud from './components/PacksCrud'

function App() {
  return (
    <>
      <div className="app-container">
        <div className="component">
          <PacksCalculator />
        </div>
        <div className="component">
          <PacksCrud />
        </div>
      </div>
    </>
  )
}

export default App
