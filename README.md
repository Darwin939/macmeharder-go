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

<h3>GraphQL mutations and queries</h3>

<h4>Upload Avatar</h4>
<code>mutation($file:Upload!) {
  uploadAppAvatar(file:$file){
    url
  }
}</code>

<H1 style="color:red">HOW TO RUN</H1>
<p>The first step is Running docker container with postgreSQL image with command below:</p>
<code>docker compose up</code>
<p>run server with command in root directory:</p>
<code>go run server.go</code>