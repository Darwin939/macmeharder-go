// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type App struct {
	ID            string     `json:"id"`
	Created       string     `json:"created"`
	Title         string     `json:"title"`
	Language      string     `json:"language"`
	LanguageCount string     `json:"language_count"`
	Description   string     `json:"description"`
	Size          float64    `json:"size"`
	Award         string     `json:"award"`
	Place         int        `json:"place"`
	Age           int        `json:"age"`
	Developer     string     `json:"developer"`
	Chart         string     `json:"chart"`
	Version       string     `json:"version"`
	Compatibility string     `json:"compatibility"`
	DownloadURL   string     `json:"downloadUrl"`
	Category      *Category  `json:"category"`
	Avatar        *AppAvatar `json:"avatar"`
}

type AppAvatar struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	URL    string `json:"url"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Apps []*App `json:"apps"`
}

type NewApp struct {
	Title         string  `json:"title"`
	Language      string  `json:"language"`
	LanguageCount string  `json:"language_count"`
	Description   string  `json:"description"`
	Size          float64 `json:"size"`
	Award         string  `json:"award"`
	Place         int     `json:"place"`
	Age           int     `json:"age"`
	Developer     string  `json:"developer"`
	Chart         string  `json:"chart"`
	Version       string  `json:"version"`
	Compatibility string  `json:"compatibility"`
	DownloadURL   string  `json:"downloadUrl"`
	CategoryID    string  `json:"category_id"`
}

type NewImage struct {
	ID string `json:"id"`
}
