# GoMap App
# 1. Purpose 
* Take in command line inputs
* Read in a section of a remote file
* Sort and return X largest values
* Dockerize the application

It's possible to achive the above purposes without using the database, however, it means we need to store all the data in the memory which is not scalble for a large input file. Therefore, I decided to import the data into a Postgres database.  

The basic flow of this application is as follows. 
1. Read the parameters (`--url`, `--num`)
2. Download the file from the given `url`
3. Remove partial of the file from beginning (default: 500 Bytes)
4. Insert the data into database
5. Get the first `num` of the records sorted by value
6. Print out the ID and value to terminal

# 2. Stack 
* Programming language   : Go 1.15
* Database               : Postgres
* Unit Test              : Go Testing package
# 3. Usage 
## (1) Run the database 
```bash
docker-compose up db
```

## (2) Set the database's IP address to $DBHOST env var
```bash
source ./set_db_host.sh
```

## (3a) Run the application in docker (recommend)
### Build the application 
```bash
docker build -t gomap_app .
```
### Run the application
```bash
docker run -it -e DBHOST=$DBHOST --net=gomap_default gomap_app --num 20 --url "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
```
### Show the help
```bash
docker run gomap_app --help
```

## (3b) Run the application in host machine
By this method (3b), we can run the app in the host machine. However, the DBHost and DBPort in `repo.go` should be correctly set. If the DB container is running in the host machine, then we can set the DBHOST to `DBHost=localhost` and DBPORT to `DBPort=5432` to run the app.
### Build the application 
```bash 
go build -o GoMap . 
```
### Run the application 
```bash
./GoMap --num 20 --url "https://bucket-ian-1.s3.amazonaws.com/data_full.txt"
```
## (3c) Run the application with docker-compose
### Run with docker-compose 
By this method (3c), we can start app and database together. It's convinient but we can't assign the arguments `--num` and `--url` in command line directly. Besides, we need to set DBHost in `repo.go` as `DBHost=db` for using this approach.

```bash
docker-compose up app 
```

# 4. Complexity

Assume the data size is N, and the input flag `--num` is K.
### Time complexity
- Insert N records: O(N)
- Sort the data set: O(NlogN)
- Read first K records: O(K)
```
Total time complexity: O(NlogN)
```
### Space complexity
- Store N records: O(N)
```
Total space complexity: O(N)
```

# 5. Test
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
go test -v ./... | grep -e PASS -e FAIL -e ok
--- PASS: TestRetrieveData (1.05s)
    --- PASS: TestRetrieveData/Case1:_Success (1.05s)
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
ok  	github.com/ianlai/GoMap	(cached)
--- PASS: TestInsertRecord (0.00s)
    --- PASS: TestInsertRecord/Case1:_Successful (0.00s)
    --- PASS: TestInsertRecord/Case2:_Failed (0.00s)
--- PASS: TestGetRecordsSortedByVal (0.00s)
    --- PASS: TestGetRecordsSortedByVal/Case1:_Successful (0.00s)
    --- PASS: TestGetRecordsSortedByVal/Case2:_Failed (0.00s)
PASS
ok  	github.com/ianlai/GoMap/data	(cached)
```