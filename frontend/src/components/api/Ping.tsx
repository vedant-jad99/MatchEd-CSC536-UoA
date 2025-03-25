import axios from "axios";
import { useEffect, useState } from "react";

const PingComponent = () => {
  const [message, setMessage] = useState<string>("");

  useEffect(() => {
    axios.get("http://localhost:3000/api/ping")
      .then(response => setMessage(response.data.message))
      .catch(error => console.error("Error:", error));
  }, []);

  return <div>Server says: {message || "Waiting for response..."}</div>;
};

export default PingComponent;
