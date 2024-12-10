import React, { useState } from "react";
import axios from "axios";

const VerifyArmstrong = () => {
  const [userId, setUserId] = useState("");
  const [number, setNumber] = useState("");
  const [response, setResponse] = useState("");



  const handleSubmit = async (e) => {
    console.log(handleSubmit, "on");
    e.preventDefault();
    try {
      console.log({ userId, number: Number(number) });
  
      const result = await axios.post(
        "http://localhost:8080/verify",
        {
          userId,
          number: Number(number),
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
  
      console.log("Server Response:", result); // Log the full server response
      setResponse("Success: " + (result.data.message || "No message received")); // Set success response
    } catch (error) {
      console.error("Error Response:", error.response); // Log the error response
      setResponse(
        error.response?.data?.message || "An error occurred during verification"
      );
    }
  };
  
  

  return (
    <div>
      <h1>Verify Armstrong Number</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="User ID"
          value={userId}
          onChange={(e) => setUserId(e.target.value)}
        />
        <input
          type="number"
          placeholder="Enter number"
          value={number}
          onChange={(e) => setNumber(e.target.value)}
        />
        <button type="submit">Verify</button>
      </form>
      {response && <p>{response}</p>}
    </div>
  );
};

export default VerifyArmstrong;
