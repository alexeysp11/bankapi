# bankapi

`bankapi` is a part of mocking payment system, that establishes connection with [banking](https://github.com/alexeysp11/banking) (a backend server of the payment system), and redirects end users' requests to it. 

## How to configure and run the application 

1. Edit config file 

There's a `appsettings.json` file inside the `config` folder. The file contains configurational settings for the application, for example:
```JSON 
{
    "Environment": "test", 
    "ServerHost": "127.0.0.1", 
    "ServerPort": "8081", 
    "ServerReadTimeout": 10,
    "ServerWriteTimeout": 10,
    "BankCoreAddress": "http://127.0.0.1:8080",
    
    "Atm": ["v1", "v2"],
    "Eftpos": ["v1", "v2"], 

    "HttpPathsAtm": [
        "/pin/enter/", 
        "/pin/change/",
        "/balance/get/",
        "/cash/deposit/",
        "/cash/withdraw/",
        "/transfer/tobankaccount/",
        "/transfer/tophonenumber/",
        "/transfer/fps/"
    ],
    "HttpPathsEftpos": [
        "/pin/enter/",
        "/transfer/"
    ]
}
```

2. Build the application 

In order to build the app, use the following command: 
```
build.cmd
```

If you want to run the application only in test environment, you can skip this step. 

3. Running the application 

If you built your application, execute: 
```
"bin/bankapiservice.exe"
```

If you skipped the previous step, use the following command: 
```
run.cmd
```

<!--
If the app is already built, you can deploy it as Windows Service:
1. Create a folder for hosting Windows Service (e.g. `D:\Services\bankapi`); 
2. Copy all the files from `bin` folder into the specified folder; 
3. In `bin` folder you will also be able to find batch files to install and uninstall Windows Service (filenames: `install.bat` and `uninstall.bat` respectively). Make sure that all of these files are copied into a folder for hosting the service;
4. Open file `install.bat` with any text editor and check if the parameter `binPath` is equal to the path to the folder for hosting Windows Service; 
5. Run `install.bat` as administrator. 
-->

## Testing 

1. Test the running server 

In order to test server, run the following command in command line:  
```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8081/banking/eftpos/v1/pin/enter/
```

## How to contribute 

In order to find information on how this app could be improved, please read [todo](docs/todo.md) file. 
