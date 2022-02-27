# GoMap App
## 1. Purpose 
* Take in command line inputs
* Read in a section of a remote file
* Sort and return X largest values
* Dockerize the application

It's possible to achive the above purposes without using the database, however, it means we need to store all the data in the memory which is not scalble for a large input file. Therefore, I decided to import the data into a Postgres database.  

The basic flow of this application is as follows. 
1. Read the parameters (`url`, `num`)
2. Download the file from the given `url`
3. Remove partial of the file from beginning (default: 500 Bytes)
4. Insert the data into database.
5. Get the first `num` of the records sorted by value 
6. Print out the ID and value to terminal 

## 2. Stack 
* Programming language   : Go 1.15
* Database               : Postgres
* Unit Test              : Go Testing package
## 3. Usage 
### (1) Run the database 
```bash
docker-compose up db
```
### (2a) Run the application (in host machine)

#### Build the application 
```bash 
go build -o GoMap .
```

#### Run the application 
```bash
./GoMap --num 20 --url "https://amp-technical-challenge.s3.ap-northeast-1.amazonaws.com/sw-engineer-challenge.txt"
```
### (2b) Run the application (in docker)

#### Build the application 
```bash
docker build -t gomap_app .
```

#### Run the application
```bash
docker run -it --net=gomap_default gomap_app --num 20 --url "https://amp-technical-challenge.s3.ap-northeast-1.amazonaws.com/sw-engineer-challenge.txt"
```

### (2c) Run the application (with docker-compose)
#### Run with docker-compose 
By this way, we can start app and database together. It's convinient but we can't assign the `num` and `url` in command line directly. Besides, we need to change the `DBHost` in the code to `DBHost=db` first for using this approach.

```bash
docker-compose up app 
```

## 4. Test
We can run the test of handler and DAL by the following command.
```bash
go test -v ./... 
```