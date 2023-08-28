# go_htmx_todo_app
A simple to-do list web app using:
- Go, Go Templates, and the standard net/http package
- htmx
- Tailwind CSS
- SQLite
It is barebones and minimally styled, as it was used as a project to explore the basics of different technologies. 
## Getting Starting
There are a few steps required to get the project working.
### Set up SQLite DB
Ensure that [SQLite](https://www.sqlite.org/quickstart.html) is installed on your machine. Then, initialize the SQLite DB by creating it in the terminal:
```bash
sqlite3 todos.db
```
Then in the sqlite shell prompt, create the table used in the app:
```sql
CREATE TABLE `todos` (
        `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
        `title` VARCHAR(64) NULL,
        `detail` VARCHAR(64) NULL,
        `created` DATE NULL
    );
```
### Initialize Tailwind CSS
This project uses the standalone [TailwindCLI](https://tailwindcss.com/blog/standalone-cli). The base input.css file is included in this repo. After installing the standalone CLI, generate the css used in the file by running the following command in the shell from the project root:
```shell
./tailwindcss -i ./static/input.css -o ./static/output.css
```

## Next Steps
- Improve error handling of Go code
- Better organize html with Go Templates
- Add ability to edit To Do items in-line
- Add confirmation modal before deleting a record
- Update code stlying
- Add proper header and footer to base html page
