## bninja html

HTML Encoder/Decoder

### Synopsis

Encode and Decode HTML

```
bninja html [flags]
```

### Examples

```

// HTML Encode data
echo "Fran & Freddie's Diner" <tasty@example.com>" | bninja html

// HTML Decode data
echo "Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;" | bninja html -d

```

### Options

```
  -d, --decode   HTML Decode
  -h, --help     help for html
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

