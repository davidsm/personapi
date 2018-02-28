# Hämta och kompilera

Det finns flera sätt att hämta och kompilera tjänsten.
Vad som är enklast beror lite på vad man har installerat.

## go get

Om man redan har installerat och konfigurerat go kan man hämta och
kompilera direkt från repot. Kör följande kommando: `go get github.com/masenius/personapi`

Detta hämtar repot, och bygger en binär till `$GOPATH/bin/personapi[.exe]`.
Eftersom repot är privat behöver man ha tillgång till det på sitt github-konto.

## Bygg i Docker-container

Detta fungerar utan att ha go installerat, men man behöver istället förstås ha Docker.
Klona repot, och bygg från Docker-filen som finns tillgänglig: `docker build --rm -t personapi .`

Denna kan sedan köras med `docker run --rm -p 3000:3000 personapi`. Tjänsten kan sedan nås på
http://localhost:3000 på Linux, eller http://ip-till-vm:3000 på Mac eller Windows

## Bygg själv från lokal kopia
Klona repot till `$GOPATH/src/github.com/masenius/personapi`. Bygg med `go install`. Resultatet är det samma som `go get`.

## Kompilera och kör direkt
Från en lokal kopia, kör `go run server.go`. Smidigt vid utveckling om man vill testa en ändring.

## Bygg minimal Docker-avbild
I och med att go bygger statiskt länkade binärer utan externa beroenden kan man köra dem i väldigt grundläggande containers
som blir minimalt större än binärens storlek.

Bygg en binär för 64-bitars Linux. Om man sitter på Windows kan detta göras med följande batch-skript

``` batchfile
@echo off

SETLOCAL
set GOARCH=amd64
set GOOS=linux
go build -v %*
ENDLOCAL
```

Spara som t.ex. go-cross.bat och kör som `go-cross .` (i projektets rot).

Sitter man redan på 64-bitars Linux räcker `go build .`. På Mac bör `GOOARCH=amd64 GOOS=linux go build .` fungera.

När detta är klart ska det finnas en binär, `personapi` i rotmappen. Bygg en Docker-avbild med denna med `docker build --rm -t personapi-tiny --file Dockerfile.prebuilt .`. Kör med `docker run --rm -p 8080:8080 personapi-tiny`.

# Köra
Kör `personapi --help` för att se alternativ

```
  -bind string
    	Bind to address. Default is empty, meaning 0.0.0.0
  -logfile string
    	Log to file path. If not specified, log to stdout
  -port int
    	Port to use (default 8080)
  -seed int
    	Specify seed for the random generator. 0 means seed with current time. Not including this argument has the same effect as 0
```

Seed till slumpgeneratorn (valfri siffra annat än 0) kan sättas om man vill ha reproducerbara resultat, det vill säga samma requests i samma ordning ger samma resultat varje gång programmet körs.
