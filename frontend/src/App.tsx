import React from 'react';
import logo from './logo.svg';
import './App.css';
import { useGetUserQuery, User } from './apiStore/userApiSlice';
import { useAppSelector } from './store/hooks';
import { useDispatch } from 'react-redux';
import { userSlice } from './stateSlices/userSlice';
import { RouterProvider } from 'react-router-dom';
import routesConfig, { setupBrowserRouter } from './router/Routing';

function App() {

const {data: getUserData, error: getUserError, isLoading: getUserIsLoading} = useGetUserQuery({username: "joji"})
const dispatch = useDispatch()

// const user: User = dispatch(userSlice.actions.setUser(getUserData)).payload

// if (!user) {
  // return null
// }


console.log(getUserData)

  return (
    <div className="App">
      <RouterProvider router={setupBrowserRouter(routesConfig)}/>
    </div>
  );
}

export default App;
