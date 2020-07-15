package kodfs_nameserver

type NameNode struct {
	Nameserver_name string

	Nameserver_ip string

	Nameserver_port int
}

func NewNameNode(nameserver_ip string, nameserver_port int) *NameNode {
	return &NameNode{Nameserver_ip: nameserver_ip, Nameserver_port: nameserver_port}
}
