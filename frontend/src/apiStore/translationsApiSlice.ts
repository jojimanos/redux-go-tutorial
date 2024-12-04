import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {TranslationType} from '../types/types'

export const translationApiSlice = createApi({
  reducerPath: "api", // Optional: name for the reducer path
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:8000" }), // Base URL for the API
  endpoints: (builder) => ({
    getTranslations: builder.query<TranslationType, void>({
      query: (args) => {
        return {
          url: "/translation",
          method: "GET",
        }
      }
    }), 
  }),
});

// Export the generated hooks
export const {
  useGetTranslationsQuery,
} = translationApiSlice;
