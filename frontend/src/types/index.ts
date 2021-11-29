export type Message = {
  id: string
  topic: string
  attributes: {
    [key: string]: string
  },
  data: string
}

export type ReceiveMessageResponse = {
  messages: Message[]
}

export type Attribute = {
  id: number
  key: string
  value: string
}