func main() {
	name := "Todd McLeod
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello world!</title>
	</head>
	<h1> +
					name +
					`</h1>
	</body>
	</html>
		`>
		nf, err := os.Create("index.html")
		if err != nil {
			log.Fatal("error creating file" err)
		}
		defer nf.Close()

		io.Copy(nf, strings.NewReader(str))
	
}