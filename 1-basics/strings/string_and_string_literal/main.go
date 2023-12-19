package main

import "fmt"

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	//Go just take this string as it is, that's why you don't need \n\t like above
	s = `
<html>
	<body>"Hello"</body>
</hmtl>`
	fmt.Println(s)


	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)

}