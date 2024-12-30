import axios from "axios";

const axiosCustom = axios.create({
    baseURL: process.env.API_BASE_URL || 'http://localhost:8080/api', // Fallback to default
});

export default axiosCustom;
