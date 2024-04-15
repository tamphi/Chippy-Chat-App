import { useState, useContext } from "react"
import { useEffect } from "react"
import { API_URL } from '@/constants'
import { useRouter } from "next/router"
import { UserInfo, AuthContext } from "@/modules/auth_provider"

const index = () => {
  const router = useRouter()
  const {authenticated} = useContext(AuthContext)
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  const buttonController = async (e: React.SyntheticEvent) => {
    //todo: implement submit button handler --> done
    e.preventDefault()

    try{
      const response = await fetch(`${API_URL}/signup`,{
        method: 'POST',
        headers: {'Content-Type':'application/json'},
        body: JSON.stringify({username,password})
      })

      if (response.ok){
        return router.push('/login')
      }

    } catch (err){
      console.log(err)
    }

  }
  

  return (
    <div className='flex flex-col items-center justify-center min-w-full min-h-screen'>
      <form className='flex flex-col md:w-1/5'>

        <div className='text-3xl font-bold text-center'>
          <span className='text-green'>Chippy Chat!</span>
        </div>

        <input 
        placeholder='username' 
        className='p-3 mt-8 rounded-md border-2 border-black focus:outline-none focus:border-green'
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        />

        <input 
        type='password'
        placeholder='password' 
        className='p-3 mt-4 rounded-md border-2 border-black focus:outline-none focus:border-green'
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        />

        <button 
        type="submit"
        className='p-3 mt-6 rounded-md bg-green font-bold text-white'
        onClick={buttonController}
        >
          Signup
        </button>
      
      </form>
    </div>
  )
}

export default index