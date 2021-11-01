# Pluggedin

Chat en tiempo real 

### Meta actual

Tener un room para chatear, encriptacion RSA, usar websockets con SSL

### Pasos futuros

Agregar contactos
Tener diferentes rooms (crear id de chats, cada chat es un websocket)
Room para solo dos personas
Usar Casbin para autorizacion

### Configuración

1. Crear una carpeta que posteriormente sera usada para la persistencia de la base de datos

```
mkdir ${HOME}/postgres-data/
```

2. Desplegar un contenedor con docker para Postgres

```
docker run -d 
	--name jaoksdb 
	-e POSTGRES_PASSWORD=123456 
	-v ${HOME}/postgres-data/:/var/lib/postgresql/data 
        -p 5434:5432
        postgres
```

En caso tener un contenedor ya configurado ejecutar lo siguiente
```
docker start jaoksdb
```

3. Correr el servidor

```
go run server.go
```

4. Pruebas de Web Socket utilizando wscat

```
wscat -c ws://localhost:1323/chats/:id
```

Reemplazar :id por un número entero
