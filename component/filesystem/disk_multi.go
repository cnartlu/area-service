package filesystem

import (
	"strings"
	sync "sync"
)

type multiResults struct {
	results map[string]Result
}

func (m multiResults) Name() string {
	return "multiResults"
}

func (m multiResults) Raw() any {
	return m
}

func (m multiResults) GetResult(key string) Result {
	if r, ok := m.results[key]; ok {
		return r
	}
	return nil
}

func (m multiResults) GetResults() map[string]Result {
	return m.results
}

func IsMultiResult(r Result) bool {
	_, ok := r.(*multiResults)
	return ok
}

func FromMultiResult(r Result) multiResults {
	result, ok := r.(multiResults)
	if !ok {
		return multiResults{}
	}
	return result
}

type multiErrors map[string]error

func (m multiErrors) Error() string {
	var s string
	for k, e := range m {
		s += k + ": " + e.Error() + "\n"
	}
	s = strings.TrimRight(s, "\n")
	return s
}

type multiDisks map[string]Disk

func (m multiDisks) Len() int {
	return len(m)
}

func (m multiDisks) Append(name string, disk Disk) {
	m[name] = disk
}

func (m multiDisks) GetDisk(name string) Disk {
	if name == "" {
		return m
	}
	if r, ok := m[name]; ok {
		return r
	}
	return nil
}

func (m multiDisks) Exists(key string, options ...HandleFunc) bool {
	for _, driver := range m {
		driver := driver
		if driver.Exists(key, options...) {
			return true
		}
	}
	return false
}

func (m multiDisks) Upload(filename string, key string, options ...HandleFunc) (Result, error) {
	var err multiErrors
	var wg = sync.WaitGroup{}
	var result = &multiResults{
		results: make(map[string]Result),
	}
	for driverName, driver := range m {
		driver := driver
		driverName := driverName
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err1 := driver.Upload(filename, key, options...)
			if err1 != nil {
				if err == nil {
					err = make(map[string]error)
				}
				err[driverName] = err
				return
			}
			result.results[driverName] = r
		}()
	}
	wg.Wait()
	return result, err
}

func (m multiDisks) Url(key string, options ...HandleFunc) string {
	for _, driver := range m {
		driver := driver
		if url := driver.Url(key, options...); url != "" {
			return url
		}
	}
	return ""
}

func (m multiDisks) Delete(key string, options ...HandleFunc) error {
	var err multiErrors
	var wg = sync.WaitGroup{}
	for driverName, driver := range m {
		driver := driver
		driverName := driverName
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err1 := driver.Delete(key, options...); err1 != nil {
				if err == nil {
					err = make(map[string]error)
				}
				err[driverName] = err1
			}
		}()
	}
	wg.Wait()
	return err
}

func NewMultiDisks(disks map[string]Disk) Disk {
	var result multiDisks = disks
	return result
}
