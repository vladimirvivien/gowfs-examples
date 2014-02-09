# Gowfs-Examples
A collection of sample code showing usage of the  [gowfs HDFS client](https://github.com/vladimirvivien/gowfs) .

### hdfstat
This a basic example that shows how to get the status of a given path.

##### Usage
```
hdfstat -path=/user/hadoop -user=hadoop
```

### hdfls
A more functional example that prints the directory listing of a given path.

##### Usage
```
hdfls -path=/user/hadoop -user=hadoop
drwxr-xr-x    - hadoop	supergroup	          0  2014-01-29 15:09:02 data-directory
drwxr-xr-x    - hadoop	supergroup	          0  2014-01-29 15:28:59 output-directory
```