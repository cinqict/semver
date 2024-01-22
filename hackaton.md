
# Hackaton: Semver cli

## ToDo Hackaton

Prep:

- [x] eigen idee van functies ()
- [ ] Cli basics
- [.] Linting
- [x] Unit test: test repo
- [x] Unittest: example
- [x] Create go files
- [x] Create structs
- [ ] op cinq github
- design: @Bouke hoe doen jullie developers dat?

## Prep collega's

- Installeer Golang v1.21 (via je app manager of via https://go.dev/doc/install)
- Handige IDE:  Goland JetBrains 30 days free trail: https://www.jetbrains.com/go/download
  - Zet GOROOT (path naar go binary)
- Chat GPT account: https://chat.openai.com/
- Basic: Eventueel oefeningen/lezen over: Go functions, structs
- Pro: We zullen de go-git library veel gebruiken: https://github.com/go-git/go-git

## Agenda

- Intro: idee, show basics
- Design functions (allemaal)
    - input, output, go file, pseudo algorithm
- Tweetallen: 
  - kies een functie
  - Ci: maak een github action met
    - linting
    - run tests
    - trigger on pr
    - merge only allowed when pipeline succeeds
  - Dockerfile
- Zorg dat linting + tests altijd goed zijn

## uitnodiging

Yo, De laatste tijd ben ik wat met programmeren bezig en zo was ik aan een CLI tooltje begonnen om een semantic versioning tag te genereren op basis van de git history.

En toen dacht ik:
Als we dat nou eens met zo'n alle doen en daarna in tweetallen - plus Chatgpt als senior developer naast je - ieder een onderdeeltje maken?  Volgens mij is dat best wel leuk.
Dus daarom wil ik een hackaton organiseren!

Nu had ik het liefst eens iets anders dan donderdag avond, maar dat blijkt niet echt haalbaar. Dus toch maar weer een donderdag. Mogen jullie wel kiezen welke en meeste stemmen gelden.

/poll "Hackaton avond" "donderdag 25 jan" "donderdag 1 feb"