# ceelei

# What is Ceelei?
Ceelei is a little CLI tool that wants to help you to use the CLI more efficient.
Working with the CLI can be hard or time consuming at first but when you get used to it its much faster than clicking around. How often you have to Google a CLI command because you forgot it? Then there is the possibility to use a note taking system to write down your favourite CLI commands. This is pretty inconvenient if you ask me as you have to switch to another program while it would be nice to stay inside the CLI while you are working in it.
Here comes Ceelei!

# A CLI tool for the CLI

# Saved locally
Your CLI list will be saved locally inside your home directory inside a file called ceelei.db.
If you want to start fresh you can just delete this file.
After adding a new command with ceelei this file will be generated again.

# Future ideas
Some ideas that i might want to add in the future:
 - edit functionality
 - grouping the commands
 
If you have suggestions or ideas please create a PR.

# Credits
In this project i used [bbolt](https://github.com/etcd-io/bbolt) as a database solution and [cobra](https://cobra.dev/) as the base for the CLI.