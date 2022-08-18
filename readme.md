# Welcome to GoBooking (UNDER CONSTRUCTION!!)
## A Modular Monlolith REST API made with Go!
This is my first Go personal project and I hope you enjoy it. This project is a REST API to manage a fictional hotel. I tried to follow the best practices of modern software development, such as: Test-Driven-Development, Domain-Driven-Design, Clean Architecture, Dependency Injection, REST and more.

Please notice that I'm still learning Go, DDD, Clean Architecture, etc so I might have some mistakes. Feel free to give me feedback :)

### Some features you will find in this project:
- Domain-Driven-Design (I tried at least)
- Clean Architecture
- Dependency Injection
- Repository Pattern
- REST
- Docker
- PostgreSQL
- Test-Driven-Development
- Git hooks
- Commit message conventions
- Small commits
- CI/CD with GitHub Actions
- Principles such as DRY, KISS, YAGNI
- And much more...

----------

### INSTALLING
### go
This project uses the Go language. Make sure you install it before you start. I'm using the version 1.18.

https://go.dev/doc/install

#### docker
To run tests, development mode, migrations, etc. with ease you must use Docker, otherwise you can get
unexpected behaviors. You wont worry about setting database, docker will do it for you.

https://docs.docker.com/engine/install/

#### node.js
This project uses git hooks and I'm using node.js to run them. Make sure you have installed it, otherwhise hooks wont work. Version 14 or higher is recommended, but other versions should work.

https://nodejs.org/en/

----------

### SETTING UP
#### Environment
You must create three files called: .env.test, .env.dev and .env.prod. Each file represents a different environment setup.
It will be used by Docker and the application. Make sure you follow the instructions in .example.env.

#### Makefile
The Makefile contains several scripts to run the application, test or migration. However, some of the scripts should run only with docker, so make sure you use it correctly.

#### Git hooks
Make sure you run `make init` to initialize the git hooks. It will automatically add the hooks to your repository in order to run lint, test, etc. before commits, etc.

----------

### RUNNING
After you have setup the environment, you can run the application. To do so, run `make docker_dev` or `make docker_prod`.

----------

### TESTING
To run the tests, simply run `make test`. Make sure you've setup the environment files before.
