if not exist bin mkdir "bin"
go build -o bin/bankapiservice.exe src/main.go src/server.go src/client.go src/confighelper.go 
copy configs\*.* "bin"
