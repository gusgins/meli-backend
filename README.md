<p align="center">
  <h3 align="center">Meli-Backend Exercise</h3>

  <p align="center">
    API for Magneto's Mutant Searching needs
  </p>
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
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
  </ol>
</details>


<!-- ABOUT THE PROJECT -->
## About The Project

This is an exercise to find, given a string array (DNA), if it corresponds to a mutant or not.

Considerations:
* DNA is represented by a NxN table, with only (A,T,C,G) as possible values
* A DNA corresponds to a mutant if and only if more than one sequence of 4 equal characters is found in the table, this can be either horizontal, vertical or diagonal

The requirements include:
* A method to find if a given string array belongs to a mutant (located in utils/dna_checker.go)
* An API Handler to respond via POST request on /mutant (with DNA as JSON) if it corresponds to a mutant, storing the value in a database to avoid rechecking same dna
* An API Handler to respond via GET request on /stats the current ammount of different humans and mutants checked by the /mutant requests

### Built With

* [Go](https://golang.org)
* [MySQL](https://mysql.com)
* [Docker](https://docker.com)



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites for Local Development
* Go Programming Language - Install Go from [https://golang.org/doc/install](https://golang.org/doc/install)
* MySQL Database Server - Install MySQL from [https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/](https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/)

### Installation 

1. Clone the repo
   ```sh
   git clone git@github.com:gusgins/meli-backend.git
   ```
2. Change directory to meli-backend
   ```sh
   cd meli-backend
   ```
3. Install go module dependencies
   ```sh
   go mod download
   ```
4. Edit `config.yml` with the configuration you want to use (the database DATABASE_NAME has to be created and DATABASE_USER has to have access to it)
   ```yml
   database:
       host: DATABASE_HOST
       port: DATABASE_PORT
       name: DATABASE_NAME
       user: DATABASE_USER
       password: DATABASE_PASSWORD
   api:
       port: API_PORT
   ```
5. Run api server
   ```sh
   go run main.go
   ```

### Using Docker Compose
1. Clone the repo
   ```sh
   git clone git@github.com:gusgins/meli-backend.git
   ```
2. Change directory to meli-backend
   ```sh
   cd meli-backend
   ```
3. Edit `docker/app/.env` to match your environment
4. Build and start containers with docker-compose
   ```sh
   docker-compose up
   ```

<!-- USAGE EXAMPLES -->
## Usage

* Open http://localhost:8080/stats (Changing 8080 to API_PORT if modified) to check stats handler.
* Using [Postman](https://www.postman.com/downloads/) you can import the collection in `postman/the [postman folder](https://github.com/gusgins/meli-backend/tree/master/postman) to test the API from there.
* You can also use this script in `generator/main.go` to populate the database and test the API with random dna strings where:
  - n: number of requests to execute
  - maxsize: max size of dna array
  - concurrent: use goroutines for faster requests
```sh
   cd generator
   go run main.go -concurrent=1 -n=100 -maxsize=10
   ```


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

Gustavo Gingins - gusgins@gmail.com

Project Link: [https://github.com/gusgins/meli_backend](https://github.com/gusgins/meli_backend)

