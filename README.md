# reCharge inspired from [micro-template](https://p-source.780.ir/emad.gh/micro-template)

Based on Data Driven Design, we have segregated our charge in a separate, isolated domain. 

main parts of the project:

- Application (app)
- Controllers 
- Domain
- services


## Setup:
Before starting the application, one can migrate the database and check the connections with third party clients.

#### config file
for generate configs:
```
    make config
```
now we generate 2 configs for you:
- config.yaml (application reads data from this file for start)
- config.test.yaml (for test cases application read data from this file)

for migrate database and seed data into database
```
    go run main.go migrate init
    go run main.go migrate up
    go run main.go seed
```
    

## Submit a new service

We are *http1.1* agnostic. Everything should be came in **grpc** and well-defined in **protobufs**. To add a new service do as follows:

- explicitly define the service ib ```api/proto```
- run the following command in terminal
``` make proto```
- one should manipulate the ```app/router``` (clarifications needed)

#### services:
 - service tracer: jaeger
 - config server: etcd
 - cmd: cobra
 - configs: from file(dev), from config server(prod)
 - logger: zap
 - database: postgres(go-pg)
 - database ui: pgAdmin
