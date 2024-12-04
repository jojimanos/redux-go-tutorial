import { configureStore } from '@reduxjs/toolkit'
import { userApiSlice } from '../apiStore/userApiSlice'
import userReducer from '../stateSlices/userSlice'
// ...

export const store = configureStore({
  reducer: {
    // orders: ordersReducer,
    [userApiSlice.reducerPath]: userApiSlice.reducer,
    userSlice: userReducer

  },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(userApiSlice.middleware),
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch
