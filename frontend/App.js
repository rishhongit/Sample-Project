import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import { setUsername } from './store/gameSlice';
import Game from './components/Game';
import Leaderboard from './components/Leaderboard';

function App() {
  const dispatch = useDispatch();
  const [user, setUser] = useState('');

  const handleLogin = () => {
    dispatch(setUsername(user));
  };

  return (
    <div>
      <input type="text" value={user} onChange={(e) => setUser(e.target.value)} placeholder="Enter username" />
      <button onClick={handleLogin}>Enter Game</button>
      <Game />
      <Leaderboard />
    </div>
  );
}

export default App;
