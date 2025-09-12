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

- **REST APIs**
  - REST APIs built with **Gin**    
  - Auto-generated Swagger documentation  

---

## 🛠️ Tech Stack

- **Language**: Go 1.22+  
- **Framework**: Gin  
- **Database**: PostgreSQL + SQLC    
- **Auth**: JWT & PASETO   
- **CI/CD**: GitHub Actions  

---

## ⚙️ Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/FilledEther20/Reg_Bank.git
cd Reg_Bank
