import axios from "axios";
import { webserviceUrl } from "./baseUrl";

export const apiLogin = async (loginData: Object) => {
  const apiUrl = `${webserviceUrl}auth/login`;
  try {
    const response = await axios.post(apiUrl, loginData, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
        const errorMessage = JSON.stringify(error.response?.data);
        throw errorMessage
      } else {
        throw error;
      }
  }
};
