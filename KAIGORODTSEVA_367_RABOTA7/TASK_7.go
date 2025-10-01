package main

import (
	"fmt"
)

type Task struct {
	Name   string
	Status string
	ID     int
}

type TaskManager struct {
	List map[int]Task
}

func (tm *TaskManager) Add(task Task) {

	if tm.List == nil {
		tm.List = make(map[int]Task)
	}

	tm.List[task.ID] = task
	fmt.Println("Задача ", task.Name, " добавлена")
}

func (tm *TaskManager) Delete(task Task) {

	if _, ok := tm.List[task.ID]; ok {
		delete(tm.List, task.ID)
		fmt.Println("Задача ", task.Name, "  удалена")
	} else {
		fmt.Println("Задача ", task.Name, " не найдена для удаления")
	}
}

func (tm *TaskManager) ChangeStatus(taskID int, newStatus string) {

	if task, ok := tm.List[taskID]; ok {

		task.Status = newStatus

		tm.List[taskID] = task
		fmt.Println("Статус задачи ", task.Name, "  изменен на ", newStatus)
	} else {
		fmt.Println("Задача ", task.Name, " не найдена ")
	}
}

func main() {

	var tm TaskManager

	tm.List = make(map[int]Task)

	t1 := Task{
		Name:   "TASK 1",
		Status: "not done",
		ID:     1,
	}
	t2 := Task{
		Name:   "TASK 2",
		Status: "not done",
		ID:     2,
	}

	tm.Add(t1)
	tm.Add(t2)
	tm.ChangeStatus(t1.ID, "done")

	tm.Delete(t2)

	t3 := Task{Name: "TASK 3", ID: 3}
	tm.Delete(t3)
	tm.ChangeStatus(3, "done")

	fmt.Println("Текущее состояние TaskManager:")
	if len(tm.List) == 0 {
		fmt.Println("Нет активных задач.")
	} else {
		for id, task := range tm.List {
			fmt.Printf("ID: %d, Name: %s, Status: %s\n", id, task.Name, task.Status)
		}
	}

}
