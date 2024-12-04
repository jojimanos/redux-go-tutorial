import { createSlice } from '@reduxjs/toolkit'
import { TranslationType } from '../types/types'

// Define a type for the slice state
interface TranslationSliceInterface {
    translations: TranslationType | undefined;
}

const initialState: TranslationSliceInterface = {
    translations: {}
}

export const translationSlice = createSlice({
  name: 'translation',
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
   setTranslation(state, {payload}) {
   },
  },
})

export const { setTranslation } = translationSlice.actions

// Other code such as selectors can use the imported `RootState` type
export const getCurrentTranslation = (state: TranslationSliceInterface) => state.translations

export default translationSlice.reducer
