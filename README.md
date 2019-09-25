# What is gohunting?

gohunting is a golang process hunting tool which you can use to retrieve useful informations from running process. Also, you can pass a file as a parameter an run it, the result is the same. 

For now the tool itself is pretty simple, but still a nice tool to use as your first step in a Forensic Analysis.

![](https://github.com/leandrofroes/gohunting/blob/master/pictures/gohunting_output.png)

## Installation

``` 
git clone https://github.com/leandrofroes/gohunting.git
cd gohunting
make
```

## Usage

```
./gohunting [-p PID] [-f FILE]
              
  -p PID
    Specificy the Process ID
  -f FILE
    Specify the file to run
```

## Why Go?

- Because I'm trying to learn Go
- Because the language is cool

## Future Features

- Process memory dump
- JSON output

## Feedbacks

Have any idea? Any feedback? Feel free to contact me, I really enjoy it!
