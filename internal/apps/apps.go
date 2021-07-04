package apps

import (
	"github.com/Darwin939/macmeharder-go/internal/categories"
)

type App struct {
	ID            string     
	Created       string     
	Title         string    
	Language      string     
	LanguageCount string   
	Description   string   
	Size          float64    
	Award         string     
	Place         int       
	Age           int        
	Developer     string     
	Chart         string    
	Version       string    
	Compatibility string    
	DownloadURL   string   
	Category      *categories.Category 
	Avatar        *AppAvatar 
}
