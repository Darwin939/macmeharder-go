package categories

import (
	"github.com/Darwin939/macmeharder-go/internal/apps"
)

type Category struct {
	ID   string 
	Name string 
	Apps []*apps.App 
}