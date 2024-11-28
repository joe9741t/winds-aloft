# winds-aloft
Easily calculate the best altitude to fly

## Problem Statement
Pilots can often choose the altitude at which they fly. Aviation weather sources provide the wind direction and speed at various altitudes. Pilots would like to quickly calculate the best altitude to fly in order to minimize headwind or maximize tailwind. This saves time and fuel. 

Possible data sources: 
https://aviationweather.gov/data/api/
https://aviationweather.gov/api/data/windtemp?region=all

## Formula
Reference: https://en.wikipedia.org/wiki/Headwind_and_tailwind
A = Angle from direction of travel
WS = Windspeed
Headwind = cos(A) * WS 

## Examples 
1) Airplane traveling north at an altitude where windspeed is 20kts @ 135 degrees (wind out of the southwest): 20 × cos(135) = -14.14kts of tailwind. 
2) Airplane traveling north at an altitude where windspeed is 20kts @ 45 degrees (wind out of the northeast): 20 x cos(45) = 14.14kts of headwind. 

## User Interface
Pilot would enter the nearest airport identifier, for example: KLJF. 
System displays the best altitude to fly, for example: 9000ft. System also displays the amount of headwind and tailwind. 

## Possible development direction/tasks: 
- Start with a cli app using cobra and make it take the inputs and do the calculation. 
- Write some tests and use GitHub actions to run them with each commit
- Add a `serve` option to the cli which allows input/outputs via htmx on a web page
- Put it in a Docker image and run it somewhere
- Use fyne to build a mobile app 
- Package for Android
- Enhance so that it can fetch the data prior to the flight, and rely on that data store when internet is not available. 


## Data sources
This stations file contains the gps coordinates of airports: 
https://aviationweather.gov/data/cache/stations.cache.json.gz

```
jq '.[] | select(.icaoId == "KAXN")' stations.cache.json 
{
  "icaoId": "KAXN",
  "iataId": "AXN",
  "faaId": "AXN",
  "wmoId": "-",
  "lat": 45.868,
  "lon": -95.394,
  "elev": 433,
  "site": "Alexandria/Chandler Fld",
  "state": "MN",
  "country": "US",
  "priority": 3
}
```