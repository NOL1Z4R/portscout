package domain

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
	StatusCancelled Status = "cancelled"

)

type Scan struct{
	ID 			string
	Target     	string
    Ports      	[]int
    Status     	Status
    Results    	[]PortResult
    CreatedAt  	time.Time
    StartedAt  	*time.Time
    FinishedAt 	*time.Time
}

type PortResult struct {
	Port	int
	State 	string	
	Service string
}