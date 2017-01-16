# gohistgrep
Command line app for quick search through the command history of shell

Sometime we need to see the command that we ran for a very long time. In the command shell, `bash` is used for this file `.bash_history`. 
The history in this file to be kept for a very long time. So I had the idea to store the history of my shell commands forever! 

Then I found somewhere on the internet this script:

```
HISTFILESIZE=2000
mkdir -p $(dirname  $HISTFILE)
export HISTFILE="$HOME/.history/$(date -u +%Y/%m/%d.%H.%M.%S)_${HOSTNAME_SHORT}_$$"
histgrep()
{
  grep -r "$@" ~/.history
  history | grep "$@"
}
```

Just add these lines to your `.bashrc` file. And that's all :)

Now you will have to store the history of your commands carefully. 
So you have a new feature in `Bash`. 
You can write `histgrep` continue to specify a part of your command and `Bash` will run on all history files and show matches!

Much time has passed and I have accumulated my command history in `Bash` for 3 years!
Search vehicles `grep` was very long... 

So I decided to write this small tool in Golang. 
This app reads all files of your history and stores data from them in sqlite database. 
Next command works only with this database, so search through the command history is very fast!

Additionally, this app is able to count popular commands in history, which makes search more relevant :) 
Your little Google right in your console :)
