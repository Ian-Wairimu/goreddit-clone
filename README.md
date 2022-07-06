# goreddit-clone
this is a reddit clone implemented using golang, sqlx for querying  and pq as the postgres driver
it allows a user to: 
- create a post
- delete a post 
- vote for a post 
- create comments
- vote for a comment 
- get all posts 

## utilities
- golang
- chi for html routing
- makefile to run commands easily
-  docker - to call the postgres docker image and run it
- adminer which is more of an admin panel to check the databases that i have created
- migrate with is a database migration tool
- sqlx for easier querying of data
- pq is the postgres driver am using to connect to the database
