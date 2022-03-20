# GoMap App
# 1. Purpose 
## (1) Importer
* Take command line inputs
* Import a remote file 
* Remove the header of the file (default 500 bytes)
* Insert the records into database

## (2) Show Top-K result on command line
* Sort and return X largest values in logs

## (3) Serve a HTTP server 
* Provide APIs
  * /status
  * /records
  * /records/{uid}

## (4) Dockerize the application
* Use docker to create the database
* Use docker-compose to manage the containers

# 2. App Flow
It's possible to achive the above purpose of returning Top-K values without using the database which has lower time complexity, however, it means we need to store all the data in the memory which is not scalble for a large input file. 

The basic flow of this application is as follows. 
1. Read the parameters (`--url`, `--num`)
2. Download the file from the given `url`
3. Remove the header from beginning (default: 500 Bytes)
4. Insert the data into database
5. Get the first `num` of the records sorted by value
6. Print out the ID and value to terminal
7. Start the server 

# 3. Stack 
* Programming language   : Go 1.15
* Database               : Postgres
* Unit Test              : Go Testing package
# 4. Usage 
## (1) Run the database 
```bash
docker-compose up db
```
## (2a) Run the application in docker (recommend)

### Set the database's IP address to $DBHOST env var
```bash
source ./set_db_host.sh
```
### Build the application
```bash
docker build -t gomap_app .
```
### Run the application (can set the arguments)
```bash
docker run -it -e DBHOST=$DBHOST --net=gomap_default gomap_app --num 20 --url "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
```
### Show the help
```bash
docker run gomap_app --help
```

## (2b) Run the application in host machine

### Set the database's IP address to $DBHOST env var
```bash
export DBHOST=localhost
```
### Build the application 
```bash 
go build -o GoMap . 
```
### Run the application (can set the arguments)
```bash
./GoMap --num 20 --url "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
```
## (2c) Run the application with docker-compose
### Run the application (cannot set the arguments)
By this method (3c), we can start both app and database with docker-compose. It's convinient because we don't need to set the network by our own (2b needs). We don't need to set `DBHOST` env var (both 2a and 2b need) also because it is defined in `docker-compose.yml` already. However, the downside is that we can't assign the arguments `--num` and `--url` in command line directly.

```bash
docker-compose up app 
```
# 5. Complexity 

Assume the data size is N, and the input flag `--num` is K.
## (1) Time complexity
- Insert N records: O(N)
- Sort the data set: O(NlogN)
- Read first K records: O(K)
```
Total time complexity: O(NlogN)
```
## (2) Space complexity
- Store N records: O(N)
```
Total space complexity: O(N)
```

## (3) Discussion: Time complexity without DB
We can also choose to provide Top-K values without using database or sorting algorithm. If so, we basically have 2 approaches to achive the same goal.
```
Max-Heap approach: O(N + KlogN)
Min-Heap approach: O(K + (N-K)logK)
```

# 6. Test
We can run the test of handler and DAL by the following command.
```bash
go test -v ./... 
```

For increasing the readability, we can use filter the necessary information only. 
```bash
go test -v ./... | grep -e PASS -e FAIL -e ok
```

Example result: 
```
go test -v ./... --cover=True | grep -e PASS -e FAIL -e cover --color=never
--- PASS: TestRetrieveData (1.24s)
    --- PASS: TestRetrieveData/Case1:_Success (1.24s)
    --- PASS: TestRetrieveData/Case2:_Failed_-_URL_Format (0.00s)
    --- PASS: TestRetrieveData/Case3:_Failed_-_removeLength_ (0.00s)
--- PASS: TestRemovePrefixData (0.00s)
    --- PASS: TestRemovePrefixData/Case1:_Success (0.00s)
    --- PASS: TestRemovePrefixData/Case2:_Failed (0.00s)
--- PASS: TestGetLinesFromReader (0.00s)
    --- PASS: TestGetLinesFromReader/Case1:_Success (0.00s)
    --- PASS: TestGetLinesFromReader/Case2:_Failed (0.00s)
--- PASS: TestInsertRecords (0.00s)
    --- PASS: TestInsertRecords/Case1:_Success (0.00s)
    --- PASS: TestInsertRecords/Case2:_Failed (0.00s)
--- PASS: TestGetTopKRecords (0.00s)
    --- PASS: TestGetTopKRecords/Case1:_Success (0.00s)
    --- PASS: TestGetTopKRecords/Case2:_Failed (0.00s)
PASS
coverage: 53.6% of statements
ok  	github.com/ianlai/GoMap	1.966s	coverage: 53.6% of statements
--- PASS: TestHandleShowRecord (0.00s)
    --- PASS: TestHandleShowRecord/Case1:_Success_(Found) (0.00s)
    --- PASS: TestHandleShowRecord/Case2:_Failed_(Not_found) (0.00s)
PASS
coverage: 48.9% of statements
ok  	github.com/ianlai/GoMap/app	(cached)	coverage: 48.9% of statements
--- PASS: TestInsertRecord (0.00s)
    --- PASS: TestInsertRecord/Case1:_Successful (0.00s)
    --- PASS: TestInsertRecord/Case2:_Failed (0.00s)
--- PASS: TestListRecords (0.00s)
    --- PASS: TestListRecords/Case1:_Successful_(sorted_by_val) (0.00s)
    --- PASS: TestListRecords/Case2:_Successful_(unsorted) (0.00s)
    --- PASS: TestListRecords/Case3:_Failed (0.00s)
PASS
coverage: 46.8% of statements
ok  	github.com/ianlai/GoMap/data	(cached)	coverage: 46.8% of statements
```