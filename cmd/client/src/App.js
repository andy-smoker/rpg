import React from 'react';
import './App.css';
import { DGgeneral } from './components/DeltaGreen';
import General from './components/Content'
import { Header } from './components/Header';
import { SideBar } from './components/SideBar';
import { Footer } from './components/Footer';
import { BrowserRouter, Route} from 'react-router-dom';


let isAuth = false

function App() {
  
  return (
    <BrowserRouter>
      <div className='main_page'>
        <Header />
        <SideBar />
        <div className='content'>
          <Route exact path='/general' component={General} />
          <Route path='/dg' component={DGgeneral} />
        </div>
        <Footer />
        <div>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;
