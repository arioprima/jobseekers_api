import { Route, Routes as ReactRoutes} from "react-router-dom";
import LandingPage from "../pages/LandingPage/LandingPage";
import Login from "../pages/auth/Login";
import Register from "../pages/auth/Register";
const Routes = () => {
  return (
    <ReactRoutes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </ReactRoutes>
  );
};

export default Routes;