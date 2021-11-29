import { Message } from "../types"
import React, { useState } from "react"
const MessageItem = (props: { message: Message, onAck: () => void }) => {
  const [expand, setExpand] = useState(false)

  const handleExpandToggle = () => {
    setExpand(prev => !prev)
  }

  const handleAck = (e: React.MouseEvent<HTMLInputElement, MouseEvent>) => {
    e.stopPropagation()
    props.onAck()
  }
  return (
    <>
      <div className="flex flex-col py-1 my-1 border-b-solid border-b-2 border-gray-200">
        <div onClick={handleExpandToggle} className="grid grid-cols-12 gap-2">
          <div className="col-span-3 overflow-x-auto">{props.message.id}</div>
          <div className="col-span-3 overflow-x-auto">{props.message.topic}</div>
          <div className="col-span-5 overflow-x-auto">{props.message.data}</div>
          <input className="col-span-1 text-white px-2 py-1 bg-yellow-400 hover:bg-yellow-500 rounded" type="button" value="ack" onClick={e => handleAck(e)} />
        </div>
        {expand &&
          <textarea disabled className="h-60" value={JSON.stringify(props.message, null, 4)} />
        }
      </div>
    </>
  )
}

export default MessageItem