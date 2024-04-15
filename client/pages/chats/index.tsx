import { useState, useContext } from "react"
import { useEffect } from "react"
import { API_URL, WEBSOCKET_URL } from '@/constants'
import { WebsocketContext } from '@/modules/websocket_provider'
import { useRouter } from "next/router"
import { AuthContext } from '@/modules/auth_provider'

const index = () => {
    const router = useRouter()
    const [usernames, setUsernames] = useState([])
    const { conn, setConn } = useContext(WebsocketContext)
    const { user } = useContext(AuthContext)

    useEffect(() => {
        (async function() {
            try{
                const response = await fetch(`${API_URL}/chats`,{
                  method: 'GET',
                  headers: {'Content-Type':'application/json'}
                })
          
                const data = await response.json()
                if (response.ok){
                  setUsernames(data["users"])
                  console.log("GET OK")
                }
          
              } catch (err){
                console.log(err)
              }
        }())
      }, [])
    
    const logOutButton = (username: string) => {
      console.log("goodbye %s", username)
      sessionStorage.removeItem("user_info");
      location.reload()
      return router.push("/login")
    }

    const joinChatButton = (username: string) => {
        var chatIdArray = new Array(user.username,username)
        var sortedChatIdArray = chatIdArray.sort()
        const chatId = sortedChatIdArray.join("__--__")
        const ws = new WebSocket(
            `${WEBSOCKET_URL}/ws/joinChat/${chatId}?username=${user.username}&receiver=${username}`
        )
        if (ws.OPEN) {
            console.log("OPEN")
            console.log(ws)
            setConn(ws)
            console.log("conn is:", conn)
            router.push('/chat')
            return
        }
    }
    return (
    <>
      <div className='my-8 px-4 md:mx-32 w-full h-full'>
        <div className='mt-6'>

          <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6 text-white rounded-md top-0 right-0'>
            <button 
              type="submit"
              className='bg-green rounded-md'
              onClick={() => logOutButton(user.username)}
            > 
              Logout 
            </button>
          </div>

          <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6 text-black rounded-md '>
            Welcome, {user.username}
          </div>
          <div className='font-bold'>Available Users</div>
          <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
            {usernames.map((username, index) => (
                username !== user.username ? 
                <div
                    key={index}
                    className='border border-green p-4 flex items-center rounded-md w-full'
                >
                    <div className='w-full'>
                    <div className='text-sm'>User</div>
                    <div className='text-green font-italic text-lg'>{username}</div>
                    </div>
                    <div className=''>
                    <button
                        className='px-4 text-white bg-green rounded-md'
                        onClick={() => joinChatButton(username)}
                    >
                        Chat
                    </button>
                    </div>
              </div> : 
                null
            )
            )}
          </div>
        </div>
      </div>
    </>
    )
      
}
export default index