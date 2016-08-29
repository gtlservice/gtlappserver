package server

import "github.com/gtlservice/gutils/logger"
import "github.com/gtlservice/gzkwrapper"

type AppServerHandler interface {
	gzkwrapper.INodeNotifyHandler
}

func (appserver *AppServer) OnZkWrapperNodeHandlerFunc(append_nodes []*gzkwrapper.NodeInfo, remove_nodes []*gzkwrapper.NodeInfo) {
	//AppServer不需要实现NodeHandler回调
}

func (appserver *AppServer) OnZkWrapperPulseHandlerFunc(key string, nodedata *gzkwrapper.NodeData, err error) {
	//zk节点心跳PulseHandler回调
	if err != nil {
		logger.ERROR("[#server#] zkwrapper pulse keepalive error %s, %s", key, err)
		return
	}
	logger.INFO("[#server#] zkwrapper pulse keepalive %s", key)
}

func (appserver *AppServer) OnZkWrapperWatchHandlerFunc(path string, data []byte, err error) {
	//AppServer不需要实现WatchHandler回调
}
