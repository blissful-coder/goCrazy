package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type task struct {
	text string
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Eisenhower Matrix")

	quadrants := make([]*fyne.Container, 4)
	taskLists := make([]*widget.List, 4)
	taskData := make([][]task, 4)

	for i := 0; i < 4; i++ {
		taskData[i] = []task{}
		taskLists[i] = widget.NewList(
			func() int {
				return len(taskData[i])
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(j widget.ListItemID, o fyne.CanvasObject) { // j is the list item index
				o.(*widget.Label).SetText(taskData[i][j].text) // Use i and j correctly!
			},
		)

		addTaskButton := widget.NewButton("Add Task", func() {
			entry := widget.NewEntry()
			dialog.ShowForm("New Task", "Add", "Cancel", []*widget.FormItem{
				widget.NewFormItem("Task", container.New(layout.NewMaxLayout(), entry)),
			}, func(b bool) {
				if b {
					newTask := task{text: entry.Text}
					taskData[i] = append(taskData[i], newTask)
					taskLists[i].Refresh()
				}
			}, myWindow)
		})

		taskLists[i].OnSelected = func(id widget.ListItemID) {
			selectedTask := taskData[i][id]
			editEntry := widget.NewEntry()
			editEntry.SetText(selectedTask.text)

			dialog.ShowConfirm("Edit or Delete", "Do you want to edit or delete this task?", func(deleteConfirmed bool) {
				if deleteConfirmed {
					taskData[i] = append(taskData[i][:id], taskData[i][id+1:]...)
					taskLists[i].Refresh()
				} else {
					dialog.ShowForm("Edit Task", "Save", "Cancel", []*widget.FormItem{
						widget.NewFormItem("Task", container.New(layout.NewMaxLayout(), editEntry)),
					}, func(saveConfirmed bool) {
						if saveConfirmed {
							taskData[i][id].text = editEntry.Text
							taskLists[i].Refresh()
						}
					}, myWindow)
				}
			}, myWindow)
		}

		quadrants[i] = container.NewVBox(
			addTaskButton,
			taskLists[i],
		)
	}

	grid := container.New(layout.NewGridLayout(2),
		container.NewVBox(
			canvas.NewText("Do First", color.White),
			quadrants[0]), // Important & Urgent
		container.NewVBox(
			canvas.NewText("Schedule", color.White),
			quadrants[1]), // Important & Not Urgent
		container.NewVBox(
			canvas.NewText("Delegate", color.White),
			quadrants[2]), // Not Important & Urgent
		container.NewVBox(
			canvas.NewText("Don't Do", color.White),
			quadrants[3]), // Not Important & Not Urgent
	)

	myWindow.SetContent(
		container.New(layout.NewMaxLayout(),
			canvas.NewRectangle(theme.BackgroundColor()), // Background
			grid,
		),
	)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}
