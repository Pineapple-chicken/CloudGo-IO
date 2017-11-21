# CloudGo-IO
Assignment for service computing : CloudGo-IO

## 要求1：支持静态文件服务：
我用了Negroni的logging中间件，得到如下结果：
```
[negroni] listening on :8080
[negroni] 2017-11-21T22:45:46-08:00 | 200 | 	 1.896475ms | localhost:8080 | GET / 
[negroni] 2017-11-21T22:45:46-08:00 | 200 | 	 64.387µs | localhost:8080 | GET /index.css 
[negroni] 2017-11-21T22:45:46-08:00 | 200 | 	 24.369µs | localhost:8080 | GET /cloud_go.js 
```
可以看到对客户端的请求做出了响应，返回了主页，以及静态目录的css文件和js文件。

## 要求2：支持简单 js 访问
访问localhost:8080/api会返回JSON数据，可以在cloud_go.js中使用Jquery的Ajax方法请求这个地址，得到返回的数据：
```
1268
Pineapple-chicken
```
可见其支持简单js访问。

## 要求3：提交表单，并输出一个表格
在服务端程序可看到，请求了/api路径：
```
[negroni] 2017-11-21T22:49:54-08:00 | 200 | 	 101.344µs | localhost:8080 | GET /api
```
同时输出了表格。

## 要求4：对<code>/unknown</code>给出开发中的提示，返回码<code>5xx</code>
在访问localhost:8080/unknown时，浏览器会出现<code>501 not implement</code>。
服务端也会显示出现了501错误：
```
[negroni] 2017-11-21T23:02:39-08:00 | 501 | 	 57.614µs | localhost:8080 | GET /unknown
```
