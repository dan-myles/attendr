# Watcher
This GO microservice is responsible for scraping class schedules and checking
if there is open seats available. 

### Database Interactions
Database connectivity is handled by the [ent](https://entgo.io/) go ORM library. However, database
schema is *not* handled by it. Currently database schema is handled by [prisma](https://www.prisma.io/ecosystem)
ORM on the front end side of the application. This was due to prisma being very
easy to work with. Database schema sometimes needs to be updated since prisma
will be pushing changes to the database schema every once in a while. 

Please be aware we are using a prisma -> ent schema generation tool called [entimport](https://github.com/ariga/entimport).
Make sure to update your schema frequently!


