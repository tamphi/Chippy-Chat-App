import React, { useState, useRef, useContext, useEffect } from 'react'
import ChatBody from '../../components/chat_body'
import { WebsocketContext } from '../../modules/websocket_provider'
import { useRouter } from 'next/router'
import { API_URL } from '../../constants'
import autosize from 'autosize'
import { AuthContext } from '../../modules/auth_provider'

export type Message = {
  content: string
  client_id: string
  username: string
  room_id: string
  type: 'recv' | 'self'
}

const index = () => {
  const [messages, setMessage] = useState<Array<Message>>([])
  const textarea = useRef<HTMLTextAreaElement>(null)
  const { conn } = useContext(WebsocketContext)
  const { user } = useContext(AuthContext)

  const router = useRouter()
  useEffect(() => {
    console.log("conn is: ", conn)
    if (conn === null) {
      router.push('/chats')
      return
    }
    (async function() {
      try{
          const lastIndexUrl = conn.url.lastIndexOf("/")
          const response = await fetch(`${API_URL}/chat${conn.url.slice(lastIndexUrl)}`,{
            method: 'GET',
            headers: {'Content-Type':'application/json'}
          })
    
          const data = await response.json()
          if (response.ok){
            console.log("GET OK")
            const allJSONMessages = data.data
            const allMessages = allJSONMessages.map(function(message:any){
              const jsonMessage = JSON.stringify(message)
              const m : Message=JSON.parse(jsonMessage)
              user?.username == m.username ? (m.type = 'self') : (m.type = 'recv')
              return m
            })
            setMessage([...messages, ...allMessages])
          }
    
        } catch (err){
          console.log(err)
        }
  }())
  }, [])

  useEffect(() => {
    if (textarea.current) {
      autosize(textarea.current)
    }

    if (conn === null) {
      router.push('/chats')
      return
    }

    conn.onmessage = (message) => {
      const m: Message = JSON.parse(message.data)
      console.log("message.data: ",message.data)
      user?.username == m.username ? (m.type = 'self') : (m.type = 'recv')
      setMessage([...messages, m])
    }

    conn.onclose = () => {}
    conn.onerror = () => {}
    conn.onopen = () => {}
  }, [textarea, messages, conn])

  const sendMessageButton = () => {
    if (!textarea.current?.value) return
    if (conn === null) {
      router.push('/')
      return
    }

    conn.send(textarea.current.value)
    textarea.current.value = ''
  }

  return (
    <>
      <div className='flex flex-col w-full'>
        <div className='p-4 md:mx-6 mb-14'>
          <ChatBody data={messages} />
        </div>
        <div className='fixed bottom-0 mt-4 w-full'>
          <div className='flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md'>
            <div className='flex w-full mr-4 rounded-md border border-green'>
              <textarea
                ref={textarea}
                placeholder='type your message here'
                className='w-full h-10 p-2 rounded-md focus:outline-none'
                style={{ resize: 'none' }}
              />
            </div>
            <div className='flex items-center'>
              <button
                className='p-2 rounded-md bg-green text-white'
                onClick={sendMessageButton}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default index