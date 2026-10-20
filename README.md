# varo-sample-app

App di prova per la dimostrazione end-to-end di [Varo](https://github.com/k4roshi/varo) (fase 0): ogni push su `main` passa dalla CI di Varo (`.varo/workflows/ci.yml`), viene costruito sul nodo e deployato senza interruzioni.

Endpoint: `/` (saluto con la versione), `/up` (health check), `/version`, `/slow?ms=` (richieste lunghe, per i test di carico).

Deployed by Varo from a push.
