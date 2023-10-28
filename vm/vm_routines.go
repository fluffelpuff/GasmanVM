package vm

func (o *ScriptContainerVM) AddNewRoutine() {
	o.wait.Add(1)
}

func (o *ScriptContainerVM) RemoveRoutine() {
	o.wait.Done()
}

func (o *ScriptContainerVM) Wait() {
	o.wait.Wait()
}
