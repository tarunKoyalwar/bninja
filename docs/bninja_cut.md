## bninja cut

Similar to GNU cut but better

### Synopsis

Similar to GNU cut but with clipboard access,multicharacter delimeter and can cut multiple whitespaces

```
bninja cut [flags]
```

### Examples

```

// cut with delimeter as go.mod and use data from clipboard
bninja cut -d "go.mod" -f 2 -i

```

### Options

```
  -d, --delim string    Delimeter to Use(Supports Multiple Chars)
  -f, --fields string   fields/colums but it can be 1-3 or 1,5,6 or 1-3,8 (Inclusive)(starts from 1)
  -h, --help            help for cut
      --skip            Skip Lines not containing delimeter
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

