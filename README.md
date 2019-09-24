# What is gohunting?

gohunting is a golang process hunting tool which you can use to retrieve useful informations about running processes. Also, you can pass a file as a parameter an run it, the result is the same. The tool was result of some nights reading about golang and it's pretty simple but still a nice tool to use as your first step in a Incident Response Analysis.

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

1 - I'm trying to learn Go

2 - It's cool

## Feedbacks

Have any idea? Any feedback? Feel free to contact me, I really enjoy it!