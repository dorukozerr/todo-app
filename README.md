# CLI Todo App

-   You need Go installed on your system.
-   There should be no ~/.todo directory.
-   `/usr/local/bin` folder should exist and there should be no `td`, `tda`, `tdd`, and `tdr` files inside that folder.

#### Install command for macos

`git clone git@github.com:dorukozerr/todo-app.git ~/.todo && cd ~/.todo && go build -o td main.go && chmod +x td && sudo mv ~/.todo/td /usr/local/bin && cd /usr/local/bin && sudo ln -s td tda && sudo ln -s td tdd && sudo ln -s td tdr && cd ~`

### Notes

-   Todos will be saved in `~/.todo/todos.json` file.
-   If you want you can add `td` to your `.zshrc` or `.bashrc` file to display active todos when you open your terminal.

| Command | Description        |
| :------ | :----------------- |
| `td`    | list active todos  |
| `td -a` | list all todos     |
| `tda`   | Add a todo         |
| `tdd N` | check todo N       |
| `tdr N` | remove todo N      |
