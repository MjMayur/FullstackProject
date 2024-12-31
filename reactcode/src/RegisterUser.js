import React, { useState } from "react";
import {
  Card,
  CardBody,
  Form,
  FormGroup,
  Label,
  Input,
  Button,
  CardHeader,
} from "reactstrap";

function RegisterUser() {
  const [formData, setFormData] = useState({
    name: "",
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

  const handleSubmit = (e) => {
    e.preventDefault();
    const apiUrl = `http://localhost:8000/user/create/`;
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
        console.log(res);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  return (
    <Card
      style={{
        maxWidth: "500px",
        margin: "auto",
        marginTop: "5rem",
        padding: "20px",
        borderRadius: "10px",
        boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
      }}
    >
      <CardHeader className="bg-white mb-4 text-center">Register</CardHeader>
      <CardBody>
        <Form onSubmit={handleSubmit}>
          <FormGroup>
            <Label for="name">Name</Label>
            <Input
              type="text"
              name="name"
              id="name"
              placeholder="Enter your name"
              value={formData.name}
              onChange={handleChange}
              required
            />
          </FormGroup>
          <FormGroup>
            <Label for="email">Email</Label>
            <Input
              type="email"
              name="email"
              id="email"
              placeholder="Enter your email"
              value={formData.email}
              onChange={handleChange}
              required
            />
          </FormGroup>
          <FormGroup>
            <Label for="password">Password</Label>
            <Input
              type="password"
              name="password"
              id="password"
              placeholder="Enter your password"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </FormGroup>
          <Button color="primary" style={{ width: "100%" }}>
            Register
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
            href="/login"
            style={{
              textDecoration: "none",
              color: "#007bff",
              marginTop: "1rem",
            }}
          >
            Login
          </a>
        </div>
      </CardBody>
    </Card>
  );
}

export default RegisterUser;
