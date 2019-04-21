# gowifiscanner
WiFiScanner in go fo learning purposes

## Arch

* Main program is a standalone server
* Server runs as a daemon and collects data
* Server makes data available over socket file and HTTPS
* Server transmits data using JSON format
* Most basic client is command line utility communicating over socket file
* Most basic client doesn't use any graphical library

## Database

* Data will be kept either in json files or sqlite
* DB structure is described in file DB.md


WiFi logger
Logger WiFi W oparciu o odczyt nazw sieci bezprzewodowych oraz pozycji GPS zapisywanie mocy sygnału wraz z pozycją. Można rozszerzyć o interfejs graficzny wizualizujący moc sygnału

Wymagania funkcjonalne:
Wybudzanie karty WiFi w trybie nasłuchu
Przechodzenie przez poszczególne kanały cyklicznie w krótkich okresach czasu (do ustalenia doswiadczalnie)
Pobranie pozycji GPS, nazwy sieci i siły sygnału
Identyfikacja dubletów sieci, zapis tylko raz
Możliwość zaprzestania pracy na żądanie użytkownika
Drukowanie listy sieci na żądanie użytkownika
Oczyszcześnie bazy na żądanie użytkownika
Połączenie z bazą danych
Odczyt z bazy danych
Zapis do bazy danych
Interfejs umożliwiający komunikację z użytkownikiem
Wymagania niefunkcjonalne:
Program powinien być podzielony na dwa elementy: serwer/daemon, który będzie dokonywać faktycznego "próbkowania" i zapisywać do bazy oraz interfejsu użytkownika, którym można połączyć się do daemona. Pozwala to na dowolną wymianę konsoli klienckiej: WWW, ncurses, nieinteraktywna komenda cli, graficzna aplikacja desktopowa. W najprostszej postaci powinien mieć interfejs konsolowy pozwalający na rozpoczęcie i zakończenie pracy, wydruk dotychczasowych zapisów, oczyszczenie bazy i zamknięcie programu. To byłaby wersja zarządzania do uruchomienia bezpośrednio na "loggerze". Musi zatem zajmować mało miejsca oraz zabierać mało mocy procesora (zużycie ewentualnej baterii). Powinien mieć prosty protokół komunikacji przez sieć internet używając szyfrowanego połączenia. Dla uproszczenia konsole klienckie muszą być uwierzytelnione ręcznie przez administratora loggera. Dane przesyłane przez internet powinny być jak najmniejsze, aby generować jak najmniejsze obciążenie dla sieci.


To generate self signed certificate without password:
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -subj '/CN=wifiscanner' -nodes

The pair in docs is for testing only
