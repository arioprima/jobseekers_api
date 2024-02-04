import axios from "axios";
import { webserviceUrl } from "./baseUrl";

export const apiLogin = async (loginData: any) => {
  const apiUrl = `${webserviceUrl}auth/login`;
  try {
    const response = await axios.post(apiUrl, loginData, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    if (error.response && error.response.data) {
      return error.response.data;
    } else {
      throw error;
      
    }
  }
};
