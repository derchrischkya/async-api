<a name="readme-top"></a>
![golang](https://img.shields.io/badge/golang-1.16.5-blue)
![docker](https://img.shields.io/badge/docker-20.10.7-blue)
![docker-compose](https://img.shields.io/badge/docker--compose-1.29.2-blue)
![nsq](https://img.shields.io/badge/nsq-1.2.0-blue)
![nats](https://img.shields.io/badge/nats-2.6.2-blue)
![rabbitmq](https://img.shields.io/badge/rabbitmq-3.9.5-blue)
![kafka](https://img.shields.io/badge/kafka-2.8.0-blue)

<br />
<div align="center">
  <h3 align="center">ASYNC-API</h3>

  <p align="center">
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/derchrischkya/async-api"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/derchrischkya/async-api/issues">Report Bug</a>
    ·
    <a href="https://github.com/derchrischkya/async-api/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

The projects provides a simple API-Server which has the following endpoints:
- GET /api/v1/ping
- GET /api/v1/process/run
- GET /api/v1/process/state
- GET /api/v1/internal/backendRun

The enduser can call the endpoint /api/v1/process/run to start a long running process. The process is executed in the background and the enduser can check the state of the process by calling the endpoint /api/v1/process/state. The endpoint /api/v1/internal/backendRun is used by the backend message broker system to start the process.


<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With
#### API-Server
* Golang

#### Message Broker
* NSQ
* NATS
* RabbitMQ
* Kafka
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Clone the repository
Run executeable which matches your operating system in `api`
Run any docker-compose in `message-broker`
Check out the readme.md within each tool for more information

### Prerequisites

- cloned repository
- docker
- (golang)

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

0. Clone the repo
   ```sh
   git clone git@github.com:derchrischkya/async-api.git
   ```
1. Create a new branch
   ```sh
   git checkout -b feature/AmazingFeature
   ```
2. Make your changes
3. Commit your changes
   ```sh
   git commit -m 'Add some AmazingFeature'
   ```
4. Push to the branch
   ```sh
    git push origin feature/AmazingFeature
    ```
5. Run executeable (e.g Apple Silicon M1)
   ```
   ./api/darwin_arm_64/bin/api
   ```
6. Run docker-compose for message broker (e.g NSQ)
   ```
    docker-compose -f message-broker/nsq/docker-compose.yml up -d
   ```
7. Checkout readme.md for more information about the message broker
   ```
    message-broker/nsq/README.md
   ```
8. Execute API-Call to start process
   ```
    curl -X GET http://localhost:3000/api/v1/process/run
   ```
9. Check state of process by clicking on the link in the response of the previous call or by calling the endpoint (check readme.md of message-broker for more information)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] 

See the [open issues](https://github.com/derchrischkya/aws-infra/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Christoph Richter  - christoph.richter1997@gmail.com

Project Link: [https://github.com/derchrischkya/async-api](https://github.com/derchrischkya/async-api)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [GitHub Pages](https://pages.github.com)
* [Golang](https://golang.org/)
* [NSQ](https://nsq.io/)
* [NATS](https://nats.io/)
* [RabbitMQ](https://www.rabbitmq.com/)
* [Kafka](https://kafka.apache.org/)
* [Docker](https://www.docker.com/)
* [Best README Template](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
