import React, { useState, useEffect } from "react";
import { createUser, updateUser } from "../services/userService";
import { TextField, Button } from "@mui/material";

function UserForm({ user, onClose, refreshUsers }) {
  const [formData, setFormData] = useState({ user_name: "", email: "" });

  useEffect(() => {
    if (user) {
      setFormData(user);
    }
  }, [user]);

  const handleSubmit = async (e) => {
    e.preventDefault(); // Prevent form submission reload
    try {
      if (user) {
        await updateUser(user.id, formData);
      } else {
        await createUser(formData);
      }
      refreshUsers();
      onClose();
    } catch (error) {
      console.error("Error creating/updating user:", error);
      if (error.message === "User already exists") {
        alert("User already exists. Please use a different username.");
      } else {
        alert(error.message || "Failed to create/update user.");
      }
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        label="User Name"
        value={formData.user_name}
        onChange={(e) =>
          setFormData({ ...formData, user_name: e.target.value })
        }
        required
        fullWidth
        margin="normal"
      />
      <TextField
        label="Email"
        value={formData.email}
        onChange={(e) => setFormData({ ...formData, email: e.target.value })}
        required
        fullWidth
        margin="normal"
        type="email"
      />
      <Button
        type="submit"
        variant="contained"
        color="primary"
        style={{ marginTop: "10px" }}
      >
        {user ? "Update" : "Create"}
      </Button>
    </form>
  );
}

export default UserForm;
