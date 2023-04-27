rem sc delete bankapiservice
sc create bankapiservice binPath= "C:\Services\bankapiservice\bankapiservice.exe" start= demand 
sc description bankapiservice "API for mocking payment system" 
