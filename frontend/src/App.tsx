import React, { useState } from 'react';
import './App.css';
import MessageList from './components/MessageList';
import SubmitMessage from './components/SubmitMessage';
import { API_URL } from './config';
import { Message, ReceiveMessageResponse } from './types';

function App() {
  const [messages, setMessages] = useState<Message[]>([])

  const handleSubmit = async (message: Message) => {
    message.id = (Math.random() * 10).toString()
    try {
      const result = await fetch(`${API_URL}/message`, { body: JSON.stringify(message), method: 'POST' })
      console.log(await result.json())
    } catch (err) {
      console.error(err)
    }
  }

  const handleAck = async (id: string) => {
    setMessages(prev => prev.filter(m => m.id !== id))
    try {
      const result = await fetch(`${API_URL}/message/${id}`, { method: 'PUT' })
      console.log(await result.json())
    } catch (err) {
      console.error(err)
    }
  }

  const handlePoll = async (topicName: string) => {
    console.log(`polling: ${topicName}...`)
    try {
      const result = await fetch(`${API_URL}/topic`, {
        body: JSON.stringify({
          topic: topicName,
          limit: 100,
          attributes: {}
        }), method: 'POST'
      })
      const received = await result.json() as ReceiveMessageResponse
      console.log(received)
      setMessages(prev => received.messages)
    } catch (err) {
      console.error(err)
    }
  }


  return (
    <div>
      <SubmitMessage onSubmit={handleSubmit} />

      <MessageList onPoll={handlePoll} onAck={handleAck} messages={messages} />
    </div>
  );
}

export default App;
