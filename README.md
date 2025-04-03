## Setup Instructions

### 1. Clone the Repository
Run the following command to clone the repository:
```bash
git clone git@bitbucket.org:Amartha/go-dlq-retrier.git
cd go-dlq-retrier
```

### 2. Copy and Configure the Environment File
Duplicate the `.env.sample` file as `.env` and make any necessary adjustments:
```bash
cp .env.sample .env
```

### 3. Connect to OpenVPN
Ensure you are connected to OpenVPN before running the application.

### 4. Run the Application
Execute the following command to start the application:
```bash
go run cmd/main.go
