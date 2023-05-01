if not exist bin mkdir "bin"
if not exist bin/configs mkdir "bin\configs"
go build -o bin/bankapiservice.exe src/main.go src/server.go src/client.go src/confighelper.go && copy configs\*.* "bin" && move bin\appsettings.json "bin\configs\appsettings.json"
