package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) AddToList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) deleteFromList(index int) {
	t.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	done        bool
}

func (t *task) markDone() {
	t.done = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	t := &task{
		name:        "Complete my Go course",
		description: "Complete my Go course in Platzi this week",
	}
	t2 := &task{
		name:        "Complete my Pyhon course",
		description: "Complete my Pyhon course in Platzi this week",
	}
	t3 := &task{
		name:        "Complete my Node course",
		description: "Complete my Node course in Platzi this week",
	}

	list := &taskList{
		tasks: []*task{
			t, t2,
		},
	}

	fmt.Println(list.tasks[0])
	list.AddToList(t3)
	fmt.Println(list.tasks[2])

}
