import axios from "axios";
console.log({ baseURL: process.env.NEXT_PUBLIC_API_URL });
export const publicAxios = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true,
});
export const privateAxios = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true,
});
