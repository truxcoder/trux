//go:build windows
// +build windows

package run

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/truxcoder/trux/config"
	"github.com/truxcoder/trux/internal/pkg/bubbletea"
	"github.com/truxcoder/trux/internal/pkg/helper"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var quit = make(chan os.Signal, 1)

type Run struct {
}

var excludeDir string
var includeExt string

func init() {
	CmdRun.Flags().StringVarP(&excludeDir, "excludeDir", "", excludeDir, `eg: trux run --excludeDir="tmp,vendor,.git,.idea"`)
	CmdRun.Flags().StringVarP(&includeExt, "includeExt", "", includeExt, `eg: trux run --includeExt="go,tpl,tmpl,html,yaml,yml,toml,ini,json"`)
	if excludeDir == "" {
		excludeDir = config.RunExcludeDir
	}
	if includeExt == "" {
		includeExt = config.RunIncludeExt
	}
}

var CmdRun = &cobra.Command{
	Use:     "run",
	Short:   "trux run [main.go path]",
	Long:    "trux run [main.go path]",
	Example: "trux run cmd/server",
	Run: func(cmd *cobra.Command, args []string) {
		cmdArgs, programArgs := helper.SplitArgs(cmd, args)
		var dir string
		if len(cmdArgs) > 0 {
			dir = cmdArgs[0]
		}
		base, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
			return
		}
		if dir == "" {
			cmdPath, err := helper.FindMain(base, excludeDir)

			if err != nil {
				fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
				return
			}
			switch len(cmdPath) {
			case 0:
				fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", "The cmd directory cannot be found in the current directory")
				return
			case 1:
				for _, v := range cmdPath {
					dir = v
				}
			default:
				var cmdPaths []string

				for k := range cmdPath {
					cmdPaths = append(cmdPaths, k)
				}
				sort.Strings(cmdPaths)
				dir = bubbletea.BuildListTea("Which directory do you want to run?", cmdPaths)
				dir = cmdPath[dir]
			}
		}
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		fmt.Printf("\033[35mTrux run %s.\033[0m\n", dir)
		fmt.Printf("\033[35mWatch excludeDir %s\033[0m\n", excludeDir)
		fmt.Printf("\033[35mWatch includeExt %s\033[0m\n", includeExt)
		watch(dir, programArgs)

	},
}

func watch(dir string, programArgs []string) {

	// Listening file path
	watchPath := "./"

	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	excludeDirArr := strings.Split(excludeDir, ",")
	includeExtArr := strings.Split(includeExt, ",")
	includeExtMap := make(map[string]struct{})
	for _, s := range includeExtArr {
		includeExtMap[s] = struct{}{}
	}
	// Add files to watcher
	err = filepath.Walk(watchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, s := range excludeDirArr {
			if s == "" {
				continue
			}
			if strings.HasPrefix(path, s) {
				return nil
			}
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if _, ok := includeExtMap[strings.TrimPrefix(ext, ".")]; ok {
				err = watcher.Add(path)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cmd := start(dir, programArgs)

	// Loop listening file modification
	for {
		select {
		case <-quit:
			err = killProcess(cmd)

			if err != nil {
				fmt.Printf("\033[31mserver exiting...\033[0m\n")
				return
			}
			fmt.Printf("\033[31mserver exiting...\033[0m\n")
			os.Exit(0)

		case event := <-watcher.Events:
			// The file has been modified or created
			if event.Op&fsnotify.Create == fsnotify.Create ||
				event.Op&fsnotify.Write == fsnotify.Write ||
				event.Op&fsnotify.Remove == fsnotify.Remove {
				fmt.Printf("\033[36mfile modified: %s\033[0m\n", event.Name)
				killProcess(cmd)

				cmd = start(dir, programArgs)
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}

func killProcess(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}
	// 获取进程ID
	pid := cmd.Process.Pid
	// 构造taskkill命令
	taskkill := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
	err := taskkill.Run()
	if err != nil {
		return err
	}
	return nil
}
func start(dir string, programArgs []string) *exec.Cmd {
	cmd := exec.Command("go", append([]string{"run", dir}, programArgs...)...)
	// Set a new process group to kill all child processes when the program exits

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("\033[33;1mcmd run failed\u001B[0m")
	}
	time.Sleep(time.Second)
	fmt.Printf("\033[32;1mrunning...\033[0m\n")
	return cmd
}
