package models

import (
	"net/url"
	"regexp"

	"github.com/pivotal-golang/lager"
)

var taskGuidPattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

type TaskFilter struct {
	Domain string
	CellID string
}

func (t *Task) LagerData() lager.Data {
	return lager.Data{
		"task-guid": t.TaskGuid,
		"domain":    t.Domain,
		"state":     t.State,
		"cell-id":   t.CellId,
	}
}

func (task *Task) Validate() error {
	var validationError ValidationError

	if task.Domain == "" {
		validationError = validationError.Append(ErrInvalidField{"domain"})
	}

	if !taskGuidPattern.MatchString(task.TaskGuid) {
		validationError = validationError.Append(ErrInvalidField{"task_guid"})
	}

	if task.TaskDefinition == nil {
		validationError = validationError.Append(ErrInvalidField{"task_definition"})
	} else if defErr := task.TaskDefinition.Validate(); defErr != nil {
		validationError = validationError.Append(defErr)
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}

func (t *Task) Decode(p *Payload) error {
	var err error

	switch p.Version {
	case V0:
		err = FromJSON(p.Payload, t)
	case V1:
		err = t.Unmarshal(p.Payload)
	default:
		panic("unknown version")
	}

	if err != nil {
		return err
	}
	return t.Validate()
}

func (def *TaskDefinition) Validate() error {
	var validationError ValidationError

	if def.RootFs == "" {
		validationError = validationError.Append(ErrInvalidField{"rootfs"})
	} else {
		rootFsURL, err := url.Parse(def.RootFs)
		if err != nil || rootFsURL.Scheme == "" {
			validationError = validationError.Append(ErrInvalidField{"rootfs"})
		}
	}

	action := UnwrapAction(def.Action)
	if action == nil {
		validationError = validationError.Append(ErrInvalidActionType)
	} else {
		err := action.Validate()
		if err != nil {
			validationError = validationError.Append(err)
		}
	}

	if def.CpuWeight > 100 {
		validationError = validationError.Append(ErrInvalidField{"cpu_weight"})
	}

	if len(def.Annotation) > maximumAnnotationLength {
		validationError = validationError.Append(ErrInvalidField{"annotation"})
	}

	for _, rule := range def.EgressRules {
		err := rule.Validate()
		if err != nil {
			validationError = validationError.Append(ErrInvalidField{"egress_rules"})
		}
	}

	if !validationError.Empty() {
		return validationError
	}

	return nil
}
