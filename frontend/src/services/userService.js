const BASE_URL = "http://localhost:8080/users";

export const getUsers = async () => {
  try {
    const response = await fetch(BASE_URL);

    if (!response.ok) {
      const errorData = await response.json().catch(() => {
        throw new Error("Unexpected response format from server");
      });
      throw new Error(errorData.error || "Failed to fetch users");
    }

    return await response.json();
  } catch (error) {
    console.error("Error in getUsers:", error.message);
    throw new Error(
      "An error occurred while fetching users. Please try again."
    );
  }
};

export const createUser = async (user) => {
  try {
    const response = await fetch(`${BASE_URL}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => {
        throw new Error("Unexpected response format from server");
      });

      if (response.status === 409) {
        throw new Error("User already exists");
      }

      throw new Error(errorData.error || "Failed to create user");
    }

    return await response.json();
  } catch (error) {
    console.error("Error in createUser:", error.message);
    throw error;
  }
};

export const updateUser = async (id, user) => {
  try {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(user),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => {
        throw new Error("Unexpected response format from server");
      });
      throw new Error(errorData.error || "Failed to update user");
    }

    return await response.json();
  } catch (error) {
    console.error("Error in updateUser:", error.message);
    throw new Error("An error occurred while updating the user.");
  }
};

export const deleteUser = async (id) => {
  try {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => {
        throw new Error("Unexpected response format from server");
      });
      throw new Error(errorData.error || "Failed to delete user");
    }
  } catch (error) {
    console.error("Error in deleteUser:", error.message);
    throw new Error("An error occurred while deleting the user.");
  }
};
