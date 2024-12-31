import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./Login";
import UserForm from "./Form";
import RegisterUser from "./RegisterUser";

// import Register from "./Register"; // Create this component for the Register page
// import ForgotPassword from "./ForgotPassword"; // Create this component for the Forgot Password page

function App() {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/form" element={<UserForm />} />
          <Route path="/register" element={<RegisterUser />} />
          {/* <Route path="/forgot-password" element={<ForgotPassword />} /> */}
        </Routes>
      </div>
    </Router>
  );
}

export default App;
