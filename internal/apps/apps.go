package apps

import (
	"log"

	database "github.com/Darwin939/macmeharder-go/internal/pkg/db/postgres"
)

type AppAvatar struct {
	ID     string
	Status string
	URL    string
}

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
	Category      int
	Avatar        *AppAvatar
}

func (app App) Save()  {
	statement, err := database.Db.Prepare(`INSERT INTO "apps" ("title","language","language_count","description","size","award","place","age","developer","chart","version","compatibility", "downloadurl","category")  VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_, error := statement.Exec(app.Title, app.Language, app.LanguageCount, app.Description, app.Size, app.Award, app.Place,app.Age, app.Developer, app.Chart, app.Version,
		app.Compatibility, app.DownloadURL, app.Category)
	if error != nil {
		log.Fatal(error)
	}
	log.Print("Row inserted!")

}

func GetAll() []App{
	stmt,err := database.Db.Prepare("SELECT * from apps")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var apps []App
	for rows.Next(){
		var app App
		err := rows.Scan(&app.ID,&app.Created,&app.Title,&app.Language,&app.LanguageCount,
						&app.Description,&app.Size,&app.Award,&app.Place,&app.Age,
					&app.Developer,&app.Chart,&app.Version,&app.Compatibility,&app.DownloadURL,&app.Category)
		if err != nil {
			log.Fatal(err)
		}
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return apps
}

func (appAva AppAvatar) Save(id int) AppAvatar{
	statement,err := database.Db.Prepare(`INSERT INTO "appavatars" (status,url) VALUES ($1,$2);
										INSERT INTO "apps(avatar_id) values($3) WHERE ID=$4"`)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	_,err = statement.Exec()
}