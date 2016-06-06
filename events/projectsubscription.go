package events

import (
  //"github.com/gin-gonic/gin"
  "github.com/dustin/go-broadcast"
)

const broadcastBufLen int = 10

var subscribedProjects = make(map[string]subscribedObject)

type subscribedObject struct {
  Subscribers int
  BroadcastChannel broadcast.Broadcaster
}

func proj (pid string) subscribedObject {
  p, exist := subscribedProjects[pid]
  if !exist {
    b := broadcast.NewBroadcaster(10)
    p = subscribedObject{0, b}
    subscribedProjects[pid] = p
  }
  return p
}

func deleteProj (pid string) {
  p, exist := subscribedProjects[pid]
  if !exist {
    return
  } else {
    p.BroadcastChannel.Close()
    delete(subscribedProjects, pid)
  }
}

func isProjSubscribed(pid string) bool {
  _, exist := subscribedProjects[pid]
  return exist
}

func SubscribeToProject(pid string) chan interface{}{
  newListener := make(chan interface{})
  subProj := proj(pid)
  subProj.BroadcastChannel.Register(newListener)
  subProj.Subscribers += 1

  return newListener
}

func UnsubscribeFromProject(pid string, listener chan interface{}) {
  subProj := proj(pid)
  subProj.BroadcastChannel.Unregister(listener)
  if subProj.Subscribers <= 0 {
    deleteProj(pid)
  }
}

func UpdateProject(pid string, diff string) {
  if !isProjSubscribed(pid){
    return
  } else {
    upd8Proj := proj(pid)
    upd8Proj.BroadcastChannel.Submit(diff)
  }
}
