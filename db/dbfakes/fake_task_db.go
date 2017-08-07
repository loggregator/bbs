// This file was generated by counterfeiter
package dbfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/auctioneer"
	"code.cloudfoundry.org/bbs/db"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
)

type FakeTaskDB struct {
	TasksStub        func(logger lager.Logger, filter models.TaskFilter) ([]*models.Task, error)
	tasksMutex       sync.RWMutex
	tasksArgsForCall []struct {
		logger lager.Logger
		filter models.TaskFilter
	}
	tasksReturns struct {
		result1 []*models.Task
		result2 error
	}
	TaskByGuidStub        func(logger lager.Logger, taskGuid string) (*models.Task, error)
	taskByGuidMutex       sync.RWMutex
	taskByGuidArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	taskByGuidReturns struct {
		result1 *models.Task
		result2 error
	}
	DesireTaskStub        func(logger lager.Logger, taskDefinition *models.TaskDefinition, taskGuid, domain string) (*models.Task, error)
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		logger         lager.Logger
		taskDefinition *models.TaskDefinition
		taskGuid       string
		domain         string
	}
	desireTaskReturns struct {
		result1 *models.Task
		result2 error
	}
	StartTaskStub        func(logger lager.Logger, taskGuid, cellId string) (before *models.Task, after *models.Task, shouldStart bool, rr error)
	startTaskMutex       sync.RWMutex
	startTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
		cellId   string
	}
	startTaskReturns struct {
		result1 *models.Task
		result2 *models.Task
		result3 bool
		result4 error
	}
	CancelTaskStub        func(logger lager.Logger, taskGuid string) (before *models.Task, after *models.Task, cellID string, err error)
	cancelTaskMutex       sync.RWMutex
	cancelTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	cancelTaskReturns struct {
		result1 *models.Task
		result2 *models.Task
		result3 string
		result4 error
	}
	FailTaskStub        func(logger lager.Logger, taskGuid, failureReason string) (before *models.Task, after *models.Task, err error)
	failTaskMutex       sync.RWMutex
	failTaskArgsForCall []struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}
	failTaskReturns struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}
	CompleteTaskStub        func(logger lager.Logger, taskGuid, cellId string, failed bool, failureReason, result string) (before *models.Task, after *models.Task, err error)
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		logger        lager.Logger
		taskGuid      string
		cellId        string
		failed        bool
		failureReason string
		result        string
	}
	completeTaskReturns struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}
	ResolvingTaskStub        func(logger lager.Logger, taskGuid string) (before *models.Task, after *models.Task, err error)
	resolvingTaskMutex       sync.RWMutex
	resolvingTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	resolvingTaskReturns struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}
	DeleteTaskStub        func(logger lager.Logger, taskGuid string) (task *models.Task, err error)
	deleteTaskMutex       sync.RWMutex
	deleteTaskArgsForCall []struct {
		logger   lager.Logger
		taskGuid string
	}
	deleteTaskReturns struct {
		result1 *models.Task
		result2 error
	}
	ConvergeTasksStub        func(logger lager.Logger, cellSet models.CellSet, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration time.Duration) (tasksToAuction []*auctioneer.TaskStartRequest, tasksToComplete []*models.Task, taskEvents []models.Event)
	convergeTasksMutex       sync.RWMutex
	convergeTasksArgsForCall []struct {
		logger                      lager.Logger
		cellSet                     models.CellSet
		kickTaskDuration            time.Duration
		expirePendingTaskDuration   time.Duration
		expireCompletedTaskDuration time.Duration
	}
	convergeTasksReturns struct {
		result1 []*auctioneer.TaskStartRequest
		result2 []*models.Task
		result3 []models.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskDB) Tasks(logger lager.Logger, filter models.TaskFilter) ([]*models.Task, error) {
	fake.tasksMutex.Lock()
	fake.tasksArgsForCall = append(fake.tasksArgsForCall, struct {
		logger lager.Logger
		filter models.TaskFilter
	}{logger, filter})
	fake.recordInvocation("Tasks", []interface{}{logger, filter})
	fake.tasksMutex.Unlock()
	if fake.TasksStub != nil {
		return fake.TasksStub(logger, filter)
	} else {
		return fake.tasksReturns.result1, fake.tasksReturns.result2
	}
}

func (fake *FakeTaskDB) TasksCallCount() int {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return len(fake.tasksArgsForCall)
}

