import axios from "axios";
import React from "react";

export default function Todo() {
  const sendTodoRequest = async () => {
    try {
      const todoReq = await axios.post("/hola", { message: "Hola is keyword" });
      console.log(todoReq.data);
    } catch (err) {
      console.error(`err: ${err}`);
    }
  };

  return (
    <div>
      Todo
      <button onClick={sendTodoRequest}>Click</button>
    </div>
  );
}
