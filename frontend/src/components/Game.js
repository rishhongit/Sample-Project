import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { setDeck, setUsername, setStatus } from '../store/gameSlice';

const Game = () => {
  const dispatch = useDispatch();
  const { deck, username } = useSelector(state => state.game);
  const [drawnCard, setDrawnCard] = useState(null);

  const startGame = () => {
    fetch('/start', {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      body: new URLSearchParams({ username }),
    })
      .then(res => res.json())
      .then(data => dispatch(setDeck(data.deck)));
  };

  const drawCard = () => {
    fetch('/draw', {
      method: 'POST',
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      body: new URLSearchParams({ username }),
    })
      .then(res => res.json())
      .then(data => setDrawnCard(data.drawnCard));
  };

  return (
    <div>
      <h1>Welcome {username}</h1>
      <button onClick={startGame}>Start Game</button>
      <button onClick={drawCard}>Draw Card</button>
      {drawnCard && <p>You drew: {drawnCard}</p>}
    </div>
  );
};

export default Game;
