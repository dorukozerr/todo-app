# CLI Todo App

To run this app in your machine follow theese instructions

Instructions should work for mac/linux (didn't tested for linux but I belive it should work)

-   Make sure u have no `~/.todo` folder and also no `/usr/local/bin/td`, `/usr/local/bin/tda`, `/usr/local/bin/tdd`, and `/usr/local/bin/tdr` files in given paths. You probably wont have them but mentioned it just in case :)))) .
-   Copy and run `git clone git@github.com:dorukozerr/todo-app.git ~/.todo && cd ~/.todo && go build -o td main.go && chmod +x td`
-   Copy and run `sudo mv ~/.todo/td /usr/local/bin && cd /usr/local/bin && sudo ln -s td tda && sudo ln -s td tdd && sudo ln -s td tdr`
-   Quit and reopen your terminal and now you should be able to use `td`, `td -a`, `tda`, `tdd <index>`, `tdr <index>` commands.
-   Todos will be saved in `~/todo/todos.json` file.
-   If you want you can add td to your `.zshrc` or `.bashrc` file to display active todos when you open your terminal.
-   If commands are not working try adding `export PATH=/usr/local/bin:$PATH"` to your `.zshrc` or `.bashrc` file.

Commands

-   list active todos -> td
-   list all todos    -> td -a
-   add todo          -> tda
-   check todo N      -> tdd N
-   remove todo N     -> tdr N

I made this project to getting familliar with GO, this is a minimal replica of devtodo app. I used to use it on my linux machine few years ago and couldnt find theexactly same functionality with same commands for my mac so I made this.
