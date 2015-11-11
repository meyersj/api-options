# api-options
Really basic HTTP API implementations in Go, Node.js and Python/Flask. We can use these implementations
to help evaluate what direction to go for our product. Another consideration should be data-access libraries.
Does the language provide a nice way to work with Neo4j?

Each implementation has two endpoints available. If the `id` and `pass` are both **admin** then
the user will receive a success message.

+ Index: 127.0.0.1:500x/
+ User Authenicate: 127.0.0.1:500x/user?id=<user id>&pass=<password>

After running, each implementation can be accessed at the following endpoints

+ Python/Flask: 127.0.0.1:5000
+ Go: 127.0.0.1:5001
+ Node.js: 127.0.0.1:5002



