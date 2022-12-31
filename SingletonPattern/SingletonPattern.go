package SingletonPattern

// This code defines a Singleton struct with a single field, data.
// It also defines a global variable instance of type *Singleton, which will hold the singleton instance.

// The GetInstance function is used to retrieve the singleton instance.
// If the instance variable is nil, it creates a new instance of Singleton and assigns it to the instance variable.
// Otherwise, it returns the existing instance.

type Singleton struct {
	data string
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{data: "This is the singleton instance"}
	}
	return instance
}
