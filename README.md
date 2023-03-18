# Proyecto Library-Backend

## Iniciar el proyecto

1. ```clon git clone https://github.com/Daizaikun/Back-Library```

2. Configurar las variables de entorno en el archivo .env si (esto es opcional si se usa los archivos docker por defecto)

3. ```Usar docker-compose upp``` si cambia los valores del .env por defecto debe actualizar el archivo docker-compose

4. El proyecto por defecto se ejecuta en el puerto 8080

Para despegarlo en producción recomiendo usar [fly.io](http://fly.io)

## Iniciar el proyecto en desarrollo

### Requisitos

- tener ejecutando una base de datos en postgres
- configurar el archivo .env en base a .env.example
  
#### Iniciar

1. ```clon git clone https://github.com/Daizaikun/Back-Library```

2. ```go mod tidy```

3. ```go run .```

Recomiendo usar [air](https://github.com/cosmtrek/air) para mejorar la experiencia de desarrollo
