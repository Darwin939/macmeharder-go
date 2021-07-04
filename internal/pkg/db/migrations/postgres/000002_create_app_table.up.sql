CREATE TABLE IF NOT EXISTS Apps(
    ID SERIAL NOT NULL PRIMARY KEY,
    Created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    Title VARCHAR,
    Language VARCHAR,
    Language_count VARCHAR,
    Description VARCHAR,
    Size REAL,
    Award VARCHAR,
    Place INT,
    Age INT,
    Developer VARCHAR,
    Chart VARCHAR,
    Version VARCHAR,
    Compatibility VARCHAR,
    DownloadUrl VARCHAR,
    Category INT REFERENCES Categories(ID) 
)