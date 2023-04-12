## TPII Go

Podés accerder al proyecto también desde Github: https://github.com/DamianSformo/proyecto-entregable-final

Para empezar la ejecucion de la aplicación, primero ejecutamos el script que se encuentra en esta misma carpeta para crear las tablas en la respectiva base de datos. Luego en el archivo main.go completamos los siguientes campos que hacen referencia a la base de datos a utilizar:

	userdb
	passworddb
	portdb

Luego corremos el siguiente comando ubicados en la raiz del proyecto para iniciar el proyecto: 

    go run cmd/server/main.go

Y por último probamos los endpoints con la colección de postman que se encuentra también en esta misma carpeta.