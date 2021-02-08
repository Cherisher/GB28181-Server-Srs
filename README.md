# GB28181-Server-Srs


## Docker
```shell
docker build -t gb_web:v1 .
docker run --name gb_web -itd -p 28080:80 gb_web:v1
# 当前配置运行
docker run --name gb_web -itd -v ${PWD}/conf/app.json:/data/conf/app.json -p 28080:80 gb_web:v1 
```
