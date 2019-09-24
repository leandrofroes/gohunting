package main

import (
  "fmt"
  "os"
  "os/exec"
  "flag"
  "github.com/shirou/gopsutil/process"
  "github.com/shirou/gopsutil/net"
  "github.com/fatih/color"
)

func banner(){

  fmt.Println(`
  -------------------------------------------------------
 |                _                 _   _                |
 |     __ _  ___ | |__  _   _ _ __ | |_(_)_ __   __ _    |
 |    / _  |/ _ \| '_ \| | | | '_ \| __| | '_ \ / _  |   |
 |   | (_| | (_) | | | | |_| | | | | |_| | | | | (_| |   |
 |    \__, |\___/|_| |_|\__,_|_| |_|\__|_|_| |_|\__, |   | 
 |    |___/                                     |___/    |
 |                                                       |
 |   gohunting: a golang process hunting tool            |
 |   version: 0.1                                        |
  -------------------------------------------------------
   `)
 }

type Report struct{
  Name string
  Background bool
  ParentPID int32
  ChildrenPID []int32
  Status string
  MemoryPercent float32
  StartedBy string
  WorkingDir string
  BinaryPath string
  OpenFiles []process.OpenFilesStat
  Connections []net.ConnectionStat
}

func check(err error){
  if(err != nil){
    fmt.Println(err)
    os.Exit(1)
  }
}

func print_report(r Report){

  green := color.New(color.FgGreen).SprintFunc()

  color.Cyan("[*] Process Information:\n------------------------\n\n")
  fmt.Println("[+] Process Name:", green(r.Name))
  fmt.Println("[+] Background:", green(r.Background))
  fmt.Println("[+] Parent PID:", green(r.ParentPID))
  fmt.Println("[+] Children PID:", green(r.ChildrenPID))
  fmt.Println("[+] Status:", green(r.Status))
  fmt.Println("[+] Memory Usage:", green(r.MemoryPercent))
  fmt.Println("[+] Started by:", green(r.StartedBy))
  fmt.Println("[+] Working Directory:", green(r.WorkingDir))
  fmt.Println("[+] Binary Path:", green(r.BinaryPath))

  fmt.Println("[+] Open Files:")

  for _, value := range r.OpenFiles{
    fmt.Println("\t-", green(value.Path))
  }

  fmt.Println("[+] Connections:")

  for _,conn := range r.Connections{
    fmt.Println("\tFileDescriptor:", green(conn.Fd))
    fmt.Println("\tFamily:", green(conn.Family))
    fmt.Println("\tType:", green(conn.Type))
    fmt.Printf("\tLocal Address: %s:%d\n", conn.Laddr.IP, conn.Laddr.Port)
    fmt.Printf("\tRemote Address: %s:%d\n", conn.Raddr.IP, conn.Raddr.Port)
    fmt.Printf("\tStatus: %s\n\n", green(conn.Status))
  }

}

func parse(proc *process.Process) (r Report){

  r = Report{}
  r.Name, _ = proc.Name()
  r.Background, _ = proc.Background()
  r.Status, _ = proc.Status()
  r.MemoryPercent, _ = proc.MemoryPercent()
  r.StartedBy, _ = proc.Username()
  r.WorkingDir, _ = proc.Cwd()
  r.BinaryPath, _ = proc.Exe()

  parent, err := proc.Parent()

  if(err != nil){
    r.ParentPID = 0
  }else{
    r.ParentPID = parent.Pid
  }

  var children_array []int32
  children, _ := proc.Children()

  for _,value := range children{
    children_array = append(children_array, value.Pid)
  }

  r.ChildrenPID = children_array

  path_str, _ := proc.OpenFiles()
  r.OpenFiles = path_str
  r.Connections, _ = proc.Connections()

  return r

}

func run_file (file string){
    cmd := exec.Command(file)
    check(cmd.Start())
    file_data := cmd.Process
    pid := file_data.Pid
    proc, err := process.NewProcess(int32(pid))
    check(err)
    report_data := parse(proc)
    print_report(report_data)
}

func run_proc (pid int){
    proc, err := process.NewProcess(int32(pid))
    check(err)
    report_data := parse(proc)
    print_report(report_data)
}

func main(){

  usage := `Usage: ./gohunting [-p PID] [-f FILE]
              
  -p PID
    Specificy the Process ID
  -f FILE
    Specify the file to run
`

  flag.Usage = func(){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  pid := flag.Int("p", 0,  "")
  file := flag.String("f", "", "")

  flag.Parse()

  if(flag.NFlag() != 1){
    fmt.Println("Cannot use 2 arguments at the same time.")
    flag.Usage()
    os.Exit(1)
  }

  banner()

  switch{
    case *file != "":
      run_file(*file)
    case *pid != 0:
      run_proc(*pid)
  }

}