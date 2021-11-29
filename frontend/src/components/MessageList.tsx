import { useState } from "react"
import { Message } from "../types"
import MessageItem from "./MessageItem"

const MessageList = (props: { messages: Message[], onPoll: (topicName: string) => void, onAck: (id: string) => void }) => {
  const [topicName, setTopicName] = useState('')

  return (
    <>
      <div className="shadow m-4 p-2">
        <div className="flex flex-col">
          <h2 className="text-xl font-bold mb-3 border-b-solid border-b-2 border-gray-300">Receive Messages</h2>
          <form className="flex flex-row items-center gap-2 p-2 mb-3">
            <label htmlFor="topicname">Topic Name:</label>
            <input className="p-2 border-2 rounded" id="topicname" type="text" placeholder="topic name" value={topicName} onChange={e => setTopicName(e.target.value)}></input>
            <input className="text-white px-2 py-1 bg-yellow-400 hover:bg-yellow-500 rounded" type="button" onClick={() => props.onPoll(topicName)} value="Receive Messages" />
          </form>
          <div className="grid grid-cols-12 gap-2 my-2 border-b-solid border-b-2 border-gray-200">
            <div className="col-span-3 overflow-x-auto">Id</div>
            <div className="col-span-3 overflow-x-auto">Topic</div>
            <div className="col-span-5 overflow-x-auto">Data</div>
            <div className="col-span-1 overflow-x-auto">Ack?</div>
          </div>
          {
            props.messages.map(m => <MessageItem key={m.id} message={m} onAck={() => props.onAck(m.id)} />)
          }
        </div>
      </div>
    </>
  )
}

export default MessageList