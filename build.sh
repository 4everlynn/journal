release='./release'
CGO_ENABLED=0 GOOS=windows  GOARCH=amd64 go build -o "${release}/journal.exe"
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o "${release}/journal_linux"
go build -o "${release}/journal"