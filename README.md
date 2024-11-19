# CLI Todo App

To run this app on your machine follow theese instructions.

Instructions should work for mac/linux (didn't tested for linux but I belive it should work).

-   Make sure u have no `~/.todo` folder and also no `/usr/local/bin/td`, `/usr/local/bin/tda`, `/usr/local/bin/tdd`, and `/usr/local/bin/tdr` files in given paths. You probably wont have them but mentioned it just in case.
-   Copy and run `git clone git@github.com:dorukozerr/todo-app.git ~/.todo && cd ~/.todo && go build -o td main.go && chmod +x td`.
-   Copy and run `sudo mv ~/.todo/td /usr/local/bin && cd /usr/local/bin && sudo ln -s td tda && sudo ln -s td tdd && sudo ln -s td tdr`.
-   Quit and reopen your terminal and now you should be able to use `td`, `td -a`, `tda`, `tdd <index>`, `tdr <index>` commands.
-   Todos will be saved in `~/todo/todos.json` file.
-   If you want you can add `td` to your `.zshrc` or `.bashrc` file to display active todos when you open your terminal.
-   If commands are not working try adding `export PATH=/usr/local/bin:$PATH"` to your `.zshrc` or `.bashrc` file.

| Command          | Description        |
| :--------------- | :----------------- |
| <kbd>td</kbd>    | list active todos  |
| <kbd>td -a</kbd> | list all todos     |
| <kbd>tda</kbd>   | Add a todo         |
| <kbd>tdd N</kbd> | check todo N       |
| <kbd>tdr N</kbd> | remove todo N      |
