# RegBank Backend Service

A complete backend web service for a simple bank application, built with **Golang**.  
The service exposes secure APIs to:  
- Create and manage bank accounts  
- Record all balance changes (ledger entries)  
- Perform money transfers between accounts (transaction-safe)  

---

## ✨ Features

- **Accounts & Transactions**
  - Create accounts with owner, balance, and currency  
  - Record every balance change as an entry  
  - Transfer money safely between accounts using DB transactions  

- **Authentication & Authorization**
  - User registration & login  
  - Password hashing with **Bcrypt**  
  - Secure token-based authentication with **JWT & PASETO**  
  - Role-based access control (RBAC)  

- **REST & gRPC APIs**
  - REST APIs built with **Gin**  
  - gRPC APIs with **protobuf**  
  - gRPC-Gateway for serving both HTTP & gRPC  
  - Auto-generated Swagger documentation  

- **Infrastructure & Deployment**
  - PostgreSQL with SQLC for type-safe queries  
  - Dockerized service & Docker Compose for local dev  
  - Kubernetes deployment on AWS EKS  
  - GitHub Actions for CI/CD pipelines  

- **Advanced Backend Topics**
  - Database migrations & isolation levels  
  - Background workers with Redis + Asynq  
  - Email verification via Gmail SMTP  
  - Structured logging & middleware  
  - Graceful server shutdown  

---

## 🛠️ Tech Stack

- **Language**: Go 1.22+  
- **Framework**: Gin  
- **Database**: PostgreSQL + SQLC  
- **Queue**: Redis + Asynq  
- **Auth**: JWT & PASETO  
- **Deployment**: Docker, Kubernetes (AWS EKS)  
- **CI/CD**: GitHub Actions  

---

## ⚙️ Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/FilledEther20/Reg_Bank.git
cd Reg_Bank
