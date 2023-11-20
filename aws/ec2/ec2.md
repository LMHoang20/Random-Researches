# Amazon Elastic Compute Cloud (EC2)

One of the first services offered by **AWS**, **EC2** is a *virtual machine service*. It is a web service that allows users to rent virtual computers to run their applications. 

Back in the day, if you wanted to run a server, you had to buy a physical server. This was expensive and time consuming. You had to buy the server, set it up, and maintain it. If you wanted to run multiple servers, you had to buy multiple servers. If you wanted to run a server for a short period of time, you had to buy a server and then get rid of it when you were done. Sometimes, you had to buy a server that was more powerful than you needed just to make sure that you had enough power to run your application. EC2 solves all of these problems.

Basically, **Amazon** built a massive warehouse full of servers. They then built a web service that allows users to rent virtual servers by the hour. Users can rent as many servers as they need, as powerful as they need, for as long as they need. They can return the servers when they are done. They can downgrade, upgrade or change the number of servers at any time. That's why it's called ***Elastic** Compute Cloud*.

Every EC2 instance is *managed* by a **hypervisor**. The hypervisor is a piece of software that allows multiple operating systems to run on a single physical server. It is responsible for managing the resources of the physical server and allocating them to the virtual servers. It is also responsible for managing the virtual servers and allocating resources to the applications running on them.


# EC2 Instance Types

EC2 instances are grouped into families based on their characteristics. Each family is designed for a specific use case.

## General Purpose

General purpose instances are designed for a balance of compute, memory, and networking resources. They are a good choice for a variety of workloads.

## Compute Optimized

Compute optimized instances are designed for compute-intensive workloads that require high performance processors. They are a good choice for batch processing, media transcoding, high performance web servers, and high performance computing (HPC).

## Memory Optimized

Memory optimized instances are designed for memory-intensive workloads that require high performance processors. They are a good choice for high performance databases, distributed web scale cache stores, and in-memory analytics.

## Accelerated Computing

Accelerated computing instances are designed for compute-intensive workloads that require high performance processors and hardware accelerators. They are a good choice for machine learning, high performance computing (HPC), video encoding, and graphics intensive applications.

## Storage Optimized

Storage optimized instances are designed for storage-intensive workloads that require high performance processors and high performance storage. They are a good choice for data warehousing, Hadoop, and distributed file systems.

# Payment Options

## On-Demand

Users are charged for the instances that they use by the hour.

On-demand instances are the most expensive option. They are a good choice for short-term, spiky, or unpredictable workloads.

## Reserved

Users can reserve instances for a one or three year term. Pay a one-time fee to reserve the instance and then pay a reduced hourly rate for the duration of the term.

Reserved instances are a good choice for steady-state workloads. They provide a significant discount compared to on-demand instances. They are available in three payment options: all upfront, partial upfront, and no upfront.

## Spot

Users can bid on unused EC2 capacity. Pay the current spot price for the instance. If the spot price exceeds the bid price, the instance is terminated.

Spot instances are the cheapest option. They are a good choice for workloads that are flexible about when they run. They are available in three payment options: all upfront, partial upfront, and no upfront.

# Hypervior

When EC2 was first launched, it used a hypervisor called **Xen**. In 2017, **AWS** launched a new hypervisor called **Nitro**. Nitro is a lightweight hypervisor that is designed to provide better performance and security than **Xen**.


# Amazon Machine Image (AMI)

An AMI is a template for an EC2 instance. It contains the operating system and any additional software that is required to run the application. It also contains the configuration settings for the instance. 

# Persistent Storage

EC2 instances are *ephemeral*. This means that they are temporary. When an instance is terminated, all of the data on the instance is lost. If you want to store data permanently, you need to use persistent storage.

## Storages

- Elastic Block Store (EBS)
- Elastic File System (EFS)
- Simple Storage Service (S3)
- Glacier
- ...

# Elastic Ecosystem

## Elastic Load Balancer (ELB)

Elastic Load Balancing automatically distributes incoming application traffic across multiple targets, such as Amazon EC2 instances, containers, IP addresses, and Lambda functions. It can handle the varying load of your application traffic in a single Availability Zone or across multiple Availability Zones. Elastic Load Balancing offers three types of load balancers that all feature the high availability, automatic scaling, and robust security necessary to make your applications fault tolerant.

## Elastic Container Service (ECS)

Elastic Container Service (ECS) is a highly scalable, high performance container management service that supports Docker containers and allows you to easily run applications on a managed cluster of Amazon EC2 instances.

