#Multi Stage Build

## Stage 1: Builder
``` dockerfile
FROM golang:1.24-alpine AS builder
```
Kompiliert den GO-Code. Enthältden kompletten Go-Compiler udn alle Build-Tools

### Stage 2: Final Imgae
````dockerfile
FROM alpine:latest
COPY --from=builder /app/api .
```
Kopiert nur die fertige Binary aus Stage 1. Enthält keinen Compiler mehr

#CGO_Enabled=0

CGO_Enabled=0 deaktiviert C-Bindings in Go.
Die Binary wird statisch gelinkt, das heisst es ist keine externen Libraries nötig
- Läuft in minimalen Containern wie Alpine oder sogar `scratch`
- Ohne diese Einstellung würde die Binary, C-Libraries erwarten die im
  Mini-Container nicht vorhanden sind was zu einem Crash beim Start führt

## Image Size

Single-Stage(golang:alpine) 300 MB
Multi-Stage (alpine-final) 15MB

Multi-Stage ist somit viel kleiner

# Test Egebnisse

## Produkt erstellen 
POST /products {"name":"Apple","price":1.99} → 201 Created, ID=1
POST /products {"name":"Banana","price":0.49} → 201 Created, ID=2
POST /products {"name":"Cherry","price":3.99} → 201 Created, ID=3 

##Liste 
GET /products → 200 OK, Liste mit 3 Produkten

##Update
PUT /products/1 {"name":"Green Apple","price":2.49} → 200 OK

##DELETE
DELETE /products/2 → 204 No Content

##Persistenz
Nach docker compose down und compose up:
Daten waren noch vorhanden
