<div id="top"></div>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/daikidev111/clean_arch_go/">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Static analysis tool</h3>

  <p align="center">
    Static analysis is a great tool to find problems often related to performance, coding style, and some logic errors without running the application.
    I have implemented a tool that suggests to use sync.ErrGroup instead of sync.WaitGroup as sync.ErrGroup.
    The reason for building this tool is that errGroup is generally more useful when you wish to execute groups of goroutines working on subtasks of a common task in parallel.
    <br />
    <a href="https://github.com/daikidev111/static_analysis/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/daikidev111/static_analysis/">View Demo</a>
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
        <li><a href="#Running MySQL, Redis, PHPMyAdmin, Swagger, Proxy containers with docker-compose">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


**AST fields**



<p align="right">(<a href="#top">back to top</a>)</p>


### Built With
* [Go](https://go.dev/)
* [Skeleton](https://github.com/gostaticanalysis/skeleton)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

This is an instruction of how to set up and start the project locally.
To get a local copy up and running follow these simple steps.

### Prerequisites

* Install Go

  Follow the steps shown in this [website](https://golang.org/dl/)

### Installation

1. Clone this repo
   ```sh
   git clone https://github.com/daikidev111/clean_arch_go/
   ```
2. Install goimports and golangci-lint
   ```sh
   make local-install
   ```
 
### Running MySQL, Redis, PHPMyAdmin, Swagger, Proxy containers with docker-compose

```sh
  docker-compose up -d
  
  The list of containers should be like this.
  
             Name                         Command               State                 Ports              
-------------------------------------------------------------------------------------------------------
clean_arch_go_mysql_1        docker-entrypoint.sh mysql ...   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
clean_arch_go_phpmyadmin_1   /docker-entrypoint.sh apac ...   Up      0.0.0.0:4000->80/tcp             
clean_arch_go_proxy_1        /docker-entrypoint.sh ngin ...   Up      0.0.0.0:3010->3010/tcp, 80/tcp   
clean_arch_go_redis_1        docker-entrypoint.sh redis ...   Up      0.0.0.0:6379->6379/tcp           
clean_arch_go_swagger-ui_1   /docker-entrypoint.sh sh / ...   Up      80/tcp, 127.0.0.1:3000->8080/tcp
```

### Check if golangci-lint and gofmt work

gofmt is a tool that automatically formats the Go source code.
   ```sh
   make fmt
   ```
golangci-lint is a tool that checks if the source code follows the coding standard.
   ```sh
   make lint
   ```

### Initiate the API
   ```sh
   go run ./cmd/main.go
   ```

### Execute the /user/create API, using the Curl command shown below;
   ```sh
curl -X 'POST' \
  'http://localhost:8080/user/create' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Joji"
}'
  ```

### 

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Daiki Kubo - [LinkedIn](https://www.linkedin.com/in/daiki-kubo/)

Project Link: [https://github.com/daikidev111/clean_arch_go/](https://github.com/daikidev111/clean_arch_go/)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
