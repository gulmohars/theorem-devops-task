# Theorem-SRE-Task
Theorem LLC - DevOps Prototyping Exercise. This Document specify in 4 parts, all the resources needed to complete the Theorem-SRE-Task assigned:

### 1. Application.

I used golang to create a simple application exposing an HTTP API with 2 endpoints:
  -  (/) Endpoint 1: Json payload response
  -  (/health) Endpoint 2: used for healthcheck

The application will run/listen on port: 8080

### 2. Containerization.

A multi stage Dockerfile, divided in two, the first part include the building steps and the second part the final stage with the minimal size.

### 3. Infrastructure. (AWS)

For simplicity I decided to use cloudformation for the IaC. It can also be done with other tool like Ansible or Terraform, but this implies more setup time, for the terraform and Ansible users policies and I prefer to dedicate more time to deeply explain the 4 points needed for the task.

####   - VPC.

I started with the VPC/Network setup, it includes 2 Subnets in different Availability zones for High Availability, can be extended to more Subnets if needed.

I created a VPC, and specified a range of IPv4 addresses for the VPC in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.20.0.0/16. This is the primary CIDR block for your VPC.

A VPC spans all the Availability Zones in the region. After creating a VPC, you can add one or more subnets in each Availability Zone. In our case I created 2. When I created a subnet, I specified the CIDR block for the subnet, which is a subset of the VPC CIDR block. Each subnet must reside entirely within one Availability Zone and cannot span zones. Availability Zones are distinct locations that are engineered to be isolated from failures in other Availability Zones. By launching instances in separate Availability Zones, you can protect your applications from the failure of a single location.

There is also a setup for the Internet gateway and the needed routing table to get access from and to the Internet. An internet gateway enables communication over the internet.

Each subnet must be associated with a route table, which specifies the allowed routes for outbound traffic leaving the subnet. Every subnet that you create is automatically associated with the main route table for the VPC.

AWS provides two features that I can use to increase security in the VPC created: security groups and network ACLs. Security groups control inbound and outbound traffic for the instances, and network ACLs control inbound and outbound traffic for our subnets. Security groups are enough for my actual needs in this exercise.

As outputs I setup, the Azs, the Subnets, the VPC Id and the VPC DefaultSG, all this with the idea to create a Cloudformation Nested Stack.

####   - Loadbalancer.

In the loadbalancer yaml file I added the Loadbalancer definition, the healthcheck setup (path and httpcode expected) used during the Blue-green deployment, the Security Group and the cookie stickiness in case a cookies auth is needed.

Each target group is used to route requests to one or more registered targets. When I created the listener rule, I specify a target group and condition (/). When a rule condition is met, traffic is forwarded to the corresponding target group. I defined the health check settings of the load balancer on a per target group basis.

After I specify a target group in a rule for a listener, the load balancer continually monitors the health of all targets registered with the target group that are in an Availability Zones enabled for the load balancer. The load balancer routes requests to the registered targets that are healthy.

As output I added the Loadbalancer Targetgroup, the Security Group created and the Url where the service is going to be ready.

####   - ECS Cluster.

