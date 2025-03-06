import './App.css'
import PacksCalculator from './components/PackCalculator'
import PacksCrud from './components/PacksCrud'

function App() {
  return (
    <>
      <p className="app-description">
        This is a simple React application that calculates the optimal number of packs for a given order quantity.
        See the source on <a href="https://github.com/cvele/reptask">GitHub</a>.
      </p>
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
