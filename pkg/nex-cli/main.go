package nexcli

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/mitranim/gg"
	"github.com/rjeczalik/notify"
)

type Cmd struct {
	sync.Mutex
	Cmd *exec.Cmd
}

func (cm *Cmd) Deinit() {
	defer gg.Lock(cm).Unlock()
	cm.DeinitUnSync()
}

func (cm *Cmd) DeinitUnSync() {
	cm.BroadcastUnSync(syscall.SIGTERM)
	cm.Cmd = nil
}

func (cm *Cmd) Restart(fileName string) {
	defer gg.Lock(cm).Unlock()
	cm.DeinitUnSync()

	cmd := exec.Command("go", "run", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Println(`unable to start subcommand:`, err)
		return
	}

	cm.Cmd = cmd
	go cm.CmdWait(cmd)
}

func (cm *Cmd) BroadcastUnSync(sig syscall.Signal) {
	proc := cm.ProcUnSync()
	if proc != nil {
		gg.Nop1(syscall.Kill(-proc.Pid, sig))
	}
}

func (cm *Cmd) ProcUnSync() *os.Process {
	cmd := cm.Cmd
	if cmd != nil {
		return cmd.Process
	}
	return nil
}

func (cm *Cmd) CmdWait(cmd *exec.Cmd) {
	err := cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

type WatchNotify struct {
	Cmd    Cmd
	Done   gg.Chan[struct{}]
	Events gg.Chan[notify.EventInfo]
}

func (wn *WatchNotify) Init() {
	wn.Done.Init()
	wn.Events.InitCap(1)

	path := filepath.Join(".", `...`)
	gg.Try(notify.Watch(path, wn.Events, notify.All))
}

func (wn *WatchNotify) Deinit() {
	wn.Done.SendZeroOpt()
	if wn.Events != nil {
		notify.Stop(wn.Events)
	}
}

func (wn *WatchNotify) Run(fileName string) {
	for {
		select {
		case <-wn.Done:
			return
		case event := <-wn.Events:
			if filepath.Ext(event.Path()) == ".go" && event.Event() == notify.Write {
				wn.Cmd.Restart(fileName)
			}
		}
	}
}
