### Download the quote sqlite DB file here:
https://www.dropbox.com/s/rqugvxwvc1aq8e3/quotes_v1.sqlite?dl=0

- 500k quotes from Goodreads
- Original source  > https://www.kaggle.com/datasets/manann/quotes-500k

### Run locally
```bash
export PORT=8080
go run main.go
```
Then open `http://localhost:8080/random`

### Build the image
```bash
docker build -t quotes-app .
```

### Run the container
```bash
docker run -it -p 8080:8080 -v $(pwd):/app quotes-app
```
