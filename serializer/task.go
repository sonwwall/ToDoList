package serializer

import "ToDoList/model"

type Task struct {
	ID        uint   `json:"id" form:"id" example:"1"`                 //任务ID
	Title     string `json:"title" form:"title" example:"第一个任务"`       //任务标题
	Content   string `json:"content" form:"content" example:"第一个任务内容"` //任务内容
	View      uint64 `json:"view" form:"view" example:"32"`            //任务查看次数
	Status    int    `json:"status" form:"status" example:"0"`         //任务状态
	CreateAt  int64  `json:"create_at" form:"create_at" `              //任务创建时间
	StartTime int64  `json:"start_time" form:"start_time" `            //任务开始时间
	EndTime   int64  `json:"end_time" form:"end_time" `                //任务结束时间
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
		//View:      task.View,
		Status: task.Status,
		//CreateAt:  task.CreateAt,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
