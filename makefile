build_run:
		docker build -t lagashcol/go-comics:dev .
		docker run -it -p 8080:8081 lagashcol/go-comics:dev
build:
		docker build -t lagashcol/go-comics:dev .
run:		
		docker run -it -p 8080:8081 lagashcol/go-comics:dev
run_into:
		docker run -it lagashcol/go-comics:dev /bin/bash
run_test:
		docker run -it lagashcol/go-comics:dev go test