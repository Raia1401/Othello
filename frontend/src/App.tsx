import './App.css';
import { BrowserRouter,Routes, Route } from 'react-router-dom';

import { BoardDataProvider } from './providers/boardDataProvider';
import GamePage from './pages/gamePage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* <Route path='/' element={
          <BoardDataProvider>
            <LoginPage />
          </BoardDataProvider>
        } /> */}

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
