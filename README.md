# GraphQL Exercise

The main goal of this project is to be a place to learn about the implementation of a GraphQL server in Go.

A makefile is available to help you to run / build your code.

To run the project you'll need a mongodb. You can either run one with
```
make run-mongo
```
... or else provide it by yourself and replace the `MONGODB_URL` from `.env`

Then you'll start the project with
```
make run
```