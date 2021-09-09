# ceelei

# What is Ceelei?
Ceelei is a little CLI tool that wants to help you to use the CLI more efficient.
Working with the CLI can be hard or time consuming at first but when you get used to it its much faster than clicking around. How often you have to Google a CLI command because you forgot it? Then there is the possibility to use a note taking system to write down your favourite CLI commands. This is pretty inconvenient if you ask me as you have to switch to another program while it would be nice to stay inside the CLI while you are working in it.
Here comes Ceelei!

# A CLI tool for the CLI
Adding and listing CLI commands inside the CLI is easy and pleasant for the user.

## ceelei
![ceelei base command](https://media-files-dtf.s3.eu-central-1.amazonaws.com/github_media/ceelei/base_command.png)

After you installed ceelei you can just type ceelei into your CLI and get an overview about what ceelei is and how to use it.

## ceelei add
![ceelei base command](https://media-files-dtf.s3.eu-central-1.amazonaws.com/github_media/ceelei/add_command.png)

With ceelei add "command" "description" you can add a command with a description to your list.

## ceelei help
![ceelei base command](https://media-files-dtf.s3.eu-central-1.amazonaws.com/github_media/ceelei/add_help.png)

Of course every command has a -help command flag when you need help.

## ceelei list
![ceelei base command](https://media-files-dtf.s3.eu-central-1.amazonaws.com/github_media/ceelei/list_command.png)

With ceelei list you will see all the commands with the description you gave them in a list. Each command also has an unique ID.

## ceelei remove
![ceelei base command](https://media-files-dtf.s3.eu-central-1.amazonaws.com/github_media/ceelei/remove_command.png)

With ceelei remove "id" you can easily remove a command from your list. Just use the ID of the command what you can show with ceelei list.

# Saved locally
Your CLI list will be saved locally inside your home directory inside a file called ceelei.db.
If you want to start fresh you can just delete this file.
After adding a new command with ceelei this file will be generated again.

# Future ideas
Some ideas that i might want to add in the future:
 - edit functionality
 - grouping the commands
 - ...

If you have suggestions or ideas please create a PR.

# Credits
In this project i used [bbolt](https://github.com/etcd-io/bbolt) as a database solution and [cobra](https://cobra.dev/) as the base for the CLI.