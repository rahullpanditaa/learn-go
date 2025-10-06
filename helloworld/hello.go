package helloworld

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return "Hello, " + name
}
