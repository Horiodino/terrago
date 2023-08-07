package api

type K8sErrors interface {
	MEtriceClient(err error)
}

type Errors struct {
	K8sErrors
}

func (E *Errors) MEtriceClient(err error) {
	if err != nil {
		panic(err)
	}
}
