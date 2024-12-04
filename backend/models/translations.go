package models

// TranslationEntry represents the translations for each supported language.
type TranslationEntry struct {
	En string `bson:"en" json:"en"`
	De string `bson:"de" json:"de"`
	Fr string `bson:"fr" json:"fr"`
	El string `bson:"el" json:"el"`
}

// Translations represents a collection of translation entries mapped by a unique key.
type Translations struct {
	TranslationEntry map[string]TranslationEntry `bson:"translations" json:"translations"`
}
