import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "./Login.css";


const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
  
    function validateForm() {
      return email.length > 0 && password.length > 0;
    }
  
    async function handleSubmit(event) {
      event.preventDefault();

      setIsLoading(true);

      try {
        const user = {
          email: fields.email,
          password: fields.password,
        }
  
        // Sign the user in
        await fetch('http://localhost:8080/api/v1/login', {
          method: "POST",
          body: JSON.stringify(user),
          headers: {"Content-type": "application/json; charset=UTF-8"}
        })
        .then(response => {
          localStorage.setItem(response);
        }) 
        .catch(err => console.log(err));
  
        userHasAuthenticated(true);
        // Redirect to the homepage
        history.push("/room");
      } catch (e) {
        alert(e);
        setIsLoading(false);
      }
    }

    return (
        <div className="Login">
          <Form onSubmit={handleSubmit}>
            <Form.Group size="lg" controlId="email">
              <Form.Label>Email</Form.Label>
              <Form.Control
                autoFocus
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </Form.Group>
            <Form.Group size="lg" controlId="password">
              <Form.Label>Password</Form.Label>
              <Form.Control
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </Form.Group>
            <Button block size="lg" type="submit" disabled={!validateForm()}>
              Login
            </Button>
          </Form>
        </div>
      );
};

export default Login;
