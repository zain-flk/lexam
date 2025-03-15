
**Developer Guide: Local Setup**

## Prerequisites
Before setting up the project locally, ensure you have the following installed:
- **Go** (latest stable version)
- **Node.js** (LTS version)
- **npm** or **yarn**

## Clone the Repository
```sh
git clone https://github.com/zain-flk/lexam.git
cd lexam
```

## Install Dependencies
```sh
go mod tidy
npm install
```

## Set Up Environment Variables
Create a `.env` file in the project root with the following variables:
```env
ACCESS_KEY=your-access-key
SECRET_KEY=your-secret-key
R2_ENDPOINT=your-r2-endpoint
BUCKET_NAME=your-bucket-name
DEST_PATH=your-destination-path
```

## Run the Project
```sh
go run main.go
```
```

## Troubleshooting
- Ensure all dependencies are installed correctly.
- Verify environment variables are correctly set.


