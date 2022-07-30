## bninja hcut

Cut Data horizantally similar to GNU cut but horizantal

### Synopsis

Cut Data horizantally similar to GNU cut but horizantal

```
bninja hcut [flags]
```

### Examples

```

// cut data horizantally
cat somefile.txt | bninja hcut -d "delimeter" -f 1,7-8

```

### Options

```
  -d, --delim string    Delimeter to Use(Supports Multiple Chars)
  -f, --fields string   fields/colums but it can be 1-3 or 1,5,6 or 1-3,8 (Inclusive)(starts from 1)
  -h, --help            help for hcut
```

### Options inherited from parent commands

```
  -i, --clipin        Import data from clipboard
  -o, --clipout       Export data to clipboard
      --file string   Import data from File
      --out string    Export data to File
      --stdout        Print to Stdout even if data if data is exported to clipboard or file
```

### SEE ALSO

* [bninja](bninja.md)	 - A Handy Toolkit for bash ninja's

