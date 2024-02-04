import React from 'react'
import jobImage from '.././../assets/istockphoto-1198240109-1024x1024.jpg'

const Hero = () => {
  return (
    <div id='home' class='pt-36'>
    <div className="container">
        <div className='flex flex-wrap'>
            <div className="w-full self-center px-4 lg:w-1/2">
                <h1 className='font-semibold text-base text-primary md:text-xl lg:text-2xl'>Hello Para pejuang Loker, selamat datang di jobseekers</h1>
                <h2 className='text-slate-600 mt-5 mb-10 md:text-lg'>Aplikasi Pencari Lowongan kerja terbaik</h2>
                <a href="#" className='text-base font-semibold text-white bg-sky-500 py-3 px-8 rounded-xl
                hover:shadow-lg hover:opacity-85 transition duration-300 ease-in-out
                '>Daftar Sekarang</a>
            </div>
            <div className='w-full self-end px-4 lg:w-1/2'>
              <div className="mt-10">
                <img src={jobImage} alt="jobseekers" className='max-w-full mx-auto' />
              </div>
            </div>
        </div> 
    </div>
    </div>
  )
}

export default Hero