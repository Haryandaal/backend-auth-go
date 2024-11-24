# Backend API for Merchant & Bank Transactions

This project is a backend API designed to simulate interactions between merchants, customers, and banks using JSON files. It features token-based authentication (JWT) and activity logging.

## Features
- **Login**: Separate login for customers and merchants.
- **Payments**:
    - Customer-to-Customer.
    - Merchant-to-Bank.
- **Logout**: Ends user session.
- **Activity Logging**: All actions are logged in a history file.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Haryandaal/backend-auth-go
   cd backend-auth-go
   ```
2. Install dependencies:
    ```bash
   go mod init
   go get github.com/gorilla/mux github.com/golang-jwt/jwt/v5 github.com/google/uuid
   ```
3. Run the application:
    ```bash
   go run main.go
   ```

## API Endpoints
1. **Customer Login**
    
    POST `/login` 

      Request
   ```json
   {
      "email" : "haryanda@example.com",
      "password" : "password"
   }
   ```
    Response
    ```json
   {
      "token" : "<JWT_TOKEN>"
   }
   ```
2. **Merchant Login**

   POST `/login`  -> login for customer
   
   POST `/merchant/login` -> login for merchant

   Request
   ```json
   {
      "email" : "merchant@example.com",
      "password" : "password"
   }
   ```
4. **Payments**

   POST `/payment`

   POST `/merchant/payment`

   Request
   ```json
   {
      "to_email": "recipient@example.com",
      "amount": 100
   }
   ```
   Response
    ```json
   {
      "message" : "message successful"
   }
   ```
5. **Payments**

   POST `/logout`

## JSON Data Files
- `customers.json` = Customer data
- `merchants.json` = Merchant data
- `banks.json` = Bank data
- `history.json` = Logs all action

## Example Logs
```json
  [
    {
      "id": "66e167f6-fcf3-4a22-9e28-22f306bd3869",
      "action": "Login",
      "detail": "haryanda@example.com logged in",
      "date": "2024-11-24 12:29:38"
    }
  ]
```
