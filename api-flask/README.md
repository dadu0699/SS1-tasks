# API
## Start
Existen dos maneras de poner en marcha nuestra api con flask siendo por medio de variables de entrono o por código.

### VENV
La primera opción es mediante la declaración de las variables de entorno, las cuales se pueden declarar en nuestro entorno virtual en los archivos `activate`, `actívate.bat` y `Activate.ps1`.

```bash
# venv/bin/activate
export FLASK_APP="app.py"
export FLASK_ENV="development"
```

```bat
rem .\venv\Scripts\Activate.bat
set "FLASK_APP=run.py"
set "FLASK_ENV=development"
```

```ps1
# .\venv\Scripts\Activate.ps1
$env:FLASK_APP = "app.py"
$env:FLASK_ENV = "development"
```

Al declarar nuestras variables de entorno es posible correr el servidor mediante el comando:
```bash
flask run 

# different port
flask run --host 0.0.0.0

# All hosts
flask run --port 6000
```

### By code
La opción mediante código se realiza creando el main en Python en el cual se llama al servidor mediante su método run, también se le pueden pasar parámetros para su ejecución.

```py
if __name__ == '__main__':
    app.run(debug=True)
```