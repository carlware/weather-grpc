# This example expose a GRPC server with a weather service

## How to test it
```
# run grpc server
# this will expose a server on 127.0.0.1:9003
go run cli/main.go serve

# if we want to test it
# evans tool can be used
evans -p 9003 -r

# or if we want to test by command line
# the following command can be executed
# where -c is a City
# Temperature is shown in Kelvin but will be improved
go run cli/main.go  get-temp -c New\ York

```

## TODO
* [x] Create models
* [x] Configure and generate grpc code
* [x] Add CLI
* [x] Support display temp in kelvin, celsius or farenheit
* [ ] Add test cases
* [ ] Refactor
* [ ] Improve readme
* [ ] Add make files



