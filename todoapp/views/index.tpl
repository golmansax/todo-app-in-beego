<!DOCTYPE html>

<html>
<head>
  <title>Todo App</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1>My Todo List</h1>
  </header>
  <div>
    <a href="/new-todo">New Todo</a>
    {{range $index, $todo := .todos}}
      <div>
        {{$todo.Name}}
      </div>
    {{end}}
  </div>
</body>
</html>
