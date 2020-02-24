package scheduler

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	"strconv"
	"strings"
)

// ExtenderArgs is copied from https://godoc.org/k8s.io/kubernetes/pkg/scheduler/api/v1#ExtenderArgs
type ExtenderArgs struct {
	// Pod being scheduled
	Pod *apiv1.Pod `json:"pod"`
	// List of candidate nodes where the pod can be scheduled; to be populated
	// only if ExtenderConfig.NodeCacheCapable == false
	Nodes *apiv1.NodeList `json:"nodes,omitempty"`
	// List of candidate node names where the pod can be scheduled; to be
	// populated only if ExtenderConfig.NodeCacheCapable == true
	NodeNames *[]string `json:"nodenames,omitempty"`
}

func (args ExtenderArgs) String() string {
	str := fmt.Sprintf("Pod:%s\n", args.Pod.Name)
	if args.NodeNames != nil {
		str += fmt.Sprintf("Nodes:%s\n", strings.Join(*args.NodeNames, ","))
	}
	if args.Nodes != nil {
		for i, node := range args.Nodes.Items {
			capacity, ok := node.Annotations["topolvm.cybozu.com/capacity"]
			if ok {
				c, _ := strconv.ParseInt(capacity, 10, 64)
				str += fmt.Sprintf("Nodelist[%d]:%s  Capacity:%s (%dGB)\n", i, node.Name, capacity, c >> 30)
			} else {
				str += fmt.Sprintf("Nodelist[%d]:%s  No capacity\n", i, node.Name)
			}
		}
	}
	return str
}


// HostPriority is copied from https://godoc.org/k8s.io/kubernetes/pkg/scheduler/api/v1#HostPriority
type HostPriority struct {
	// Name of the host
	Host string `json:"host"`
	// Score associated with the host
	Score int `json:"score"`
}

func (hp HostPriority) String() string {
	return fmt.Sprintf("host:%s   priority:%d", hp.Host, hp.Score)
}

// HostPriorityList is copied from https://godoc.org/k8s.io/kubernetes/pkg/scheduler/api/v1#HostPriorityList
type HostPriorityList []HostPriority

func  HostPriorityList2String(hpl []HostPriority) string {
	str := ""
	for _, hp := range hpl {
		str += hp.String() + "\n"
	}
	return str
}

// ExtenderFilterResult is copied from https://godoc.org/k8s.io/kubernetes/pkg/scheduler/api/v1#ExtenderFilterResult
type ExtenderFilterResult struct {
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if ExtenderConfig.NodeCacheCapable == false
	Nodes *apiv1.NodeList `json:"nodes,omitempty"`
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if ExtenderConfig.NodeCacheCapable == true
	NodeNames *[]string `json:"nodenames,omitempty"`
	// Filtered out nodes where the pod can't be scheduled and the failure messages
	FailedNodes FailedNodesMap `json:"failedNodes,omitempty"`
	// Error message indicating failure
	Error string `json:"error,omitempty"`
}


func (args ExtenderFilterResult) String() string {
	str := ""
	if args.NodeNames != nil {
		str += fmt.Sprintf("Nodes:%s\n", strings.Join(*args.NodeNames, ","))
	}
	if args.Nodes != nil {
		for i, node := range args.Nodes.Items {
			capacity, ok := node.Annotations["topolvm.cybozu.com/capacity"]
			if ok {
				c, _ := strconv.ParseInt(capacity, 10, 64)
				str += fmt.Sprintf("Nodelist[%d]:%s  Capacity:%s (%dGB)\n", i, node.Name, capacity, c >> 30)
			} else {
				str += fmt.Sprintf("Nodelist[%d]:%s  No capacity\n", i, node.Name)
			}
		}
	}
	str += fmt.Sprintf("FailedNodes:%v\n", args.FailedNodes)
	str += fmt.Sprintf("Error:%s", args.Error)
	return str
}


// FailedNodesMap is copied from https://godoc.org/k8s.io/kubernetes/pkg/scheduler/api/v1#FailedNodesMap
type FailedNodesMap map[string]string

