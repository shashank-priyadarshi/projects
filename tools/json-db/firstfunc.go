package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}
	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s'(database already exists)\n", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record(no name)!")
	}
	mutex := d.getOrCreateMutex()
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tempPath := finalPath + ".tmp"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b = append(b, byte('\n'))
	if err := ioutil.WriteFile(tempPath, b, 0644); err != nil {
		return err
	}
	return nil
}

func (d *Driver) Read() error {
	return nil
}

func (d *Driver) ReadAll() {}

func (d *Driver) Delete() error {
	return nil
}

func (d *Driver) getOrCreateMutex() *sync.Mutex {
	return &sync.Mutex{}
}

func Stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}
