import { createSlice } from '@reduxjs/toolkit'
import { RootState } from '../store/store'


// Define a type for the slice state
interface UserSliceInterface {
    user: User | null;
    accessToken: string;
}

type User = {
    username: string | null,
    email: string | null
}

// Define the initial state using that type
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

export const { setUser, login, logout } = userSlice.actions

// Other code such as selectors can use the imported `RootState` type
export const getCurrentUser = (state: UserSliceInterface) => state.user
export const getAccessToken = (state: UserSliceInterface) => state.accessToken

export default userSlice.reducer