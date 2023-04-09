
**CICD GO Exercise**

Mittels Mux in Go eine Rest Schnittstelle test driven geschrieben.

Das Tutorial war ganz gut und es hat keine gröberen Probleme gegeben.

Habe dann dazu noch eine Schnittstelle gebaut, die einen alle Produkte je nach Übergabewert sortiert übergibt.
Hierbei habe ich dann erst bemerkt, dass ich eigentlich die Programmiersprache Go gar nicht spreche und es hat dann mindestens doppelt so lange gebraucht ein kleines Feature zu implementieren, wie das ganze Tutorial durchzuarbeiten.
Nun kann man aber mit dem folgenden Endpoint alle Produkte nach dem Preis sortiert auf und absteigend ausgeben lassen.

    /sort/{asc/desc}

Mittels nvim und Makefile war es auch möglich schnell zu entwickeln.
Es war auf jeden Fall interessant, test driven zu entwickeln, da ich ja auch bei meinem Feature die Tests geschrieben habe, bevor ich etwas implementiert habe. Da ist anfangs der Frust zwar immer groß, allerdings die Freude wenn es funktioniert umso größer.



