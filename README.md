# sqliteDBPractice

Install Golang 

clone directory in GOPath

run commands:
go mod init 
go mod tidy 

start server: 

go run main.go 

can query database using query parameters 

/payments?date=<1642248000>&status=<failed/success>

without any query parameters returns entire payments table
