### Funciones de Lambda con Go
Para poder realizar una función Lambda con el lenguaje Go es necesario tener en cuenta si la misma va a ser consumida mediante de alguna de las formas que se presentan a continuación, esto debido a que se codifican con algunas diferencias.
- [API Gateway](./lambda-apigateway-go)
- [AWS SDK](./lambda-go)

Luego es necesario tener en cuenta que el archivo a subir en AWS es un compilado para la arquitectura seleccionada al momento de la creación de la función y para el sistema operativo Linux. Esto se puede realizar mediante la ejecución del siguiente comando:
```bash
GOOS=linux GOARCH=amd64 go build -o main
```

Es importante tener en cuenta que si estas en un sistema operativo Windows el comando anterior no funcionara completamente para esto es necesario cambiar y verificar las variables de entorno. Esto se puede realizar mediante los siguientes comandos:
```bash
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
go env GOOS GOARCH
go build -o main
```