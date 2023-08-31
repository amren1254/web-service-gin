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

# This is a basic workflow to help you get started with Actions

name: CICD

# Controls when the action will run. Triggers the workflow on push or pull request
# events but only for the master branch
on:
  push:
    branches:
    - master
    - staging
    - develop
  pull_request:
    branches:
    - master
    - staging
    - develop

# A workflow run is made up of one or more jobs that can run sequentially or in parallel
jobs:
  # The "build" workflow
  build:
    # The type of runner that the job will run on
    runs-on: ubuntu-latest

    # Steps represent a sequence of tasks that will be executed as part of the job
    steps:
    # Checks-out your repository under $GITHUB_WORKSPACE, so your job can access it
    - uses: actions/checkout@v2
    
    # Setup Go
    - name: Setup Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.14.0' # The Go version to download (if necessary) and use.
    
    # Install all the dependencies
    - name: Install dependencies
      run: |
        go version
        go get -u golang.org/x/lint/golint
        
    # Run build of the application
    - name: Run build
      run: go build . 
      
    # Run vet & lint on the code
    - name: Run vet & lint
      run: |
        go vet .
        golint .
    
    # Run testing on the code
    - name: Run testing
      run: cd test && go test -v
    
    # Send slack notification
    - name: Send slack notification
      uses: 8398a7/action-slack@v3
      with:
        status: ${{ job.status }}
        fields: repo,message,commit,author,action,eventName,ref,workflow,job,took # selectable (default: repo,message)
      env:
        SLACK_WEBHOOK_URL: ${{ secrets.SLACK_WEBHOOK_URL }} # required
      if: always() # Pick up events even if the job fails or is canceled.

  # The "deploy" workflow
  deploy:
    # The type of runner that the job will run on
    runs-on: ubuntu-latest
    needs: [build] # Only run this workflow when "build" workflow succeeds
    if: ${{ github.ref == 'refs/heads/master' && github.event_name == 'push' }} # Only run this workflow if it is master branch on push event
    steps:
    - uses: actions/checkout@v2

    # Deploy to Docker registry
    - name: Deploy to Docker registry
      uses: docker/build-push-action@v1
      with:
        username: ${{ secrets.DOCKER_USERNAME }}
        password: ${{ secrets.DOCKER_PASSWORD }}
        repository: wilsontanwm/gosimple
        tag_with_ref: true