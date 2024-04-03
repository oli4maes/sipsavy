# SipSavy

SipSavy is a REST API to store cocktail recipes.

# Architecture

This project is an attempt to use vertical slice architecture with the mediator pattern.
All http requests will be handled in the http module. The http module will use the mediator implementation to send
an in-memory message `TRequest` which will be handled by the corresponding handler which in turn will return
a response `TResponse`. Following this pattern, the only dependencies should lie in the mediator handlers, this
will ensure that there are no circular dependencies.

## Project structure

### cmd/api

This is the entry point of the application.

### internal/features

All mediator handler are defined here.

### internal/infrastructure/http

The actual api server and its handlers.

### internal/infrastructure/mediator

The mediator implementation.

### internal/infrastructure/persistence

All repositories.

# Local development environment

The easiest way to get going with this codebase is to run a docker container for the sql database.
Since this project makes use of Microsoft SQL Server and this database does not support ARM64 architecture
we will run an instance of Azure SQL Edge which has all the features we require.  

`docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD={{PASSWORD_01}}" -p 1433:1433 --name sql --hostname sql 
-d mcr.microsoft.com/azure-sql-edge:latest`

You will then need to set an environment variable: `CONNECTION_STRING` with the value: `Server=localhost,1433\;Database=sipsavy\
;User=SA\;Password={{PASSWORD_01}}\;TrustServerCertificate=True`
