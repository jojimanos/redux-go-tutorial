export type LanguageType = 'en' | 'de' | 'fr' | 'el';

export type TranslationType = {[k: string]: {en: string, de: string, fr: string, el: string}}

export type UserType = {
    username: string | null,
    email: string | null
}
