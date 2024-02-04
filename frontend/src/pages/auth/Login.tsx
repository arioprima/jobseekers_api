import React, { useState } from 'react'
import jobImage from '.././../assets/istockphoto-1198240109-1024x1024.jpg'
import { apiLogin } from '../../Service/api'

interface LoginData {
  email: string
  password: string

}

const Login: React.FC = () => {
  const [loginData, setLoginData] = useState<LoginData>({
    email: '',
    password: ''
  })

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setLoginData((prevData) => ({ ...prevData, [name]: value }));
  }
  

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await apiLogin(loginData);
      console.log(response);
      if (response.status === 200) {
        alert('Login Success');
      } else  {
        setLoginData((prevData)=>({...prevData, password: ''}))
        alert(response.message);
      }
    } catch (err) {
      console.log(err);
    }
  }
  
  
  

  return (
    <div>
      <div className="flex justify-center items-center h-screen">
      <div className="flex w-6/12 h-4/6 shadow-2xl">
        {/* kiri */}
          <div className='w-1/2 hidden md:flex shadow-xl relative'>
          <div className="absolute inset-0 bg-black opacity-50"></div>
              <img className='w-full h-full' src={jobImage} />
              <div className="absolute inset-0 flex items-center justify-center text-white">
              <div className="animate-pulse text-6xl font-bold">
                Jobseekers
              </div>
            </div>
          </div>
          {/* kanan */}
          <div className="flex flex-col justify-center items-center w-1/2">
            <h1 className='mb-4 text-2xl font-semibold text-gray-800'>Login</h1>
            <form onSubmit={handleSubmit}>
            <input
              type="email"
              name='email'
              placeholder="Email"
              className="w-72 p-2 mb-4 rounded border-b border-gray-300 focus:outline-none focus:border-blue-500"
              onChange={handleInputChange}
            />
            <input
              type="password"
              name='password'
              placeholder="Password"
              className="w-72 p-2 mb-4 rounded border-b border-gray-300 focus:outline-none focus:border-blue-500"
              value={loginData.password}
              onChange={handleInputChange}
            />
             <button className="w-72 bg-blue-500 text-white  py-2 rounded hover:bg-blue-600 focus:outline-none focus:bg-blue-600">
                Login
              </button>
            </form>
          </div>
      </div>
      </div>
    </div>
  )
}

export default Login