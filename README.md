Modelado y Programación
=======================

Proyecto 1: Chat
-----------------------

### Fecha de entrega: viernes 25 de septiembre, 2026

El proyecto consta de dos ejecutables, de los cuales
uno es para el Servidor, y el otro para el Cliente.

Para ejecutar el servidor 

```
$ ./Servidor
```
Para ejecutar el cliente
```
$ ./Cliente <IP:HOST> <NOMBRE_DE_USUARIO>
```

------------------------
### Uso del cliente

Para cambiar el estado de un cliente.
```
$ STATUS <estado>
```
Posibles estados:
- ACTIVE (default)
- BUSY
- AWAY

Pedir la lista de usuarios de la sala general.
```
$ USERS
```

Enviar un texto público.
```
$ <mensaje>
```

Enviar un texto privado a otro cliente.
```
$ TEXT <nombre> <mensaje>
```

Para desconectarse.
```
$ exit
```


  
