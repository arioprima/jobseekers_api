import { BrowserRouter as Router } from "react-router-dom";
import Routes from "./routes/routes";
function App() {
  return (
    <Router basename="/jobseekers">
      <Routes />
    </Router>
  );
}

export default App;
