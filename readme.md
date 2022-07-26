# Welcome to GoBooking
## A REST API made with Go!

----------

### Install
#### go
This project uses the Go language. Make sure you install it before you start. I'm using the version 1.18.

https://go.dev/doc/install

##### docker
To run tests, development mode, migrations, etc. with ease you must use Docker, otherwise you can get
unexpected behaviors. You wont worry about setting database, docker will do it for you.

https://docs.docker.com/engine/install/

##### golangci-lint
This project follow strict lint rules. To achieve this, I am using golangci-lint. Make sure you have installed it.

https://golangci-lint.run/

##### node.js
This project uses git hooks and I'm using node.js to run them. Make sure you have installed it, otherwhise hooks wont work. Version 14 or higher is recommended, but other versions should work.

https://nodejs.org/en/

----------

### Setup
##### Environment
You must create three files called: .env.test, .env.dev and .env.prod. Each file represents a different environment setup.
It will be used by Docker and the application. Make sure you follow the instructions in .example.env.

##### Makefile
The Makefile contains several scripts to run the application, test or migration. However, some of the scripts should run only with docker, so make sure you use it correctly.

##### Git hooks
Make sure you run `make init` to initialize the git hooks. It will automatically add the hooks to your repository in order to run lint, test, etc. before commits, etc.

----------

### Run
After you have setup the environment, you can run the application. To do so, run `make docker_dev` or `make docker_prod`.

----------

### Test
To run the tests, simply run `make test`. Make sure you've setup the environment files before.
