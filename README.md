# goansi

Colour and style for the terminal.

## Usage
Run the example:
```
go run examples/example.go
```
Use the library:
```
s := goansi.Anstring("I am black on green.").Black().BgGreen().Underline()
fmt.Println(s)

Or

style := goansi.NewStyle().Gray().BgBlue().Bold().BlinkSlow()
fmt.Println(style.Render("I am white on blue."))
```

## License
BSD 3-Clause: [LICENSE.txt](LICENSE.txt)

[<img alt="LICENSE" src="http://img.shields.io/pypi/l/Django.svg?style=flat-square"/>](LICENSE.txt)
