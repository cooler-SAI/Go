package main

type Plug interface {
	Working() string
}

type HomeDevice struct {
	Name  string
	Power int
}

func (h HomeDevice) Working() string {
	return h.Name
}

func SetUp(plug Plug) string {
	return plug.Working()
}

func main() {

	t := HomeDevice{
		Name:  "Toaster",
		Power: 10,
	}

}
