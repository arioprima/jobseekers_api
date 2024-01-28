import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
// import Home from './pages/home/Home';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';

function App() {
  return (
    <Router basename='/jobseekers'>
    <Routes>
      {/* <Route path="/" element={<Home />} /> */}
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  </Router>
  );
}

export default App
