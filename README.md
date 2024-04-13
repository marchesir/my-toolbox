# Indroduction
One Billion Row Challenge (1BRC) is a fun coding exercise how far a program can be psuhed to aggregate  one billion rows from a text file.

# Problem
Input: Given a [text file](https://github.com/gunnarmorling/1brc/blob/main/data/weather_stations.csv) containing temperature values for a range of weather stations.  Each row is one measurement in the format:
*<string: station name;<float:measurement>*.

Output: For each unique station find the min/max/average temperature recorded abd emit the final result on STDOUT in the station name's alphabetical 
order with the format:
*{<station name}:<min>/max/average}*.

Assumptions:
* The temperature value is with the range [-99.9, 99.9];
* The temperature value has only one fractional digit;
* The length odf the staion name is within [1,100];
* Value rounding must be done using the semantcis of IEEE 754 rounding-direction "roundTowardPositive";

# References
This is loosly based off [1BRC Java](https://github.com/gunnarmorling/1brc). 
