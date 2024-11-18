import React, { useState, useEffect } from "react";
import { getUsers, deleteUser } from "../services/userService";
import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from "@mui/material";
import UserForm from "./UserForm";

function UserList() {
  const [users, setUsers] = useState([]);
  const [open, setOpen] = useState(false);
  const [currentUser, setCurrentUser] = useState(null);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      const data = await getUsers();
      setUsers(Array.isArray(data) ? data : []);
    } catch (error) {
      console.error("Error fetching users:", error);
    }
  };

  const handleDelete = async (id) => {
    try {
      await deleteUser(id);
      fetchUsers();
    } catch (error) {
      console.error("Error deleting user:", error);
    }
  };

  const handleOpenDialog = (user = null) => {
    setCurrentUser(user);
    setOpen(true);
  };

  const handleCloseDialog = () => {
    setOpen(false);
    setCurrentUser(null);
  };

  return (
    <div style={{ padding: "20px" }}>
      <h2>User Management</h2>
      <div style={{ marginBottom: "20px" }}>
        <Button
          onClick={fetchUsers}
          variant="contained"
          color="primary"
          style={{ marginRight: "10px" }}
        >
          Get All Users
        </Button>
        <Button
          onClick={() => handleOpenDialog()}
          variant="contained"
          color="secondary"
        >
          Create User
        </Button>
      </div>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>User Name</TableCell>
              <TableCell>Email</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {users.map((user) => (
              <TableRow key={user.id}>
                <TableCell>{user.id}</TableCell>
                <TableCell>{user.user_name}</TableCell>
                <TableCell>{user.email}</TableCell>
                <TableCell>
                  <Button
                    onClick={() => handleOpenDialog(user)}
                    color="primary"
                  >
                    Update
                  </Button>
                  <Button
                    onClick={() => handleDelete(user.id)}
                    color="secondary"
                  >
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      <Dialog open={open} onClose={handleCloseDialog}>
        <DialogTitle>{currentUser ? "Update User" : "Create User"}</DialogTitle>
        <DialogContent>
          <UserForm
            user={currentUser}
            onClose={handleCloseDialog}
            refreshUsers={fetchUsers}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseDialog} color="primary">
            Cancel
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default UserList;
