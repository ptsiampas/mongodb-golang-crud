# Golan MongoDB Driver Example
This is reference code for the article at [Wired Elf](https://wiredelf.com/....)

It's a Simple website that explores the basic use of the official MongoDB Driver using Go.

To use this example, clone the repository and then move into the cloned directory.

## Setup
To get a simple mongo database setup and running, just use docker-compose to launch the admin front end and a secure version of mongo db.

```console
petert@dogmeat:~/go/src/mongodb-golang-crud $ docker-compose up -d
petert@dogmeat:~/go/src/mongodb-golang-crud $ docker exec -it mongodb_mongo_1 bash
# mongo -u "root" --authenticationDatabase "admin" -p 1234
root@7efbbcbe89b4:/# mongo -u "root" --authenticationDatabase "admin" -p 1234
MongoDB shell version v3.6.4
connecting to: mongodb://127.0.0.1:27017
MongoDB server version: 3.6.4
> db.books.insertMany([
  {isbn:"978-1503261969",title:"Emma",author:"Jayne Austen",price:9.44},
  {isbn:"978-1505255607",title:"The Time Machine",author:"H. G. Wells",price:5.99},
  {isbn:"978-1503379640",title:"The Prince",author:"Niccolò Machiavelli",price:6.99},
  {isbn:"978-0062457714",title:"The Subtle Art of Not Giving a F*ck",author:"Mark Manson",price:32.28},
  {isbn:"978-0730324218",title:"The Barefoot Investor 2019 Update: The Only Money Guide You'll Ever Need",author:"Scott Pape",price:18.98}
])
```
This will get you setup with dummy data to play with and examine.

