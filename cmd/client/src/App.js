import React from 'react'
import { useRoutes } from "./routes";
import {BrowserRouter as Router} from 'react-router-dom'

function App() {
  let isAuth = false
  
  if (localStorage["token"] != null){
    isAuth = true
  }
  const routes = useRoutes(isAuth)
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
