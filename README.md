# Go Location Temp

## How can I test online?

You can test the application online using the following link:
https://go-location-temp-700422746697.us-central1.run.app/temperature/{cep_numbers}
Replace `{cep_numbers}` with the desired Brazilian postal codes (CEPs) numbers. For example:

https://go-location-temp-700422746697.us-central1.run.app/temperature/13484000

## Instructions:

1. Clone the repository:

```zsh
git clone https://github.com/pablohen/go-location-temp.git
```

2. Add a valid `WEATHER_API_KEY` to the `docker-compose.yaml` file:

```yaml
environment:
  - WEATHER_API_KEY=your_api_key
```

3. Run the Docker container:

```zsh
docker compose up -d
```
