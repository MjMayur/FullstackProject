import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Card, CardBody, Form, FormGroup, Input, Button } from "reactstrap";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function UserForm() {
  const token = localStorage.getItem("authToken") || "";
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    phone: "",
    message: "",
  });
  const [isUpdate, setIsUpdate] = useState({
    id: 0,
    update: false,
  });
  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Here you can add your form submission logic
  };
  const [records, setRecords] = useState([]);

  React.useEffect(() => {
    if (!token) {
      navigate("/login");
    }
  }, []);

  const fetchData = () => {
    const apiUrl = "http://localhost:8000/user/list/";
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    fetch(apiUrl, {
      method: "GET",
      headers: headers,
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        setRecords(res.data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  React.useEffect(() => {
    // Fetch API
    fetchData();
  }, []);

  const handleEdit = (id) => {
    const apiUrl = `http://localhost:8000/user/get/${id}`;

    // Fetch API
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    fetch(apiUrl, {
      method: "GET",
      headers: headers,
    })
      .then((res) => res.json())
      .then((res) => {
        setIsUpdate({ id: id, update: true });
        console.log(res);
        setFormData({
          name: res.data.Name,
          email: res.data.Email,
          phone: res.data.Phone,
          message: res.data.Message,
        });
      })
      .catch((error) => {
        console.log(error);
      });
    fetchData();
  };

  const handleDelete = (id) => {
    const apiUrl = `http://localhost:8000/user/delete/${id}`;

    // Fetch API
    let headers = new Headers();
    headers.append("Content-Type", "application/json");
    fetch(apiUrl, {
      method: "DELETE",
      headers: headers,
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.code === 200) {
          toast.success("Record deleted successfully!");
        } else {
          toast.error("Failed to fetch data!");
        }
      })
      .catch((error) => {
        console.log(error);
      });
    fetchData();
  };

  const submitForm = () => {
    const apiUrl = isUpdate.update
      ? `http://localhost:8000/user/update/${isUpdate.id}`
      : `http://localhost:8000/user/add/`;

    const method = isUpdate.update ? "PATCH" : "POST";
    fetch(apiUrl, {
      method: method,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(formData),
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.code === 200) {
          toast.success(
            isUpdate.update
              ? "Record updated successfully!"
              : "Record added successfully!"
          );
          setFormData({ name: "", email: "", phone: "", message: "" });
          setIsUpdate({ id: 0, update: false });
          fetchData();
        } else {
          toast.error("Something went wrong. Please try again.");
        }
      })
      .catch((error) => {
        toast.error("Failed to submit the form!");
        console.error(error);
      });
  };

  return (
    <>
      <Card
        style={{
          maxWidth: "500px",
          margin: "auto",
          marginTop: "3rem",
          padding: "20px",
          borderRadius: "10px",
          boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
        }}
      >
        <CardBody className="p-2">
          <h2 style={{ textAlign: "center", marginBottom: "2rem" }}>
            Contact Us
          </h2>
          <Form onSubmit={handleSubmit}>
            <FormGroup>
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
              <Input
                type="tel"
                name="phone"
                id="phone"
                placeholder="Enter your phone number"
                value={formData.phone}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <FormGroup>
              <Input
                type="textarea"
                name="message"
                id="message"
                placeholder="Enter your message"
                value={formData.message}
                onChange={handleChange}
                required
              />
            </FormGroup>
            <Button
              color="primary"
              style={{ width: "100%" }}
              onClick={submitForm}
            >
              {isUpdate.update ? "Update" : "Submit"}
            </Button>
          </Form>
        </CardBody>
      </Card>
      <div
        style={{ padding: "20px", display: "flex", justifyContent: "center" }}
      >
        <table
          style={{
            width: "80%",
            borderCollapse: "collapse",
            boxShadow: "0 2px 10px rgba(0,0,0,0.1)",
          }}
        >
          <thead>
            <tr style={{ backgroundColor: "#f2f2f2" }}>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Name
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Email
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Phone
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Message
              </th>
              <th
                style={{
                  padding: "10px",
                  border: "1px solid #dddddd",
                  textAlign: "center",
                }}
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {records.map((record) => (
              <tr key={record.id} style={{ backgroundColor: "#ffffff" }}>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.name}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.email}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.phone}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  {record.message}
                </td>
                <td style={{ padding: "10px", border: "1px solid #dddddd" }}>
                  <button
                    onClick={() => handleEdit(record.ID)}
                    style={{ marginRight: "10px" }}
                  >
                    Edit
                  </button>
                  <button onClick={() => handleDelete(record.ID)}>
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
}

export default UserForm;
