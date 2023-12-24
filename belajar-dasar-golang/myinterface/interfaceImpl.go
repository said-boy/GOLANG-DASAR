package myinterface

type Hero struct {
	Name string
}

func (h *Hero) PrintName() string {
	return h.Name
}

func (h *Hero) Run(r int) int {
	run := 0
	for i := 0; i < r; i++ {
		run += 10
	}
	return run
}

type Enemy struct {
	Name string
}

func (h *Enemy) Run(r int) int {
	run := 0
	for i := 0; i < r; i++ {
		run += 5
	}
	return run
}
