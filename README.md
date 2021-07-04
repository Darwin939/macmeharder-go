This is backend of the Macmeharder.com. All code have written in golang. 
PostgresSQL used as the database

<H1> How to setup</H1>

<h3>migrate db</h3>

<code>migrate -database postgresql://postgres:postgres@/postgres?sslmode=disable -path internal/pkg/db/migrations/postgres up</code>

<h3>or Force to version</h3>

<code>migrate -database postgresql://postgres:postgres@/postgres?sslmode=disable -path internal/pkg/db/migrations/postgres force 1</code>