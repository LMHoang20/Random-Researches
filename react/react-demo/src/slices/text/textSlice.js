import { createSlice } from '@reduxjs/toolkit'

export const textSlice = createSlice({
  name: 'text',
  initialState: {
    value: "Hello World",
  },
  reducers: {
    setText: (state, action) => {
      state.value = action.payload
    }
  },
})

export const { setText } = textSlice.actions

export const selectText = (state) => state.text.value

export default textSlice.reducer
