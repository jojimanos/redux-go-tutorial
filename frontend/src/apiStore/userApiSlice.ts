// src/features/apiSlice.ts
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

// Define the types for the data you're fetching
export interface User {
  username: number;
  email: string;
}

// Define the API slice
export const userApiSlice = createApi({
  reducerPath: "api", // Optional: name for the reducer path
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:8000" }), // Base URL for the API
  endpoints: (builder) => ({
    login: builder.mutation<{token: string}, {username: string, password: string}>({
      query: (args) => {
        return {
          url: "/user/login",
          method: "POST",
          body: {username: args.username, password: args.password}
        }
      }
    }), 
    getUser: builder.query<User, { username: string }>({
      query: (args) => {
        return {
          url: "/user",
          method: "GET",
          params: { username: args.username },
        };
      },
    }),
    getCurrentUser: builder.query<User, { token: string }>({
      query: (args) => {
        return {
          url: "/user/current",
          method: "GET",
          headers: {Authorization: `Bearer ${args.token}`}
        };
      },
    }),
    getAllUsers: builder.query<User[], void>({
      query: () => "api/users", // Replace with your API endpoint
    }),
  }),
});

// Export the generated hooks
export const {
  useLoginMutation,
  useGetUserQuery,
  useLazyGetUserQuery,
  useGetAllUsersQuery,
  useLazyGetAllUsersQuery,
  useGetCurrentUserQuery,
  useLazyGetCurrentUserQuery
} = userApiSlice;
