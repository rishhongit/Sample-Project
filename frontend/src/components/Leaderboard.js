import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { setLeaderboard } from '../store/gameSlice';

const Leaderboard = () => {
  const dispatch = useDispatch();
  const leaderboard = useSelector(state => state.game.leaderboard);

  useEffect(() => {
    fetch('/leaderboard')
      .then(res => res.json())
      .then(data => dispatch(setLeaderboard(data)));
  }, [dispatch]);

  return (
    <div>
      <h2>Leaderboard</h2>
      <ul>
        {Object.keys(leaderboard).map(user => (
          <li key={user}>{user}: {leaderboard[user]}</li>
        ))}
      </ul>
    </div>
  );
};

export default Leaderboard;
