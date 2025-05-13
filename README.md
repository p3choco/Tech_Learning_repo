**Zadanie 1** Docker

:white_check_mark: 3.0 obraz ubuntu z Pythonem w wersji 3.10 [Link do commita 1](https://hub.docker.com/r/pbednarski/kotlin-docker-example)

:white_check_mark: 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem SQLite w ramach projektu na Gradle (build.gradle) [Link do commita2 ](https://hub.docker.com/r/pbednarski/kotlin-docker-example)

:white_check_mark: 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC [Link do commita 3](https://hub.docker.com/r/pbednarski/kotlin-docker-example)

:white_check_mark: 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle [Link do commita 4](https://hub.docker.com/r/pbednarski/kotlin-docker-example)

:white_check_mark: 5.0 dodać konfigurację docker-compose [Link do commita 5](https://hub.docker.com/r/pbednarski/kotlin-docker-example)

Kod: master

**Zadanie 2** Należy stworzyć aplikację na frameworku Play w Scali 3

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo/commit/af504428353d2dc30262d3736ee0a17bb86e82e5)

:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy [Link do commita2 ](https://github.com/p3choco/Tech_Learning_repo/commit/40f6087ac8e66d45328f7d212b08f795718f00f2)

:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo/commit/af504428353d2dc30262d3736ee0a17bb86e82e5)

:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo/commit/d391d8af5375e487de35dd2631e72939e15a0a2d)

:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD
 [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo/commit/d391d8af5375e487de35dd2631e72939e15a0a2d)


Kod: scala

**Zadanie 3** Kotlin

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo/commit/79e81c48abce6b8c924b8621fcd5a55c368a7a14)

:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota) [Link do commita2 ](https://github.com/p3choco/Tech_Learning_repo/commit/7515ad9a0fd64a3db60eea2ae7a23864abfa1c06
)

:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo/commit/f93c438fe50e1b1ded3dc49ce4f96339b3a64b5e
)

:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo/commit/f0d40983164238caa5dbd847b0a611d09123e9c1
)

:white_check_mark: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,
Webex
 [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo/commit/763dcb5168ae8f9d3c2a6af8883977ab822e441e)




https://github.com/user-attachments/assets/c5e1152c-99b6-4050-ba53-fcd8fc2216ab


Kod: kotlin


**Zadanie 4** Go

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do
stworzenia 5 modeli, gdzie pomiędzy dwoma musi być relacja. Należy
zaimplementować proste endpointy do dodawania oraz wyświetlania danych
za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym
jednak pozostać przy sqlite.

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
miała kontroler Produktów zgodny z CRUD [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo/commit/ecaf27663690db6cba6b881ec902a93b4bcee7e8
)

:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
listy) [Link do commita2 ](https://github.com/p3choco/Tech_Learning_repo/commit/6f2c2f0cb8f8781eded33b744acd9524f3a9ff91
)

:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo/commit/34d847137ee7f94a6d72375a74a168805c4ebffd
)

:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią,
a produktem [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo/commit/1cd2ace9d146b48967cb4e55251ff261741ca90c
)

:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y
 [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo/commit/3b55f731884e8be36adba7f6fce2650a20781527)


Kod: Go



**Zadanie 5** Frontend

Należy stworzyć aplikację kliencką wykorzystując bibliotekę React.js.
W ramach projektu należy stworzyć trzy komponenty: Produkty, Koszyk
oraz Płatności. Koszyk oraz Płatności powinny wysyłać do aplikacji
serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach
z aplikacji serwerowej. Aplikacja serwera w jednym z trzech języków:
Kotlin, Scala, Go. Dane pomiędzy wszystkimi komponentami powinny być
przesyłane za pomocą React hooks.

:white_check_mark: 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
Produktach powinniśmy pobierać dane o produktach z aplikacji
serwerowej; [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo/commit/c7091c51ccc7eeaa76341403a229fa0203ab89f4)

