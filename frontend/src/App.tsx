import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
// import Home from './pages/home/Home';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';
import LandingPage from './pages/LandingPage/LandingPage'

function App() {
  return (
    <Router basename='/jobseekers'>
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  </Router>
  );
}

export default App
