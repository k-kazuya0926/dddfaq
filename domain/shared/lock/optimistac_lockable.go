// 3.7
package lock

type OptimisticLockable interface {
	Version
}

type Version struct {
	value int
}

func InitialVersion() Version {
	return Version{value: 0}
}
