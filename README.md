# bankapi

`bankapi` is a part of mocking payment system, that establishes connection with [banking](https://github.com/alexeysp11/banking) (a backend server of the payment system), and redirects end users' requests to it. 

## How to use 

Run the app by typing the following command: 
```
_run.cmd
```

In order to build the app, use: 
```
_build.cmd
```

If the app is already built, you can deploy it as Windows Service:
1. Create a folder for hosting Windows Service (e.g. `D:\Services\bankapi`); 
2. Copy all the files from `bin` folder into the specified folder; 
3. In `bin` folder you will also be able to find batch files to install and uninstall Windows Service (filenames: `install.bat` and `uninstall.bat` respectively). Make sure that all of these files are copied into a folder for hosting the service;
4. Open file `install.bat` with any text editor and check if the parameter `binPath` is equal to the path to the folder for hosting Windows Service; 
5. Run `install.bat`. 

In order to test server, run the following command in command line (the command is depricated):  
```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/banking/notes
```

## How to contribute 

In order to, please read [todo](docs/todo.md) file. 
