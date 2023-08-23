# Docker

docker ps
docker ps -a
docker run <container-id>
docker exec -it <container-id> /bin/bash  -> interactive terminal
docker image ls
docker container ls
docker build -tag <tag-name> <path>

docker rmi <container-id> -> if container id is known
docker rm
docker system prune -> remove all dangling images

# Postgres

# Golang

Method :- A method is a special function that is scoped to a specific type in Go. Unlike function, a method can only be called from the instance of the type it was defined on.

```GO
package main

import "fmt"

type Article struct {
 Title  string
 Author string
}

func (a Article) String() string {
 return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
 Title  string
 Author string
 Pages  int
}

func (b Book) String() string {
 return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
 String() string
}

func main() {
 a := Article{
  Title:  "Understanding Interfaces in Go",
  Author: "Sammy Shark",
 }
 Print(a)

 b := Book{
  Title:  "All About Go",
  Author: "Jenny Dolphin",
  Pages:  25,
 }
 Print(b)
}

func Print(s Stringer) {
 fmt.Println(s.String())
}
```
