package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/PavIek/cube/manager"
// 	"github.com/PavIek/cube/task"
// 	"github.com/PavIek/cube/worker"
// 	"github.com/golang-collections/collections/queue"
// 	"github.com/google/uuid"
// )

// func main() {
// 	host := os.Getenv("CUBE_HOST")
// 	port, _ := strconv.Atoi(os.Getenv("CUBE_PORT"))

// 	fmt.Println("Starting Cube worker")

// 	w := worker.Worker{
// 		Queue: *queue.New(),
// 		Db:    make(map[uuid.UUID]*task.Task),
// 	}
// 	api := worker.Api{Address: host, Port: port, Worker: &w}

// 	go w.RunTasks()
// 	go w.CollectStats()
// 	go api.Start()

// 	workers := []string{fmt.Sprintf("%s:%d", host, port)}
// 	m := manager.New(workers)

// 	for i := 0; i < 3; i++ {
// 		t := task.Task{
// 			ID:    uuid.New(),
// 			Name:  fmt.Sprintf("test-container-%d", i),
// 			State: task.Scheduled,
// 			Image: "strm/helloworld-http",
// 		}
// 		te := task.TaskEvent{
// 			ID:    uuid.New(),
// 			State: task.Running,
// 			Task:  t,
// 		}
// 		m.AddTask(te)
// 		m.SendWork()
// 	}

// 	go func() {
// 		for {
// 			fmt.Printf("[Manager] Updating tasks from %d workers\n", len(m.Workers))
// 			m.UpdateTasks()
// 			time.Sleep(15 * time.Second)
// 		}
// 	}()

// 	for {
// 		for _, t := range m.TaskDb {
// 			fmt.Printf("[Manager] Task: id: %s, state: %d\n", t.ID, t.State)
// 			time.Sleep(15 * time.Second)
// 		}
// 	}
// }
