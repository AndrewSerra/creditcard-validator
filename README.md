# Credit Card Validator

This project is using the Luhn algorithm to verify
a given credit card number is valid. 

The project has two components. One is the validator
directory, and the other is the API. 

## Validator

The validator is where Luhn algorithm is used to verify
a card number. Independently this directory can be used
as a CLI tool as well. 

## API

The API has a single endpoint at the moment which is a
POST request `http://localhost:3000/verify`.

The postman json file can be used to send requests or 
simply using curl.

## Build the program

Navigate to the project directory and run:

```shell
go build
```

This will create a binary. If using Windows the file extension
will be `.exe`.

## Running the API

The default port is set to 3000. To change this set an
environment variable in the format `PORT=<new-port>`. 

Another approach is to prepend the command line run command
with the environment variable. An example is the following.

```shell
PORT=3001 go run main.go
```

If the program is built, the following command will run the 
binary file. Navigate to where the build output is and run the command:

```shell
./creditcard-validator
```

 