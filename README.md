# Japan Prefectures and Cities API

This API provides information about Japan's prefectures, their cities, and train lines based on the prefecture ID.

## Endpoints

### Get Cities by Prefecture ID
<h3 align="left">Get all cities in a specific prefecture based on the prefecture ID.</h3>
<p align="left">
</p>

<h3 align="left">Languages and Tools:</h3>
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> </p>

```curl
GET /cities-by-prefecture-id?prefecture_id=<PREFECTURE ID>
```

Example Response:

```json
{
  "cities": [
    {
      "id": "1",
      "prefecture_id": 1,
      "city_en": "Tokyo",
      "city_ja": "東京",
      "special_district_ja": "特別区"
    },
    {
      "id": "2",
      "prefecture_id": 1,
      "city_en": "Yokohama",
      "city_ja": "横浜",
      "special_district_ja": ""
    }
    // ... (other cities)
  ]
}
```
<h3 align="left">Get All Japan Prefectures</h3>
Get a list of all Japan prefectures.

```curl
GET /prefectures
```

Example Response:

```json
{
  "prefectures": [
    {
      "id": 1,
      "prefecture_kanji": "東京都",
      "prefecture_romaji": "Tokyo",
      "prefecture_kana": "とうきょうと"
    },
    {
      "id": 2,
      "prefecture_kanji": "神奈川県",
      "prefecture_romaji": "Kanagawa",
      "prefecture_kana": "かながわけん"
    }
    // ... (other prefectures)
  ]
}
```
<h3 align="left">Get Train Lines By Prefecture</h3>
Get all train lines in a specific prefecture based on the prefecture ID.

```curl
GET /lines-by-pref?prefecture_id=<PREFECTURE ID>
```

Example Response:

```json
{
  "prefecture_lines": {
    "id": 1,
    "name": {
      "en": "Tokyo Metro",
      "ja": "東京メトロ"
    },
    "lat": 35.681236,
    "lng": 139.767125,
    "zoom": 11,
    "stations": [
      {
        "id": "T01",
        "name": {
          "en": "Shinjuku",
          "ja": "新宿"
        },
        "lat": 35.6895,
        "lng": 139.7009
      },
      {
        "id": "T02",
        "name": {
          "en": "Shibuya",
          "ja": "渋谷"
        },
        "lat": 35.6586,
        "lng": 139.7016
      }
      // ... (other train stations)
    ]
  }
}
```
<h3 align="left">How to Use</h3>

1. Make a GET request to the desired endpoint using the provided URLs.

2. Replace <PREFECTURE ID> with the ID of the specific prefecture you want to query.

<h3 align="left">Deployment</h3>
This API is deployed as an AWS Lambda function with the AWS API Gateway serving as the REST API endpoint.
To build the AWS Lambda function and create the function ZIP file, follow these steps:

1. Use Docker Compose to access the container environment:

```bash
  docker-compose exec api-app bash
```

2. Build the AWS Lambda function binary:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
```

3. Create the function ZIP file:

```bash
zip ./function.zip main
```

The function.zip file contains the AWS Lambda function binary and can be uploaded to AWS Lambda.

<h3 align="left">Data Sources</h3>
The data for Japan's prefectures, cities, and train lines is sourced from the following JSON files:

* prefectures.json: Contains information about Japan's prefectures.
* cities.json: Contains information about cities in each prefecture.
* prefectures-stations.json: Contains information about train lines in each prefecture.

<h3 align="left">Technologies Used</h3
                                    
* Go programming language
* AWS Lambda
* AWS API Gateway
* GitHub Actions for continuous integration
<h3 align="left">Contributions</h3>
Contributions to this project are welcome! If you find a bug or want to add new features, feel free to submit a pull request.

<h3 align="left">License</h3>
This project is licensed under the MIT License.
