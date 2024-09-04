# winds-aloft
Easily calculate the best altitude to fly

## Problem Statement
Pilots can often choose the altitude at which they fly. Aviation weather sources provide the wind direction and speed at various altitudes. Pilots would like to quickly calculate the best altitude to fly in order to minimize headwind or maximize tailwind. This saves time and fuel. 

Possible data sources: 
https://aviationweather.gov/data/api/
https://aviationweather.gov/data/windtemp/?region=chi&fcst=06&level=low

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