func (fake *FakeTaskDB) TasksArgsForCall(i int) (lager.Logger, models.TaskFilter) {
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	return fake.tasksArgsForCall[i].logger, fake.tasksArgsForCall[i].filter
}

func (fake *FakeTaskDB) TasksReturns(result1 []*models.Task, result2 error) {
	fake.TasksStub = nil
	fake.tasksReturns = struct {
		result1 []*models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDB) TaskByGuid(logger lager.Logger, taskGuid string) (*models.Task, error) {
	fake.taskByGuidMutex.Lock()
	fake.taskByGuidArgsForCall = append(fake.taskByGuidArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("TaskByGuid", []interface{}{logger, taskGuid})
	fake.taskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(logger, taskGuid)
	} else {
		return fake.taskByGuidReturns.result1, fake.taskByGuidReturns.result2
	}
}

func (fake *FakeTaskDB) TaskByGuidCallCount() int {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return len(fake.taskByGuidArgsForCall)
}

func (fake *FakeTaskDB) TaskByGuidArgsForCall(i int) (lager.Logger, string) {
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	return fake.taskByGuidArgsForCall[i].logger, fake.taskByGuidArgsForCall[i].taskGuid
}

func (fake *FakeTaskDB) TaskByGuidReturns(result1 *models.Task, result2 error) {
	fake.TaskByGuidStub = nil
	fake.taskByGuidReturns = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDB) DesireTask(logger lager.Logger, taskDefinition *models.TaskDefinition, taskGuid string, domain string) (*models.Task, error) {
	fake.desireTaskMutex.Lock()
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		logger         lager.Logger
		taskDefinition *models.TaskDefinition
		taskGuid       string
		domain         string
	}{logger, taskDefinition, taskGuid, domain})
	fake.recordInvocation("DesireTask", []interface{}{logger, taskDefinition, taskGuid, domain})
	fake.desireTaskMutex.Unlock()
	if fake.DesireTaskStub != nil {
		return fake.DesireTaskStub(logger, taskDefinition, taskGuid, domain)
	} else {
		return fake.desireTaskReturns.result1, fake.desireTaskReturns.result2
	}
}

func (fake *FakeTaskDB) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeTaskDB) DesireTaskArgsForCall(i int) (lager.Logger, *models.TaskDefinition, string, string) {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return fake.desireTaskArgsForCall[i].logger, fake.desireTaskArgsForCall[i].taskDefinition, fake.desireTaskArgsForCall[i].taskGuid, fake.desireTaskArgsForCall[i].domain
}

func (fake *FakeTaskDB) DesireTaskReturns(result1 *models.Task, result2 error) {
	fake.DesireTaskStub = nil
	fake.desireTaskReturns = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDB) StartTask(logger lager.Logger, taskGuid string, cellId string) (before *models.Task, after *models.Task, shouldStart bool, rr error) {
	fake.startTaskMutex.Lock()
	fake.startTaskArgsForCall = append(fake.startTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
		cellId   string
	}{logger, taskGuid, cellId})
	fake.recordInvocation("StartTask", []interface{}{logger, taskGuid, cellId})
	fake.startTaskMutex.Unlock()
	if fake.StartTaskStub != nil {
		return fake.StartTaskStub(logger, taskGuid, cellId)
	} else {
		return fake.startTaskReturns.result1, fake.startTaskReturns.result2, fake.startTaskReturns.result3, fake.startTaskReturns.result4
	}
}

func (fake *FakeTaskDB) StartTaskCallCount() int {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return len(fake.startTaskArgsForCall)
}

func (fake *FakeTaskDB) StartTaskArgsForCall(i int) (lager.Logger, string, string) {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return fake.startTaskArgsForCall[i].logger, fake.startTaskArgsForCall[i].taskGuid, fake.startTaskArgsForCall[i].cellId
}

