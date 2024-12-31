import React from "react";
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
  const handleLogin = () => {};
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
              required
            />
          </FormGroup>
          <FormGroup>
            <Input
              type="password"
              name="password"
              id="password"
              placeholder="Enter your password"
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
