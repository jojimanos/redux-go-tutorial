import { createSlice } from '@reduxjs/toolkit'
import { UserType } from '../types/types';

interface UserSliceInterface {
    user: UserType | null;
    accessToken: string | undefined;
}

const initialState: UserSliceInterface = {
  user: null,
  accessToken: ''
}

export const userSlice = createSlice({
  name: 'user',
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
   setUser(state, {payload}) {
    state.user = payload
   },
   login(state, {payload}) {
    state.accessToken = payload
   },
   logout(state) {
    state.user = null;
    state.accessToken = ''
   }   
  },
})

export const userActions = userSlice.actions

// Other code such as selectors can use the imported `RootState` type
export const getCurrentUser = (state: UserSliceInterface) => state.user
export const getAccessToken = (state: UserSliceInterface) => state.accessToken

export default userSlice.reducer
