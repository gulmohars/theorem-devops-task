# Theorem-SRE-Task
Theorem LLC - DevOps Prototyping Exercise

### Go Application.

I used golang to create a simple application exposing an HTTP API with 2 endpoints:
  -  (/) Endpoint 1: Json payload response
  -  (/health) Endpoint 2: used for healthcheck

The application will run on port: 8080

### Dockerfile.

A multi stage Dockerfile, divided in two, the first part include the building steps and the second part the final stage with the minimal size.

### Infrastructure. (AWS)

For simplicity I decided to use cloudformation for the IaC. It can also be done with other tool like Ansible or Terraform, but this implies more setup time, for the terraform and Ansible users policies and I prefer to dedicate more time to deeply explain the 4 points needed for the task.

####   - VPC.

I started with the VPC/Network setup, it includes 2 Subnets in different Availability zones for High Availability, can be extended to more Subnets if needed.

There is also a setup for the Internet gateway and the needed routing table to get access from and to the Internet. 

As outputs I setup, the Azs, the Subnets, the VPC Id and the VPC DefaultSG, all this with the idea to create a Cloudformation Nested Stack.

####   - Loadbalancer.

