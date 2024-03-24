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
The actual api server and it's handlers.
### internal/infrastructure/mediator
The mediator implementation.
### internal/infrastructure/persistence
All repositories.