# navodila za uporabo redovalnice
Za podrobnejša navodila in godoc dokumentacijo poženi:
```
go doc -all redovalnica
```
Več na: https://pkg.go.dev/github.com/makov3c/dn5ps@v0.0.0-20251202183731-53affebc548f/redovalnica

Najprej poženemo `go.build` nato pa poganjamo program s poljubnimi zastavicami.

## zastavica najmanjOcen
```go
./dn5ps --najmanjOcen 3
```
Povprečje bomo izpisali le, če ima študent najmanj 3 ocene (sicer javimo napako)

## zastavica minOcena
```go
./dn5ps --minOcena 7
```
Ker imamo radi dobre študente, bomo vpis ocene dovolili le, če je najmanj 8 (sicer javimo napako)

## zastavica maxOcena
```go
./dn5ps --maxOcena 9
Ker ne želimo, da imajo študenti predobre ocene, jih omejimo z 9 (sicer javimo napako)