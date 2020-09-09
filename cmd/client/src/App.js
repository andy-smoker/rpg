import React from 'react'
import { useRoutes } from "./routes";
import {BrowserRouter as Router} from 'react-router-dom'

function App() {
  const routes = useRoutes(1)
  return (
    <div>
    <Router>
    <div className="container">
      {routes}
    </div>
    </Router>
    </div>
  )
}

export default App;
