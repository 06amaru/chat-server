# Pluggedin

Chat en tiempo real con encriptación E2E para conversaciones efímeras mediante un protocolo híbrido (AES y ECIES)

### Meta actual

Tener un cola para mensajería asíncrona (caso usuario offline). Actualmente dos usuarios hablan sin problemas online, pero si un se encuentra offline y recibe mensajes durante ese tiempo necesita conseguir los mensajes desde la cola.

### Pasos futuros

Implementar un socket para cada usuario donde se mande notificaciones acerca de nuevos mensajes
### Configuración

1. Crear una carpeta que posteriormente sera usada para la persistencia de la base de datos

```
mkdir ${HOME}/postgres-data/
```

2. Desplegar un contenedor con docker para Postgres

```
docker run --name jaoksdb -e POSTGRES_PASSWORD=123456 -p 5434:5432 -d postgres
```

En caso tener un contenedor ya configurado ejecutar lo siguiente
```
docker start jaoksdb
```

3. Correr el servidor

```
go run server.go
```
