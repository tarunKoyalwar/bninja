## bninja replace

Match & Replace Data

### Synopsis

Match & Replace Data from files In Regex Mode Go(RE2) Engine is used

```
bninja replace [flags]
```

### Examples

```

// replace first 5 occurences of "programmer" to "dev"
bninja replace -m "programmer" -r "dev" -c 5

```

### Options

```
  -c, --count int        No of Occurences to replace(default -1 i.e all) (default -1)
  -h, --help             help for replace
  -m, --match string     String to Match
      --posix            Use POSIX for regex
      --regex            Regex Mode
  -r, --replace string   replace matched character with string (can be empty)
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

