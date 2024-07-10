package domains

func TodoToTodoResp(t Todo) TodoResp {
	return TodoResp{
		Id:        t.Id,
		Title:     t.Title,
		Content:   t.Content,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func RespTodos(t []Todo) []TodoResp {
	todos := []TodoResp{}
	for _, v := range t {
		todo := TodoToTodoResp(v)
		todos = append(todos, todo)
	}
	return todos

}
