import { Attribute, Message } from "../types"
import React, { useState } from "react"
import AttributeItem from "./AttributeItem"

const SubmitMessage = (props: { onSubmit: (m: Message) => void }) => {
  const [topicName, setTopicName] = useState('')
  const [data, setData] = useState('')
  const [attributes, setAttributes] = useState<Attribute[]>([])
  const handleSubmit = async () => {
    if (topicName === '') return
    props.onSubmit({
      id: '',
      topic: topicName,
      data,
      attributes: attributesToObject()
    })
    setTopicName('')
    setData('')
    setAttributes([])
  }

  const attributesToObject = (): { [key: string]: string } => {
    let temp: { [key: string]: string } = {}
    attributes.map(a => temp[`${a.key}`] = a.value)
    return temp
  }

  const handleAttributeRemove = (id: number) => {
    console.log(id)
    setAttributes(prev => prev.filter(a => a.id !== id))
  }

  const handleAttributeChange = (id: number) => {
    return (key: string, value: string) => {
      console.log(id, key, value)
      setAttributes(prev => prev.map(a => a.id === id ? { id, key, value } : a))
    }
  }

  return (
    <>
      <div className="shadow m-4 p-2">
        <h2 className="text-xl font-bold mb-3 border-b-solid border-b-2 border-gray-300">Submit Message</h2>
        <form className="flex flex-col items-start gap-2 p-2 mb-3">
          <label htmlFor="topicname">Topic Name</label>
          <input className="p-2 border-2 rounded w-full lg:w-1/2" id="topicname" type="text" placeholder="topic name" value={topicName} onChange={e => setTopicName(e.target.value)}></input>
          <label htmlFor="data">Data</label>
          <textarea className="p-2 border-2 rounded w-full lg:w-1/2" id="data" placeholder="message data" value={data} onChange={e => setData(e.target.value)}></textarea>
          <label htmlFor="attributes">Attributes</label>
          {
            attributes.map(a => <AttributeItem key={a.id} onRemove={() => handleAttributeRemove(a.id)} onChange={handleAttributeChange(a.id)} />)
          }
          <input className="text-white px-2 py-1 bg-yellow-400 hover:bg-yellow-500 rounded" type="button" value="Add" onClick={() => setAttributes(prev => [...prev, { key: '', value: '', id: Math.random() }])} />
          <input className="text-white px-2 py-1 bg-yellow-400 hover:bg-yellow-500 rounded" type="button" value="Submit" onClick={handleSubmit} />
        </form>
      </div>
    </>
  )
}

export default SubmitMessage