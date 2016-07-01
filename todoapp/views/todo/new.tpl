<!DOCTYPE html>

<html>
<head>
  <title>Todo App</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <h1>New Todo</h1>
  </header>
  <div>
    <form method="post" action="/create-todo">
      <div>
        <label for="name">Name</label>
        <input name="name" />
      </div>
      <br />
      <input type="submit" value="Create Todo" />
    </form>
  </div>
</body>
</html>
