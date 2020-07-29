import axios from "axios";

const API_URL = "http://localhost:5000";

export function addPost(postData) {
  return axios.post(`${API_URL}/posts`, postData);
}

export function updatePost(postData) {
  return axios.put(`${API_URL}/posts/${postData.id}`, postData);
}

export function getPosts() {
  return axios.get(`${API_URL}/posts`);
}

export function getUserPosts(userId) {
  return axios.get(`${API_URL}/users/${userId}/posts`);
}

export function getPost(postId) {
  return axios.get(`${API_URL}/posts/${postId}`);
}

export function updatePublishedStatus(postId, published) {
  return axios.patch(`${API_URL}/posts/${postId}/published`, {
    published,
  });
}

export function deletePost(postId) {
  return axios.delete(`${API_URL}/posts/${postId}`);
}
