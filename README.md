

# TA2
Tarea academica 2 del curso Programación Concurrente y Distribuida

Pre requisto fundamental (compilador GCC) Instalarlo antes de ejecutar el analisis de dataset:
Se debe tener go para escritorio instalado

https://github.com/jmeubank/tdm-gcc/releases/download/v9.2.0-tdm64-1/tdm64-gcc-9.2.0.exe

Para ejecutar el servidor backend local en go situarse en la raiz del proyecto y en cmd escribir
go run cmd/gopherapi/main.go

Para ejecutar el servidor frontend en node.js situarse en la raiz del proyecto y escribir npm install (para los prerequisitos) y luego
npm start 

Para ejecutar el conversor a csv del backend en go, situarse en la raiz del proyecto y ejecutar
go run getrequest.go

Para ejecutar el análisis de la dataset de Calidad de vinos, ejecutar el siguiente comando
go run winetest.go

Video demo:
https://www.youtube.com/watch?v=1zqRPRIf8DI

