# E-Partogram-Scheduler


## **Description**
`Nurse` Might missout on taking readings of a patient this scheduler reminds the `Nurse` by Periodically scanning the database and alerts the `Nurse` with a push notification if she misses out to take a `patient`'s reading.

___
## **Tech-Stack**
|<img src=https://miro.medium.com/max/920/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg height=200 width="100%" > | <img src=https://seekvectorlogo.com/wp-content/uploads/2018/12/docker-vector-logo.png height=200 width="100%" > | 
|---|---|
|<img src="https://www.freecodecamp.org/news/content/images/2020/05/unnamed-1.png" height=200 width="100%" >|<img src=https://res.cloudinary.com/practicaldev/image/fetch/s--TNgs2Fd7--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/8susytd9w6lxe9sreqvd.jpg  height=200 width="100%" >    | 

___
## **Features**

* The Scheduler Application written in `Golang` which runs a periodic cronJob to dial a `gRPC` call to [epartogram-server](https://github.com/captainirs/epartogram-server) in a `goroutine`


* Scans through patient readings in the database every 30 minutes and notifies the alloted nurse if she missed out any patient readings in last 30 minutes in real-time.

## Requirements
* [Golang](https://go.dev)
* [Docker](https://www.docker.com)

## Setup
* Create .env
  ```sh
  cp .env.example .env
  ```
* Run in Local
  ```sh
  go run main.go
  ```
* Run in Docker
  ```sh
  docker build -t "epartogram/scheduler" .
  ```
  ```sh
  docker run --network epartogram_network --name scheduler "epartogram/scheduler"
  ```