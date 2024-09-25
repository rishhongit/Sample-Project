// REDUX SLICE

import { createSlice } from '@reduxjs/toolkit';

const initialState = {
  username: '',
  deck: [],
  leaderboard: [],
  status: 'idle',
};

const gameSlice = createSlice({
  name: 'game',
  initialState,
  reducers: {
    setUsername: (state, action) => {
      state.username = action.payload;
    },
    setDeck: (state, action) => {
      state.deck = action.payload;
    },
    setLeaderboard: (state, action) => {
      state.leaderboard = action.payload;
    },
    setStatus: (state, action) => {
      state.status = action.payload;
    },
  },
});

export const { setUsername, setDeck, setLeaderboard, setStatus } = gameSlice.actions;
export default gameSlice.reducer;
