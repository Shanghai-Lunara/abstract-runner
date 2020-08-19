package proto

import (
	"sync"
	"time"
)

type Runner struct {
	mu sync.RWMutex

	Status RunnerStatus `json:"status"`

	CreateTm           time.Time `json:"createTm"`
	LastChangeStatusTm time.Time `json:"lastChangeStatusTm"`
}

type Project struct {
	// The two fields bellow were used for starting the specific project
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

// Git
type Git struct {
	GitAddress string `json:"gitAddress"`
	BranchName string `json:"branchName"`
	CommitHash string `json:"commitHash"`
	DataPath   string `json:"dataPath"`
	WorkerPath string `json:"workerPath"`
}

type TaskList struct {
	Items []Task `json:"items"`
}

type Task struct {
	Project    Project    `json:"project"`
	Status     TaskStatus `json:"status"`
	StartingTm time.Time  `json:"startingTm"`
	EndingTm   time.Time  `json:"endingTm"`
}
