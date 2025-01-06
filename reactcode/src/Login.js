import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Card,
  CardBody,
  Form,
  FormGroup,
  Input,
  Button,
  CardHeader,
} from "reactstrap";

function Login() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleLogin = (e) => {
    e.preventDefault();
    const apiUrl = `http://localhost:8000/user/login/`;
    // Fetch API
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    fetch(apiUrl, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(formData),
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.code === 200) {
          console.log(res);
          const token = res.data.Token;

          // Store the token in localStorage with a key
          localStorage.setItem("authToken", token);

          // Navigate the user to the "/form" route
          navigate("/form");
        }
      })
      .catch((error) => {
        console.log(error);
      });
  };
  return (
    <Card
      style={{
        maxWidth: "25rem",
        margin: "auto",
        marginTop: "5rem",
        padding: "20px",
        borderRadius: "10px",
        boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
      }}
    >
      <CardHeader className="bg-white mb-4 text-center" tag="h4">
        Login
      </CardHeader>
      <CardBody>
        <Form>
          <FormGroup>
            <Input
              type="email"
              name="email"
              id="email"
              placeholder="Enter your email"
              onChange={handleChange}
              required
            />
          </FormGroup>
          <FormGroup>
            <Input
              type="password"
              name="password"
              id="password"
              placeholder="Enter your password"
              onChange={handleChange}
              required
            />
          </FormGroup>
          <Button
            color="primary"
            style={{ width: "100%", marginBottom: "10px" }}
            onClick={handleLogin}
          >
            Login
          </Button>
        </Form>
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            fontSize: "0.9rem",
            marginTop: "10px",
          }}
        >
          <a
            href="/forgot-password"
            style={{ textDecoration: "none", color: "#007bff" }}
          >
            Forgot Password?
          </a>
          <a
            href="/register"
            style={{ textDecoration: "none", color: "#007bff" }}
          >
            Register
          </a>
        </div>
      </CardBody>
    </Card>
  );
}

export default Login;