:white_check_mark: 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing [Link do commita 2](https://github.com/p3choco/Tech_Learning_repo/commit/3d66229a8263cd5ef70943d3ed211d092e083e54)

:white_check_mark: 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
pomocą React hooks [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo/commit/1c1a10e831cd669f0685015a0ccc14d5b1f96c1f)

:white_check_mark: 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
kliencką na dockerze via docker-compose [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo/commit/5d59c24ba259d44bbd893c62a3f14d343b428161)

:white_check_mark: 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo/commit/5d59c24ba259d44bbd893c62a3f14d343b428161)


https://github.com/user-attachments/assets/dd2fbe6d-0466-494f-8aa8-2f0ca3166a3c


Kod: Frontend


**Zadanie 6** Testy

Należy stworzyć 20 przypadków testowych w jednym z rozwiązań:

- Cypress JS (JS)
- Selenium (Kotlin, Python, Java, JS, Go, Scala)

Testy mają w sumie zawierać minimum 50 asercji (3.5). Mają również
uruchamiać się na platformie Browserstack (5.0). Proszę pamiętać o
stworzeniu darmowego konta via https://education.github.com/pack.

:white_check_mark: 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
(Kotlin, Python, Java, JS, Go, Scala) [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo/commit/ed9aff04c1f69e312fc7d1ed672f88ac61d2d4a9)

:white_check_mark: 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
asercji [Link do commita 2](https://github.com/p3choco/Tech_Learning_repo/commit/ed9aff04c1f69e312fc7d1ed672f88ac61d2d4a9)

:white_check_mark: 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
projektu z minimum 50 asercjami [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo/commit/ed9aff04c1f69e312fc7d1ed672f88ac61d2d4a9)

:white_check_mark: 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
minimum jednym scenariuszem negatywnym per endpoint [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo/commit/ed9aff04c1f69e312fc7d1ed672f88ac61d2d4a9)

:white_check_mark: 5.0 Należy uruchomić testy funkcjonalne na Browserstacku [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo/commit/ed9aff04c1f69e312fc7d1ed672f88ac61d2d4a9)


Kod: Tests


**Zadanie 7** Sonar

Należy dodać projekt aplikacji klienckiej oraz serwerowej (jeden
branch, dwa repozytoria) do Sonara w wersji chmurowej
(https://sonarcloud.io/). Należy poprawić aplikacje uzyskując 0 bugów,
0 zapaszków, 0 podatności, 0 błędów bezpieczeństwa. Dodatkowo należy
dodać widżety sonarowe do README w repozytorium dane projektu z
wynikami.

:white_check_mark: 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
hookach gita [Link do commita 1](https://github.com/p3choco/Tech_Learning_repo_SERVER/commit/1bdfc48edabe503ab95920b31b4edb32d8cd2203)

:white_check_mark: 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
aplikacji serwerowej) [Link do commita 2](https://github.com/p3choco/Tech_Learning_repo_SERVER/commit/bbef83bc2ae4e495e5c4b1015105b09a4a22be2a)

:white_check_mark: 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
aplikacji serwerowej) [Link do commita 3](https://github.com/p3choco/Tech_Learning_repo_SERVER/commit/bbef83bc2ae4e495e5c4b1015105b09a4a22be2a)

:white_check_mark: 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
w kodzie w Sonarze (kod aplikacji serwerowej) [Link do commita 4](https://github.com/p3choco/Tech_Learning_repo_SERVER/commit/bbef83bc2ae4e495e5c4b1015105b09a4a22be2a)

:white_check_mark: 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie
aplikacji klienckiej [Link do commita 5](https://github.com/p3choco/Tech_Learning_repo_CLIENT/commit/f73979ae09d2a22147b7dd373ca934df5043bc10)


Kod: 

[Server](https://github.com/p3choco/Tech_Learning_repo_SERVER)

[Client](https://github.com/p3choco/Tech_Learning_repo_CLIENT) 


