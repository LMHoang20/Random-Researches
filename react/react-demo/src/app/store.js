import { configureStore } from '@reduxjs/toolkit';
import textReducer from '../slices/text/textSlice';

export default configureStore({
  reducer: {
    text: textReducer,
  },
});
