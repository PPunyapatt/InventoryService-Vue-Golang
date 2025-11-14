import axios from "axios";

export default axios.create({
  baseURL: import.meta.env.BASE_URL ,//"http://localhost:1234/api/v1",
  headers: {
    "Content-type": "application/json"
  }
});