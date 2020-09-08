# Helper scripts to generate tables for scale testing

```
usage: gen_tables.py [-h] -numOfTables NUMOFTABLES -numOfCols NUMOFCOLS
                     -dataType DATATYPE [-numOfPartitions NUMOFPARTITIONS]
                     [-outputFile OUTPUTFILE]

Generate Tables Schema

optional arguments:
  -h, --help            show this help message and exit
  -numOfTables NUMOFTABLES
  -numOfCols NUMOFCOLS
  -dataType DATATYPE
  -numOfPartitions NUMOFPARTITIONS
  -outputFile OUTPUTFILE
```

```
Example: ./gen_tables.py -h 10 -numofcols 100 -datatype text -numofPartitions 100 -outputFile /tmp/output.txt
```
The above command will generate CREATE and INSERT statement to create 10 root
tables with 100 child partitions each, and 100 columns of datatype text and
will place the output in the file /tmp/output.txt. The generated SQLs can then
be used to load database.

Note: The script is far from complete, but gives you starting.
