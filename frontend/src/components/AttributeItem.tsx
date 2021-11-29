import { Message } from "../types"
import React, { useEffect, useState } from "react"
const AttributeItem = (props: { onRemove: () => void, onChange: (key: string, value: string) => void }) => {
  const [key, setKey] = useState('')
  const [value, setValue] = useState('')

  useEffect(() => {
    props.onChange(key, value)
  }, [key, value])

  return (
    <div className="flex flex-row">
      <input id="key" className="p-1 mx-2 border-2 rounded" type="text" placeholder="key" value={key} onChange={(e) => setKey(e.target.value)} />
      <input id="value" className="p-1 mx-2 border-2 rounded" type="text" placeholder="value" value={value} onChange={(e) => setValue(e.target.value)} />
      <input className="text-white px-2 py-1 bg-yellow-400 hover:bg-yellow-500 rounded" type="button" value="Remove" onClick={props.onRemove} />
    </div>
  )
}

export default AttributeItem