# Simple E-commerce API

🚀 A simple backend API for an e-commerce platform, built with **Golang**, **Gin**, **Gorm**, and **MariaDB**.  
Designed for learning, experimentation, and potential future development.

---

## 📦 Features
- User Registration & Login
- Store Registration
- Product Management:
  - Add Product
  - Get All Products
  - Search Product
- Cart Management:
  - Add to Cart
  - View Cart
  - Remove from Cart
- Address Management:
  - Add Address
  - Update Address
  - Delete Address
  - Get Address
- Order Management:
  - Checkout

---

## 🛠️ Tech Stack
- **Backend:** Go (Golang)
- **Framework:** Gin
- **ORM:** Gorm
- **Database:** MariaDB

---

## 🏗️ Database Structure
Tables:
- `user`
- `role`
- `store`
- `product`
- `order`
- `orderItem`
- `cart`
- `cartItem`
- `address`

---

## 🚀 Getting Started

### 1. Clone the repository
```bash
git clone https://github.com/azmiagr/simple-ecommerce-api.git
cd simple-ecommerce-api
```

### 2. Setup Database
- Create a new MariaDB database.
- Configure your database connection in your environment variables or config file.

**Example `.env` file:**
```env
DB_USER=
DB_PASS=
DB_HOST=
DB_PORT=
DB_NAME=
```

### 3. Install Dependencies
```powershell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
```

### 4. Run the Project
```powershell
cd cmd\app\
go run main.go
```

---

## 💡 Notes
- This project is a **learning project** and still under development.
- Some features (validation, security, optimization) might be improved over time.
- Feel free to fork and adapt to your needs!

---

## 📬 Contact
Made by [azmiagr](https://github.com/azmiagr)
