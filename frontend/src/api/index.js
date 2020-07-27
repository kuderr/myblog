import axios from "axios";

const API_URL = "http://localhost:5000";

export function addPost(postData) {
  return axios.post(`${API_URL}/posts`, postData);
}

export function getPosts() {
  return axios.get(`${API_URL}/posts`);
}

export function getPost(postId) {
  return axios.get(`${API_URL}/posts/${postId}`);
}