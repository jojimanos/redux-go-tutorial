import { createSlice } from '@reduxjs/toolkit'
import { LanguageType } from '../types/types';

interface LanguageSliceInterface {
    language: LanguageType;
}

const initialState: LanguageSliceInterface = {
    language: 'en'
}

export const languageSlice = createSlice({
  name: 'language',
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
   setLanguage(state, {payload}) {
       state.language = payload
   },
  },
})

export const { setLanguage } = languageSlice.actions

// Other code such as selectors can use the imported `RootState` type
export const getCurrentLanguage = (state: LanguageSliceInterface) => state.language

export default languageSlice.reducer
