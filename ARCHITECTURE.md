# Architecture Doku

## Request Flow

HTTP Request
|
Router
|
Handöer
|
Store
|
Database

HTTP Request: Client sendet eine HTTP-Anfrage z.b GET
Router: Gorilla Max matcht die URL und leitet an den richtigen Handler weiter
Handler: Liest Request-Daten, ruft die Store-Methode auf, schreibt
JSON-Response
Store: Interface mit zwei Implementierungen: MemoryStore und PostgresStore
Database: PostgressStore speichert daten in PostgreSQL, MemoryStore im RAM

## MemoryStore vs PostgresStore

|    Vergleich   |     MemoryStore         |       PostgresStore       |
| Geschwindigkeit| Sehr schnell            | etwas langsamer           |
| Einstatz       | Tests, lokale Entw.     | Produktion                |
| Setup          | Kein Setup nütig        | Postgres Server nötig     |
| Skalierung     | nicht möglich           | Mehrere Instanzen möglich |
| Persitenz      | Daten weg nach Neustart | Daten bleben erhalten     |


## Wann wird welches verwendet ?

MemoryStore: Unit Tests, schnelle lokale Entwicklung ohne Datenbank
PostgresStore: Produktion, wenn Daten persistent gespeichert werden müssen
