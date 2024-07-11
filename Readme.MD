It's the ~~90's~~ 20's, and B-b-blockbusta is back in business! They have big plans to reopen their chain of stores and offer those sweet tapes to the public again! They need a rental system again, and now built on containers!

# About
This is a demonstrator project to show off a number of Golang and DevOps features:

| Golang | Devops | | | |
| --- | --- | --- | --- | --- |
| cobra | http | client | server | json |
| post | get | chi | package | |
| container | API | GitHub | make | |


# MVP
As a minimum viable product, the following features needs to be implemented:

Infrastructure:
 - [x] Containerised backend API
 - [x] Web frontend containers that can be distributed to stores
 - [x] Dockerbuild 

Development:
 - [x] GitHub hosted
 - [ ] makefile to build
 - [x] Golang

Features:
 - [ ] Customer can log in and browse and borrow movies
 - [ ] Employee can log in and manage move db
 - [ ] Optional SQL database backend

 # How to use
 ## Commandline
 You can either build or start the program directly using `go run`:
 ```
 go run . apiserver
 ```
 To start the client:
 ```
 go run . client
 ```
 ## Docker
 ### Build Docker image
 ```
$ docker build -t bbuster:latest .
[1/2] STEP 1/4: FROM golang:1.22 AS builder
[1/2] STEP 2/4: WORKDIR /app
--> Using cache b755d48b3312b0cdb294fa7d8c7b8bc4801993ca9fd80812c63d1ae8fd617d0c
--> b755d48b3312
[1/2] STEP 3/4: COPY . /app/
--> d7b0913e7150
[1/2] STEP 4/4: RUN CGO_ENABLED=0 GOOS=linux go build -o bbuster .
go: downloading github.com/spf13/cobra v1.8.1
go: downloading github.com/go-chi/chi/v5 v5.1.0
go: downloading github.com/spf13/pflag v1.0.5
--> 1210a1e447e0
[2/2] STEP 1/4: FROM scratch
[2/2] STEP 2/4: COPY --from=builder /app/bbuster /app/bbuster
--> 06d7a69f4778
[2/2] STEP 3/4: EXPOSE 8080
--> 55ed8f43fa16
[2/2] STEP 4/4: ENTRYPOINT ["/app/bbuster","apiserver"]
[2/2] COMMIT bbuster:latest
--> 7384d6870ad6
Successfully tagged localhost/bbuster:latest
7384d6870ad637960cd4644c2bb5ed3c05c4442ef331e6935b8dff88f248c905
 ```
### Create container network
```
$ podman network create videorental
```

### Start server
```
$  podman run -d -p 8080:8080 --rm --net videorental --name apiserver  bbuster:latest apiserver
```
### Start client
```
$ podman run --rm -p 50080:50080  --rm --net videorental --name client bbuster:latest client -s apiserver
```
## Log in
Open your browser and browse to http://localhost:50080 