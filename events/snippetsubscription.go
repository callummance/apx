package events

import (
  //"github.com/gin-gonic/gin"
  "github.com/dustin/go-broadcast"
)

var subscribedSnippets = make(map[string]subscribedObject)

func snippet (sid string) subscribedObject {
  s, exist := subscribedSnippets[sid]
  if !exist {
    b := broadcast.NewBroadcaster(10)
    s = subscribedObject{0, b}
    subscribedSnippets[sid] = s
  }
  return s
}

func deleteSnippet (sid string) {
  p, exist := subscribedSnippets[sid]
  if !exist {
    return
  } else {
    p.BroadcastChannel.Close()
    delete(subscribedSnippets, sid)
  }
}

func isSnippetSubscribed(sid string) bool {
  _, exist := subscribedSnippets[sid]
  return exist
}

func SubscribeToSnippet(sid string) chan interface{}{
  newListener := make(chan interface{})
  subSnippet := snippet(sid)
  subSnippet.BroadcastChannel.Register(newListener)
  subSnippet.Subscribers += 1

  return newListener
}

func UnsubscribeFromSnippet(sid string, listener chan interface{}) {
  subSnippet := snippet(sid)
  subSnippet.BroadcastChannel.Unregister(listener)
  if subSnippet.Subscribers <= 0 {
    deleteSnippet(sid)
  }
}

func UpdateSnippet(sid string, diff string) {
  if !isProjSubscribed(sid){
    return
  } else {
    upd8Proj := snippet(sid)
    upd8Proj.BroadcastChannel.Submit(diff)
  }
}
