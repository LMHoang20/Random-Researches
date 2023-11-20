# Amazon Aurora

Amazon Aurora is a MySQL and PostgreSQL-compatible relational database built for the cloud, that combines the performance and availability of traditional enterprise databases with the simplicity and cost-effectiveness of open source databases.

## Aurora Postgres

### Aurora Postgres Features

- 3x performance of PostgreSQL

- 1/10th the cost of commercial database

- 99.99% availability

- 6 copies of your data across 3 AZs

- Backups are continuous and incremental

- Storage is self-healing

- Aurora is designed to transparently handle the loss of up to two copies of data without affecting database write availability and up to three copies without affecting read availability

## Fast Database Cloning

Aurora Fast Database Cloning allows you to quickly and inexpensively create multiple copies of your Aurora PostgreSQL databases for development, test, and QA. You can create a clone of an Aurora PostgreSQL database in minutes, and it will be available for use immediately. Cloning is easy to use and incurs no storage costs, making it a cost-effective way to create multiple copies of your database for development, test, and QA.

### Copy on Write

A clone is a full replica of a database is unreasonable, especially when the database has terabytes of data on it. For this reason, Aurora clones by creating a copy-on-write clone. This means that each instance of the clone is a virtual pointer back to the original database. When you make a change to the clone, the change is written to the clone, not the original database. Making changes to the original database also creates a new copy of the data, which is then written to the clone.

## Aurora Serverless

Aurora Serverless is an on-demand, auto-scaling configuration for Amazon Aurora (MySQL-compatible and PostgreSQL-compatible editions), where the database will automatically start up, shut down, and scale capacity up or down based on your application's needs. It enables you to run your database in the cloud without managing any database instances. It's a simple, cost-effective option for infrequent, intermittent, or unpredictable workloads.

## Aurora Global Database

Aurora Global Database is a feature that allows you to create up to five secondary read-only replicas of your Aurora database in the same AWS Region, with latency of less than one second. This enables applications to read a single up-to-date copy of data from multiple regions. Aurora Global Database spans multiple AWS Regions, enabling low latency global reads and disaster recovery from region-wide outages.

### Region Cluster

Aurora Global Database uses a single primary cluster and up to five secondary clusters, each in a separate AWS Region. The primary cluster is the cluster that you create first. The primary cluster is the only cluster that you can write to. The secondary clusters are read-only clusters. You can create a secondary cluster in any AWS Region that Aurora supports. You can create up to five secondary clusters.

###  Reader/Writer Instances

You can scale your read/write capacity by creating multiple DB instances and distributing the read/write workload across the instances.

An Aurora cluster instance might be either a writer or a reader. Aurora clusters allow one writer and up to 15 readers. The instance role might change failover happens.

## Auto Scaling / Dynamic Resizing

When you use Aurora Serverless, you specify the minimum and maximum capacity of the DB cluster. Aurora Serverless automatically scales capacity up to the maximum capacity when demand increases, and scales capacity down to the minimum capacity when demand decreases. Scaling is rapid, with Aurora Serverless adjusting capacity within seconds, and there is no need to pre-provision capacity for database instances.

### Buffer Pool Resizing

Aurora Serverless automatically scales the buffer pool size as the capacity of the DB cluster changes. The buffer pool is a cache of data pages that have been read from the database storage into memory. The buffer pool is used to speed up access to data by reducing the amount of disk I/O that is needed to access data pages. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance. The buffer pool is sized based on the amount of memory that is available on the DB instance.

## Trusted Language Extension

Aurora Postgres is implemented in C, which is scary for most developers to work with. Aurora Postgres supports Trusted Language Extensions, which allows you to write new extensions in languages like JavaScript, PL/pgSQL, (Rust), etc.

## Amazon GuardDuty RDS Protection

Aurora GuardDuty will detect anomalies in memory usages, CPU usage, and network traffic. It will also detect anomalies in the number of connections, the number of queries, and the number of transactions. 

It will also detect against malicious activity, such as someone bruteforcing, guessing, password spraying. It will also have mechanism to protect against stolen credentials.

## Blue-Green Deployments

Aurora supports blue-green deployments. This means that you can deploy a new version of your application to a new Aurora cluster and then switch over to the new cluster when you are ready. This allows you to deploy new versions of your application without any downtime.

# References:

- [AWS Aurora](https://aws.amazon.com/rds/aurora/)
- [AWS re:Invent 2022 - Deep dive into Amazon Aurora and its innovations (DAT326)](https://www.youtube.com/watch?v=pzZydB78Eyc)
