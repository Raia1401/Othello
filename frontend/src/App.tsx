import './App.css';
import GamePage from './pages/GamePage';
import { BoardDataProvider } from './providers/BoardDataProvider';
import { UserDataProvider } from './providers/UserDataProvider';

function App() {
  return (
    <UserDataProvider>
      <BoardDataProvider>
        <GamePage/>
      </BoardDataProvider>
    </UserDataProvider>
  );
}

export default App;
