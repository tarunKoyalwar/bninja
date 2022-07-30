## bninja uniq

Print Unique lines

### Synopsis

Print Unique Lines (uses golang map to filter duplicates)

```
bninja uniq [flags]
```

### Examples

```

// Print Unique Lines Only
cat subdomains.txt | bninja uniq

```

### Options

```
  -h, --help             help for uniq
      --insensitive      ignore case/ insensitive
      --show-frequency   Show frequency of each item(i.e no of times it appeared
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