func (fake *FakeTaskDB) StartTaskReturns(result1 *models.Task, result2 *models.Task, result3 bool, result4 error) {
	fake.StartTaskStub = nil
	fake.startTaskReturns = struct {
		result1 *models.Task
		result2 *models.Task
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTaskDB) CancelTask(logger lager.Logger, taskGuid string) (before *models.Task, after *models.Task, cellID string, err error) {
	fake.cancelTaskMutex.Lock()
	fake.cancelTaskArgsForCall = append(fake.cancelTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("CancelTask", []interface{}{logger, taskGuid})
	fake.cancelTaskMutex.Unlock()
	if fake.CancelTaskStub != nil {
		return fake.CancelTaskStub(logger, taskGuid)
	} else {
		return fake.cancelTaskReturns.result1, fake.cancelTaskReturns.result2, fake.cancelTaskReturns.result3, fake.cancelTaskReturns.result4
	}
}

func (fake *FakeTaskDB) CancelTaskCallCount() int {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return len(fake.cancelTaskArgsForCall)
}

func (fake *FakeTaskDB) CancelTaskArgsForCall(i int) (lager.Logger, string) {
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	return fake.cancelTaskArgsForCall[i].logger, fake.cancelTaskArgsForCall[i].taskGuid
}

func (fake *FakeTaskDB) CancelTaskReturns(result1 *models.Task, result2 *models.Task, result3 string, result4 error) {
	fake.CancelTaskStub = nil
	fake.cancelTaskReturns = struct {
		result1 *models.Task
		result2 *models.Task
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTaskDB) FailTask(logger lager.Logger, taskGuid string, failureReason string) (before *models.Task, after *models.Task, err error) {
	fake.failTaskMutex.Lock()
	fake.failTaskArgsForCall = append(fake.failTaskArgsForCall, struct {
		logger        lager.Logger
		taskGuid      string
		failureReason string
	}{logger, taskGuid, failureReason})
	fake.recordInvocation("FailTask", []interface{}{logger, taskGuid, failureReason})
	fake.failTaskMutex.Unlock()
	if fake.FailTaskStub != nil {
		return fake.FailTaskStub(logger, taskGuid, failureReason)
	} else {
		return fake.failTaskReturns.result1, fake.failTaskReturns.result2, fake.failTaskReturns.result3
	}
}

func (fake *FakeTaskDB) FailTaskCallCount() int {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return len(fake.failTaskArgsForCall)
}

func (fake *FakeTaskDB) FailTaskArgsForCall(i int) (lager.Logger, string, string) {
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	return fake.failTaskArgsForCall[i].logger, fake.failTaskArgsForCall[i].taskGuid, fake.failTaskArgsForCall[i].failureReason
}

func (fake *FakeTaskDB) FailTaskReturns(result1 *models.Task, result2 *models.Task, result3 error) {
	fake.FailTaskStub = nil
	fake.failTaskReturns = struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTaskDB) CompleteTask(logger lager.Logger, taskGuid string, cellId string, failed bool, failureReason string, result string) (before *models.Task, after *models.Task, err error) {
	fake.completeTaskMutex.Lock()
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		logger        lager.Logger
		taskGuid      string
		cellId        string
		failed        bool
		failureReason string
		result        string
	}{logger, taskGuid, cellId, failed, failureReason, result})
	fake.recordInvocation("CompleteTask", []interface{}{logger, taskGuid, cellId, failed, failureReason, result})
	fake.completeTaskMutex.Unlock()
	if fake.CompleteTaskStub != nil {
		return fake.CompleteTaskStub(logger, taskGuid, cellId, failed, failureReason, result)
	} else {
		return fake.completeTaskReturns.result1, fake.completeTaskReturns.result2, fake.completeTaskReturns.result3
	}
}

func (fake *FakeTaskDB) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeTaskDB) CompleteTaskArgsForCall(i int) (lager.Logger, string, string, bool, string, string) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return fake.completeTaskArgsForCall[i].logger, fake.completeTaskArgsForCall[i].taskGuid, fake.completeTaskArgsForCall[i].cellId, fake.completeTaskArgsForCall[i].failed, fake.completeTaskArgsForCall[i].failureReason, fake.completeTaskArgsForCall[i].result
}

