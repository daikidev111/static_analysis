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
  <a href="https://github.com/daikidev111/static_analysis">
    <img src="image/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Static analysis tool</h3>

  <p align="center">
    The better error handling suggestion for a concurrent execution using go routine.
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
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project ##
Static analysis is a great tool to find problems often related to performance, coding style, and some logic errors without running the application.
I have implemented a tool that suggests to use sync.ErrGroup instead of sync.WaitGroup as sync.ErrGroup is generally more useful when you wish to handle error for groups of goroutines working on subtasks of a common task in parallel.


## How I did it ##
 
When you run the test from error_handling_test.go, it prints out the following AST.
Then, analyze the following AST tree that is comprised of AST nodes.
```sh
    88  .  .  .  .  .  1: *ast.GoStmt { <- AST node for Go Routine 
    89  .  .  .  .  .  .  Go: sample.go:11:2
    90  .  .  .  .  .  .  Call: *ast.CallExpr {
    91  .  .  .  .  .  .  .  Fun: *ast.FuncLit { <- AST node for function literal
    92  .  .  .  .  .  .  .  .  Type: *ast.FuncType {
    93  .  .  .  .  .  .  .  .  .  Func: sample.go:11:5
    94  .  .  .  .  .  .  .  .  .  Params: *ast.FieldList { <- AST node for field list
    95  .  .  .  .  .  .  .  .  .  .  Opening: sample.go:11:9
    96  .  .  .  .  .  .  .  .  .  .  List: []*ast.Field (len = 1) { <- contains one param
    97  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Field {
    98  .  .  .  .  .  .  .  .  .  .  .  .  Names: []*ast.Ident (len = 1) {
    99  .  .  .  .  .  .  .  .  .  .  .  .  .  0: *ast.Ident {
   100  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample.go:11:10
   101  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "ch"
   102  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Obj: *ast.Object {
   103  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Kind: var
   104  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "ch"
   105  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Decl: *(obj @ 97)
   106  .  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   107  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   108  .  .  .  .  .  .  .  .  .  .  .  .  }
   109  .  .  .  .  .  .  .  .  .  .  .  .  Type: *ast.ChanType {
   110  .  .  .  .  .  .  .  .  .  .  .  .  .  Begin: sample.go:11:13
   111  .  .  .  .  .  .  .  .  .  .  .  .  .  Arrow: - <- We care about only this arrow type for now.
   112  .  .  .  .  .  .  .  .  .  .  .  .  .  Dir: 3
   113  .  .  .  .  .  .  .  .  .  .  .  .  .  Value: *ast.Ident {
   114  .  .  .  .  .  .  .  .  .  .  .  .  .  .  NamePos: sample.go:11:18
   115  .  .  .  .  .  .  .  .  .  .  .  .  .  .  Name: "error" <- if it uses error then we can confirm that this goroutine func is used for error handling
   116  .  .  .  .  .  .  .  .  .  .  .  .  .  }
   117  .  .  .  .  .  .  .  .  .  .  .  .  }
   118  .  .  .  .  .  .  .  .  .  .  .  }
   119  .  .  .  .  .  .  .  .  .  .  }
   120  .  .  .  .  .  .  .  .  .  .  Closing: sample.go:11:23
   121  .  .  .  .  .  .  .  .  .  }
   122  .  .  .  .  .  .  .  .  }
```

<p align="right">(<a href="#top">back to top</a>)</p>


### Built With
* [Go](https://go.dev/)
* [Skeleton](https://github.com/gostaticanalysis/skeleton)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Daiki Kubo - [LinkedIn](https://www.linkedin.com/in/daiki-kubo/)

Project Link: [https://github.com/daikidev111/static_analysis](https://github.com/daikidev111/static_analysis)

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
