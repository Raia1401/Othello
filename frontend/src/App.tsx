import './App.css';
import { BrowserRouter,Routes, Route } from 'react-router-dom';

import { BoardDataProvider } from './providers/boardDataProvider';

import GamePage from './pages/gamePage';
// import LoginPage from './pages/LoginPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={
                <BoardDataProvider>
                  <GamePage/>
                </BoardDataProvider>
        }/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