func (fake *FakeTaskDB) CompleteTaskReturns(result1 *models.Task, result2 *models.Task, result3 error) {
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTaskDB) ResolvingTask(logger lager.Logger, taskGuid string) (before *models.Task, after *models.Task, err error) {
	fake.resolvingTaskMutex.Lock()
	fake.resolvingTaskArgsForCall = append(fake.resolvingTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("ResolvingTask", []interface{}{logger, taskGuid})
	fake.resolvingTaskMutex.Unlock()
	if fake.ResolvingTaskStub != nil {
		return fake.ResolvingTaskStub(logger, taskGuid)
	} else {
		return fake.resolvingTaskReturns.result1, fake.resolvingTaskReturns.result2, fake.resolvingTaskReturns.result3
	}
}

func (fake *FakeTaskDB) ResolvingTaskCallCount() int {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return len(fake.resolvingTaskArgsForCall)
}

func (fake *FakeTaskDB) ResolvingTaskArgsForCall(i int) (lager.Logger, string) {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return fake.resolvingTaskArgsForCall[i].logger, fake.resolvingTaskArgsForCall[i].taskGuid
}

func (fake *FakeTaskDB) ResolvingTaskReturns(result1 *models.Task, result2 *models.Task, result3 error) {
	fake.ResolvingTaskStub = nil
	fake.resolvingTaskReturns = struct {
		result1 *models.Task
		result2 *models.Task
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTaskDB) DeleteTask(logger lager.Logger, taskGuid string) (task *models.Task, err error) {
	fake.deleteTaskMutex.Lock()
	fake.deleteTaskArgsForCall = append(fake.deleteTaskArgsForCall, struct {
		logger   lager.Logger
		taskGuid string
	}{logger, taskGuid})
	fake.recordInvocation("DeleteTask", []interface{}{logger, taskGuid})
	fake.deleteTaskMutex.Unlock()
	if fake.DeleteTaskStub != nil {
		return fake.DeleteTaskStub(logger, taskGuid)
	} else {
		return fake.deleteTaskReturns.result1, fake.deleteTaskReturns.result2
	}
}

func (fake *FakeTaskDB) DeleteTaskCallCount() int {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	return len(fake.deleteTaskArgsForCall)
}

func (fake *FakeTaskDB) DeleteTaskArgsForCall(i int) (lager.Logger, string) {
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	return fake.deleteTaskArgsForCall[i].logger, fake.deleteTaskArgsForCall[i].taskGuid
}

func (fake *FakeTaskDB) DeleteTaskReturns(result1 *models.Task, result2 error) {
	fake.DeleteTaskStub = nil
	fake.deleteTaskReturns = struct {
		result1 *models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeTaskDB) ConvergeTasks(logger lager.Logger, cellSet models.CellSet, kickTaskDuration time.Duration, expirePendingTaskDuration time.Duration, expireCompletedTaskDuration time.Duration) (tasksToAuction []*auctioneer.TaskStartRequest, tasksToComplete []*models.Task, taskEvents []models.Event) {
	fake.convergeTasksMutex.Lock()
	fake.convergeTasksArgsForCall = append(fake.convergeTasksArgsForCall, struct {
		logger                      lager.Logger
		cellSet                     models.CellSet
		kickTaskDuration            time.Duration
		expirePendingTaskDuration   time.Duration
		expireCompletedTaskDuration time.Duration
	}{logger, cellSet, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration})
	fake.recordInvocation("ConvergeTasks", []interface{}{logger, cellSet, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration})
	fake.convergeTasksMutex.Unlock()
	if fake.ConvergeTasksStub != nil {
		return fake.ConvergeTasksStub(logger, cellSet, kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration)
	} else {
		return fake.convergeTasksReturns.result1, fake.convergeTasksReturns.result2, fake.convergeTasksReturns.result3
	}
}

func (fake *FakeTaskDB) ConvergeTasksCallCount() int {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return len(fake.convergeTasksArgsForCall)
}

func (fake *FakeTaskDB) ConvergeTasksArgsForCall(i int) (lager.Logger, models.CellSet, time.Duration, time.Duration, time.Duration) {
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return fake.convergeTasksArgsForCall[i].logger, fake.convergeTasksArgsForCall[i].cellSet, fake.convergeTasksArgsForCall[i].kickTaskDuration, fake.convergeTasksArgsForCall[i].expirePendingTaskDuration, fake.convergeTasksArgsForCall[i].expireCompletedTaskDuration
}

func (fake *FakeTaskDB) ConvergeTasksReturns(result1 []*auctioneer.TaskStartRequest, result2 []*models.Task, result3 []models.Event) {
	fake.ConvergeTasksStub = nil
	fake.convergeTasksReturns = struct {
		result1 []*auctioneer.TaskStartRequest
		result2 []*models.Task
		result3 []models.Event
	}{result1, result2, result3}
}

func (fake *FakeTaskDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tasksMutex.RLock()
	defer fake.tasksMutex.RUnlock()
	fake.taskByGuidMutex.RLock()
	defer fake.taskByGuidMutex.RUnlock()
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	fake.cancelTaskMutex.RLock()
	defer fake.cancelTaskMutex.RUnlock()
	fake.failTaskMutex.RLock()
	defer fake.failTaskMutex.RUnlock()
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	fake.deleteTaskMutex.RLock()
	defer fake.deleteTaskMutex.RUnlock()
	fake.convergeTasksMutex.RLock()
	defer fake.convergeTasksMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTaskDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.TaskDB = new(FakeTaskDB)
