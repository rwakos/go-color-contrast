# go-color-contrast

Color contrast following WCAG 2 and WAI

## Makefile

Use the make file to create a build:

```
make build
```

Then execute the program:

```
bin/main -f="#FFFFFF" -b="#000000"
```

Parameters:

`f : string Foreground color`

`b : string Background color`

Results:

- Decimal | Contrast between foreground and background
- Boolean | AA for smaller text than 18 points (ratio as 4.5 : 1)
- Boolean | AA for bigger text than 17 points (ratio as 3 : 1)
- Boolean | AAA for smaller text than 18 points (ratio as 7 : 1)
