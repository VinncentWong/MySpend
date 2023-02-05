# MySpend
> MySpend is a web-based application used to manage one's finances. To manage these finances, MySpend applies the 50/30/20 rule where 50% is for daily needs, 30% for personal expenses, and 20% for savings. This application is actually expected to be integrated with the person's bank account in real time but due to resource constraints, the bank account as well as the entire total money owned in each bank has been set up custom.
# ~~FrontEnd[NOT READY YET]~~
> *Front End implementation is not ready. Will continue soon!*
# BackEnd
### How to Install and Run the project
- *Make sure you have installed [Docker](https://www.docker.com/) and [Docker-compose](https://docs.docker.com/compose/install/)*
- *Clone this repository*
- *Type command `cd be`*
- *Make `.env` file on your local machine and copy all props from `.env.example` file into your `.env` file*
- *Fill in the value of each empty property in the .env file*
- *After filled all your new env properties, type command*
> `docker build -t api-myspend:latest --build-arg host=[VALUE] --build-arg user=[VALUE] --build-arg password=[VALUE] --build-arg port=[VALUE] --build-arg dbname=[VALUE] --build-arg secret_key=[VALUE]  --build-arg minio_endpoint=[VALUE]  --build-arg minio_user=[VALUE] --build-arg minio_password=[VALUE] --build-arg redis_address=[VALUE] --build-arg redis_server=[VALUE] .`
- *Then type command `docker-compose up -d`*
- *Then you can type `docker container logs api`*
- *If the output of the log is ended with this sentence*
>[GIN-debug] Listening and serving HTTP on :8000
- *You have successfully run the app!*
### How to use this project
- *Make sure you've run all applications and their dependencies*
- *Don't forget to install [Postman](https://www.postman.com/) or similar applications*
- *Please open my [API documentation](https://documenter.getpostman.com/view/19666540/2s935oKiDA). All the information about the endpoint needed along with additional information about the endpoint is in the postman documentation.*
- *Have fun with my project!*
### Tech stack used
- *[Go Programming Language](https://go.dev/)*
- *[Postgresql](https://www.postgresql.org/)*
- *[GORM](https://gorm.io/)* 
- *[Docker](https://www.docker.com/)* 
- *[JSON Web Token](https://jwt.io/)*
- *[MinIO](https://min.io/)*
- *[Redis](https://redis.io/)*
- *[Go-redis](https://github.com/redis/go-redis)*
- *[Go-redis-cache](https://github.com/go-redis/cache)*
- *[Testify](https://github.com/stretchr/testify)* 
### Additional Notes
> This project also already provide unit testing on 3 repositories. You can check it on app directory, happy coding :star_struck:
