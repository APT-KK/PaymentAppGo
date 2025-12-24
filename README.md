# PaymentAppGo

This is a beginner learning project for Go (Golang) and React. It demonstrates a simple payment workflow using Stripe for payment processing and Material UI for the frontend.

## Features
- Go backend with Stripe integration
- React frontend with Material UI
- Secure handling of API keys using .env files
- Example of full-stack development with Go and React

## Project Structure

```
app/
  backend/
    server.go         # Go backend server
    .env              # Backend environment variables (not committed)
    go.mod            # Go module file
    .gitignore        # Ignores .env and other files
  frontend/
    src/              # React source files
    public/           # Static files
    .env              # Frontend environment variables (not committed)
    package.json      # React dependencies and scripts
    .gitignore        # Ignores .env and other files
```

## Getting Started

### Backend (Go)
1. Install Go (https://go.dev/dl/)
2. In the `backend` folder, create a `.env` file with your Stripe secret key:
   ```
   STRIPE_SECRET_KEY=sk_test_...
   ```
3. Run the backend server:
   ```
   go run server.go
   ```
   The server will start on `localhost:4242`.

### Frontend (React)
1. Install Node.js and npm (https://nodejs.org/)
2. In the `frontend` folder, create a `.env` file with your Stripe publishable key:
   ```
   REACT_APP_STRIPE_PUBLISHABLE_KEY=pk_test_...
   ```
3. Install dependencies:
   ```
   npm install
   ```
4. Start the React app:
   ```
   npm start
   ```
   The app will run on `localhost:3000` and proxy API requests to the Go backend.

## Notes
- This project is for learning purposes only. Do not use real Stripe keys in production.
- Make sure `.env` files are listed in `.gitignore` and never committed to version control.

## License
MIT
