# **Microservices Architecture with gRPC, GraphQL & Distributed Databases**  

This project is a **scalable microservices system** leveraging **gRPC for inter-service communication**, **GraphQL as the API gateway**, and **distributed data storage** for optimized performance.  

## **Architecture Overview**  

- **Account Service** – Manages authentication & profiles (PostgreSQL for relational integrity).  
- **Catalog Service** – Handles product data & search (Elasticsearch for fast indexing).  
- **Order Service** – Processes transactions & order tracking (PostgreSQL for consistency).  
- **GraphQL API Gateway** – Aggregates responses, enabling flexible queries across services.  

## **Infrastructure**  

- **Dockerized microservices** with **health checks & service discovery**.  
- **PostgreSQL & Elasticsearch** for structured and search-optimized data storage.  
- **gRPC** for **efficient, low-latency** inter-service calls.  
- **GraphQL** for **frontend-optimized** querying.  
- Uses **stable Go versions** to maintain reliability in production.  

The decision to use a hybrid of relational and NoSQL databases, coupled with gRPC and GraphQL, was driven by real-world performance, scalability, and maintainability needs. This architecture is developer-friendly, production-ready, and designed for future growth, making it an excellent model for enterprise-grade distributed systems.
