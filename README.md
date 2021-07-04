This is backend of the Macmeharder.com. All code have written in golang. 
PostgresSQL used as the database

<H1> How to setup</H1>

<h3> regenerate graphql schema </h3>
<code> go run github.com/99designs/gqlgen generate </code>


<h3>migrate db</h3>

<code>migrate -database postgresql://postgres:postgres@/postgres?sslmode=disable -path internal/pkg/db/migrations/postgres up</code>

<h3>or Force to specific migration version</h3>

<code>migrate -database postgresql://postgres:postgres@/postgres?sslmode=disable -path internal/pkg/db/migrations/postgres force 1</code>

<h3>Create new migration</h3>


<code>migrate create -ext sql -dir postgres -seq create_avatar_table</code>
