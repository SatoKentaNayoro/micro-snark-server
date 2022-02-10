package build

const BaseVersion = "v0.0.1"

var CurrentCommit string

type Version struct {
	BaseVersion string
	Commit      string
}

func GetVersion() Version {
	return Version{
		BaseVersion: BaseVersion,
		Commit:      CurrentCommit,
	}
}
