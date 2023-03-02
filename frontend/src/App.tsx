import './App.css';
import { BrowserRouter,Routes, Route } from 'react-router-dom';

import { BoardDataProvider } from './providers/BoardDataProvider';
import { UserDataProvider } from './providers/UserDataProvider';

import GamePage from './pages/GamePage';
import LoginPage from './pages/LoginPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={
          <LoginPage />
        } />

        <Route path='/game' element={
                <UserDataProvider>
                <BoardDataProvider>
                  <GamePage/>
                </BoardDataProvider>
              </UserDataProvider>
        }/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
