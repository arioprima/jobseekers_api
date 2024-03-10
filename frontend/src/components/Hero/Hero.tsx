import jobImage from '.././../assets/istockphoto-1198240109-1024x1024.jpg'

const Hero = () => {
  return (
    <div id='home' className='pt-36 bg-gradient-to-l from-gray-200 ...'>
      <div className="container">
        <div className='flex flex-wrap-reverse'>
          <div className="w-full self-center px-4 lg:w-1/2  flex flex-wrap md:block">
            <h1 className='font-semibold text-base text-primary md:text-xl lg:text-2xl'>Hello Para pejuang Loker, selamat datang di jobseekers</h1>
            <h2 className='text-secondary mt-2 mb-10 md:text-lg'>Aplikasi Pencari Lowongan kerja terbaik</h2>
            <a href="jobseekers/register" className='text-base font-semibold text-white bg-green-500 py-3 px-8 rounded-xl
                hover:shadow-lg hover:opacity-85 transition duration-300 ease-in-out
                '>Daftar Sekarang</a>
          </div>
          <div className='w-full self-end px-4 lg:w-1/2 mt-[-2rem] lg:mt-6 mb-5'>
            <div className="relative lg:right-0">
              <img src={jobImage} alt="jobseekers" className='max-w-full mx-auto rounded-2xl' />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Hero