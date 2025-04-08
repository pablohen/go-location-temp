# Go Location Temp

## Instructions:

1. Clone the repository:

```zsh
git clone https://github.com/pablohen/go-location-temp.git
```

2. Add a valid `WEATHER_API_KEY` to the `docker-compose.yaml` file:

```yml
environment:
  - WEATHER_API_KEY=your_api_key
```

3. Run the Docker container:

```zsh
docker compose up -d
```
