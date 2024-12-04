import { LanguageType, TranslationType } from "../types/types";


export function tryTranslation(value: string, translations: TranslationType, language: LanguageType) {
    return value in translations && language in translations[value] ? translations[value][language] : value;
}
