package status

type Status interface {
	Enter()
	NextStatus()
}