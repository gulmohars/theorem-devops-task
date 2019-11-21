# Theorem-SRE-Task
Theorem LLC - DevOps Prototyping Exercise

### Go Application.

I used golang to create a simple application exposing an HTTP API with 2 endpoints:
  -  (/) Endpoint 1: Json payload response
  -  (/health) Endpoint 2: used for healthcheck

The application will run on port: 8080

### Dockerfile.

A multi stage Dockerfile, divided in two, the first part include the building steps and the second part the final stage with the minimal size.

