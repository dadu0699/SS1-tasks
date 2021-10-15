# API
La funcionalidad principal del API es permitir utilizar las funcionalidades de `DetectLabels` y `CompareFaces` del servicio `rekognition` de [AWS]( https://aws.amazon.com).


### DetectLabels
Permite reconocer las etiquetas de una imagen enviada en base64. Esto por medio del método POST en el endpoint ` http://127.0.0.1:5000/detect-labels ` y una entrada JSON con la siguiente estructura.
```json
{
  "image": "<imagen_en_base64>"
}
```


### CompareFaces
Permite comparar rostros por medio de dos imágenes en base64 y un valor numérico de similitud. Esto por medio del método POST en el endpoint ` http://127.0.0.1:5000/compare-faces ` y una entrada JSON con la siguiente estructura.
```json
{
  "imagen1": "<imagen_en_base64>",
  "imagen2": "<imagen_en_base64>",
  "similitud": <valor_numerico>,
}
```


## Packages
Se requirió la instalación de tres paquetes los cuales son
-	[Gorilla/Mux](https://github.com/gorilla/mux): Permite implementar un enrutador de solicitudes y un despachador para hacer coincidir las solicitudes entrantes con su respectivo controlador.
-	[Go CORS handler](https://github.com/rs/cors): Es un controlador `net/http` que implementa la especificación W3 de intercambio de recursos de origen cruzado en Golang.
-	[GoDotEnv](https://github.com/joho/godotenv): Carga variables de entorno desde un archivo .env.
-	[AWS SDK for Go](https://github.com/aws/aws-sdk-go): Es el SDK oficial de AWS para el lenguaje de programación Go.

```
go get -u github.com/gorilla/mux
go get -u github.com/rs/cors
go get -u github.com/joho/godotenv
go get -u github.com/aws/aws-sdk-go
```


## .env File
```env
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
REGION_NAME=
```


## Start
Primero es necesario instalar los paquetes necesarios y luego se puede realizar la ejecución del API.
```bash
go mod tidy
go run .\cmd\server\main.go
```