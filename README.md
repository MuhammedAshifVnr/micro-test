
# **Microservices with Go: User Management and Method Execution**

## **Overview**
This project implements two Go-based microservices with **gRPC** communication, focusing on user management and task processing:  
- **Microservice 1**: Manages user data with CRUD operations, backed by PostgreSQL.  
- **Microservice 2**: Handles sequential and parallel task execution using a task queue mechanism.

---

## **Features**
### **Microservice 1**:  
- Create, retrieve, update, and delete user records.  
- Offers gRPC-based APIs for inter-service communication.  
- Integrates with PostgreSQL for persistent data storage.  

### **Microservice 2**:  
- Executes two distinct methods: **Method 1** (sequential) and **Method 2** (parallel).  
- Fetches user data from **Microservice 1** through gRPC.  
- Simulates processing tasks using a custom task queue system.

---

## **Prerequisites**
- **Go** (1.19 or higher)  
- **Docker** and **Docker Compose**  
- **PostgreSQL**  
- Optional: **Redis** (for future caching enhancements).  
- Optional: **Kubernetes** and `kubectl` for deployment in a cluster.  

---

## **Setup**

### **Step 1: Clone the Repository**
```bash
git clone https://github.com/MuhammedAshifVnr/micro-test.git
cd micro2
```

### **Step 2: Run Services with Docker Compose**
Use the provided `docker-compose.yml` file to start all services:  
```bash
docker-compose up --build
```

This command will start:  
- **Microservice 1**: REST APIs available on `http://localhost:8081`  
- **Microservice 2**: gRPC server on `http://localhost:5001`  
- **PostgreSQL**: Running on `localhost:5432`

---

## **Microservice 1: User Management**

### **API Endpoints**

| Method | Endpoint         | Description                   |
|--------|------------------|-------------------------------|
| POST   | `/user`          | Create a new user.            |
| GET    | `/user/:id`      | Retrieve user by ID.          |
| GET    | `/methods`       | List all users.               |
| PUT    | `/user/:id`      | Update user details.          |
| DELETE | `/user/:id`      | Delete user by ID.            |

#### **Example Usage**  
1. **Create a User**  
   ```json
   POST /user
   {
       "name": "Alice",
       "email": "alice@example.com",
       "phone": 1234567890
   }
   ```

2. **Retrieve User by ID**  
   ```http
   GET /user/1
   ```

3. **Update User**  
   ```json
   PUT /user/1
   {
       "name": "Alice Updated",
       "email": "alice.new@example.com"
   }
   ```

4. **Delete User**  
   ```http
   DELETE /user/1
   ```

5. **List All Users**  
   ```http
   GET /user/list
   ```

---

## **Microservice 2: Task Execution**

### **Methods**

| Method  | Description                                                |
|---------|------------------------------------------------------------|
| Method1 | Executes tasks **sequentially** using a task queue.        |
| Method2 | Executes tasks **in parallel** for concurrent processing.  |

### **Behavior**  

- **Method1**:  
  - Adds tasks to a queue for sequential execution.  
  - Simulates a delay using a configurable wait time.  

- **Method2**:  
  - Processes all tasks in parallel without dependencies.  

#### **Example Usage**  
1. **Execute Method 1 (Sequential)**  
   ```json
   POST /method
   {
       "method": 1,
       "waitTime": 10
   }
   ```

2. **Execute Method 2 (Parallel)**  
   ```json
   POST /method
   {
       "method": 2,
       "waitTime": 5
   }
   ```

---


## **How It Works**

1. **Microservice 1** handles user data management and serves as the data source for Microservice 2.  
2. **Microservice 2** utilizes gRPC to communicate with Microservice 1 for fetching user data.  
3. The custom task queue in Microservice 2 executes tasks sequentially (Method1) or concurrently (Method2).  

---

## **Future Enhancements**
- **Caching**: Add Redis for caching frequently accessed user data.  
- **Kubernetes Deployment**: Deploy services in a Kubernetes cluster.  
- **Monitoring**: Implement Prometheus and Grafana for performance monitoring.  
