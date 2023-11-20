# Relational Database Service (RDS)

AWS RDS is a managed relational database service that provides an easy way to set up, operate, and scale a relational database in the cloud. It provides cost-efficient and resizable capacity while automating time-consuming administration tasks such as hardware provisioning, database setup, patching and backups.

Very similar to just running a database on an EC2 instance, but with a lot of the management taken care of for you.

## RDS Features

- Supports Amazon Aurora, PostgreSQL, MySQL, MariaDB, Oracle Database, and SQL Server database engines

- Automated backups

- Point-in-time recovery for your DB Instance

## RDS Pricing

You are charged for each hour your DB Instance is running, storage you provision for your instance, IOs you consume, backup storage you consume, and data transfer you consume

### RDS Pricing Models

- On-Demand: Pay by the hour with no long-term commitments

- Reserved Instances: Pay a one-time fee to reserve an instance for a one or three year term and then pay a reduced hourly rate for the duration of the term

## RDS Security

- Encryption at rest

- Encryption in transit

- Network isolation using Amazon VPC

- IAM integration

- Amazon RDS supports SSL to encrypt connections

## RDS Backups

- Automated backups

- Database snapshots

- Transaction logs

## RDS Read Replicas

To optimize performance, you can create one or more read replicas of your source DB Instance and serve read traffic from multiple copies of your data.

Automatically managed by RDS, so you don't have to worry about it.

## RDS Multi-AZ

Multi-Availability Zone (Multi-AZ) deployments for enhanced availability and data durability. In addition to automatic failover, Multi-AZ deployments automatically replicate data synchronously across Availability Zones to protect against data loss due to a storage failure in a single Availability Zone.
