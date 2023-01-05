setup:
	docker run -p 8080:8080 -d pkya/304:deckofcards

run:
	go run main.go