<i>foxypad</i> - simple online scratchpad with authorization wriiten in Go<br/>
<br/>
Copyright (C) 2018 Alexey "FoxyLab" Voronin<br/>
Email: support@foxylab.com<br/>
This software is licensed under the GPL v3.0 License.<br/>
<br/>
Settings (in file <b>foxypad.go</b>):<br/>
```
const filename string = "foxypad.txt" //data filename
const port string = ":8888"           //port number
const rows string = "20"              //rows number
const login string = "user"           //login
const password string = "password"    //password
```
Build:<br/>
<b>go get -u github.com/gin-gonic/gin</b> (Gin framework installation)<br/>
<b>go build foxypad.go</b><br/>
Use - on home server or VPS<br/>
<br/>
<i>foxypad</i> - простой онлайн-блокнот с авторизацией на Go<br/>
<br/>
Copyright (C) 2018 Алексей "FoxyLab" Воронин<br/>
Электронная почта: support@foxylab.com<br/>
Это программное обеспечение распространяется под лицензией GPL v3.0.<br/>
<br/>
Настройки (в файле <b>foxypad.go</b>):<br/>
```
const filename string = "foxypad.txt" //имя файла с данными
const port string = ":8888"           //номер порта
const rows string = "20"              //кол-во строк
const login string = "user"           //логин
const password string = "password"    //пароль
```
Компиляция:<br/>
<b>go get -u github.com/gin-gonic/gin</b> (установка фреймворка Gin)<br/>
<b>go build foxypad.go</b><br/>
Использование - на домашнем сервере или VPS



