// src/components/Login.tsx

import React, { useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import { NavLink, useNavigate, useOutletContext } from "react-router-dom";
import { useGetUserQuery, useLazyGetCurrentUserQuery, useLazyGetUserQuery, useLoginMutation } from "../apiStore/userApiSlice";
import {userActions} from '../stateSlices/userSlice'

const Login: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [errors, setErrors] = useState<{ username: boolean; password: boolean; }>();
const {token} = useOutletContext<{token: string}>()

  const dispatch = useDispatch();

  const navigate = useNavigate();

  const [
    login,
    { isLoading: loginIsLoading, isError: loginIsError, error: loginError },
  ] = useLoginMutation();

  const 
    [getCurrentUser]
   = useLazyGetCurrentUserQuery();

  const handleValidation = () => {
    let tempErrors = {username: false, password: false};
    let isValid = true;

    if (username.length <= 0) {
      tempErrors.username = true;
      isValid = false;
    }
    if (password.length <= 8) {
      tempErrors.username = true;
      isValid = false;
    }

    setErrors({...tempErrors})
    return isValid
  }

const handleSubmit = useCallback(
        async (e: React.FormEvent<HTMLFormElement>) => {
            e.preventDefault();

            // Basic validation
            if (!username || !password) {
                setError("Please fill in all fields.");
                return;
            }

            try {
                const token = await login({ username, password });
                console.log(token.data?.token);

                if (token.data?.token) {
                    window.localStorage.setItem("token", token.data.token as string);
                    dispatch(userActions.login(token.data.token));
                    const user = await getCurrentUser({token: token.data.token})
                    dispatch(userActions.setUser(user.data))
                    setError(""); // Reset error state
                } else {
                    setError("Login failed.");
                }
            } catch (error) {
                setError("An error occurred during login.");
            }
        },
        [username, password, login, dispatch]
    ); 

  useEffect(() => {
if (token)
      navigate("/profile");
  }, [token, handleSubmit, navigate])

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
        <h2 className="text-2xl font-bold text-center">Login</h2>
        {error && <p className="text-red-500 text-center">{error}</p>}
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label
              htmlFor="username"
              className="block text-sm font-medium text-gray-700"
            >
              Username
            </label>
            <input
              type="username"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
            {errors?.username && (
            <p className="text-red-500">Απαιτείται το Ονοματεπώνυμο!</p>
          )}
          </div>
          <div>
            <label
              htmlFor="password"
              className="block text-sm font-medium text-gray-700"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-1 block w-full p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              required
            />
            {errors?.password && (
            <p className="text-red-500">Απαιτείται το Ονοματεπώνυμο!</p>
          )}
          </div>
          <button
            type="submit"
            className="w-full py-2 px-4 font-semibold text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50"
          >
            Login
          </button>
        </form>
        <p className="text-center text-sm text-gray-600">
          Don't have an account?{" "}
          <NavLink to="/signup" className="text-blue-600 hover:underline">
            Sign up
          </NavLink>
        </p>
      </div>
    </div>
  );
};

export default Login;
