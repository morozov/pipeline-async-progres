package main

import (
	"fmt"
	"sync"

	"github.com/gosuri/uilive"
)

// Progress https://github.com/gosuri/uiprogress/blob/484b9f69ea000422e1873db136dbb80e30b5de3c/progress.go
type Progress struct {
	jobs []*Job

	writer *uilive.Writer
	update chan struct{}
	done   chan struct{}
	mtx    *sync.RWMutex
}

func NewProgress() *Progress {
	writer := uilive.New()

	return &Progress{
		writer: writer,
		done:   make(chan struct{}),
		update: make(chan struct{}),
		mtx:    &sync.RWMutex{},
	}
}

func (p *Progress) AddJob(job *Job) {
	p.jobs = append(p.jobs, job)
}

func (p *Progress) Start() {
	p.writer.Start()
	p.print()
	go p.listen()
}

func (p *Progress) Stop() {
	p.done <- struct{}{}
	p.writer.Stop()
}

func (p *Progress) Update() {
	p.update <- struct{}{}
}

func (p *Progress) listen() {
	for {
		select {
		case <-p.update:
			p.print()
		case <-p.done:
			close(p.done)
			return
		}
	}
}

func (p *Progress) print() {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	for _, j := range p.jobs {
		fmt.Fprintln(p.writer, j.String())
	}
	p.writer.Flush()
}

