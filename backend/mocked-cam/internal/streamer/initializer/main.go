package initializer

import (
	"encoding/json"
	"github.com/MrTimeGo/smart-parking-lot/backend/mocked-cam/pkg/randq"
	"github.com/pkg/errors"
	"os"
)

type DumpInfo struct {
	CarsToExit []string
}

func InitializeQueues(dumpFile string, carPathes []string) (enterq, exitq *randq.RandomizedQueue[string], err error) {
	raw, err := os.ReadFile(dumpFile)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to read dump file")
	}

	var dump DumpInfo
	if err = json.Unmarshal(raw, &dump); err != nil {
		return nil, nil, errors.Wrap(err, "failed to unmarshal dump file")
	}

	enterq = randq.New[string]()
	exitq = randq.New[string]()

	for _, path := range carPathes {
		match := false

		for _, entered := range dump.CarsToExit {
			if entered == path {
				match = true
				break
			}
		}

		if match {
			enterq.Enqueue(path)
		} else {
			exitq.Enqueue(path)
		}
	}

	return
}

func DumpExitQueue(dumpFile string, exitq *randq.RandomizedQueue[string]) error {
	pathes, err := exitq.BatchDequeue()
	if errors.Is(err, randq.ErrEmptyQueue) {
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "failed to dequeue exit queue")
	}

	var dump = DumpInfo{CarsToExit: pathes}
	raw, err := json.Marshal(dump)
	if err != nil {
		return errors.Wrap(err, "failed to marshal dump info")
	}

	if err = os.WriteFile(dumpFile, raw, 0644); err != nil {
		return errors.Wrap(err, "failed to write dump file")
	}

	return nil
}
